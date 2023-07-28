package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// run n concurrent transfer transations
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, CreateTransferParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetEntry(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check account's balance
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balance
	updateAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updateAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updateAccount1.Balance, updateAccount2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updateAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updateAccount2.Balance)
}

// 模擬 deadlock in postgresql
// BEGIN;

// INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES (1, 2, 10) RETURNING *;

// INSERT INTO entries (account_id, amount) VALUES (1, -10) RETURNING *;
// INSERT INTO entries (account_id, amount) VALUES (2, 10) RETURNING *;

// SELECT * FROM accounts WHERE id = 1 FOR UPDATE;
// UPDATE accounts SET balance = 90 WHERE id = 1 RETURNING *;

// SELECT * FROM accounts WHERE id = 2 FOR UPDATE;
// UPDATE accounts SET balance = 110 WHERE id = 2 RETURNING *;

// ROLLBACK;

// =======================================================
// reference: https://wiki.postgresql.org/wiki/Lock_Monitoring
// 查詢 lock
// SELECT blocked_locks.pid     AS blocked_pid,
// 		  blocked_activity.usename  AS blocked_user,
// 		  blocking_locks.pid     AS blocking_pid,
// 		  blocking_activity.usename AS blocking_user,
// 		  blocked_activity.query    AS blocked_statement,
// 		  blocking_activity.query   AS current_statement_in_blocking_process
// FROM pg_catalog.pg_locks         blocked_locks
// JOIN pg_catalog.pg_stat_activity blocked_activity  ON blocked_activity.pid = blocked_locks.pid
// JOIN pg_catalog.pg_locks         blocking_locks
// 	ON blocking_locks.locktype = blocked_locks.locktype
// 	AND blocking_locks.database IS NOT DISTINCT FROM blocked_locks.database
// 	AND blocking_locks.relation IS NOT DISTINCT FROM blocked_locks.relation
// 	AND blocking_locks.page IS NOT DISTINCT FROM blocked_locks.page
// 	AND blocking_locks.tuple IS NOT DISTINCT FROM blocked_locks.tuple
// 	AND blocking_locks.virtualxid IS NOT DISTINCT FROM blocked_locks.virtualxid
// 	AND blocking_locks.transactionid IS NOT DISTINCT FROM blocked_locks.transactionid
// 	AND blocking_locks.classid IS NOT DISTINCT FROM blocked_locks.classid
// 	AND blocking_locks.objid IS NOT DISTINCT FROM blocked_locks.objid
// 	AND blocking_locks.objsubid IS NOT DISTINCT FROM blocked_locks.objsubid
// 	AND blocking_locks.pid != blocked_locks.pid

// JOIN pg_catalog.pg_stat_activity blocking_activity ON blocking_activity.pid = blocking_locks.pid
// WHERE NOT blocked_locks.granted;

// =======================================================
// 列出在database中所有的locks
// SELECT a.datname,
// 		a.application_name,
//          l.relation::regclass,
//          l.transactionid,
//          l.mode,
//          l.locktype,
//          l.GRANTED,
//          a.usename,
//          a.query,
//          a.query_start,
//          age(now(), a.query_start) AS "age",
//          a.pid
// FROM pg_stat_activity a
// JOIN pg_locks l ON l.pid = a.pid
// where a.application_name = 'psql'
// ORDER BY a.pid;

// =======================================================
// 試錯的執行步驟:
// 開啟兩個terminal:分別為 a 與 b
// 分別都進入db容器內，make execdb
// a 執行：BEGIN;
// b 執行：BEGIN;
// b 執行：INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES (1, 2, 10) RETURNING *;
// b 執行：INSERT INTO entries (account_id, amount) VALUES (1, -10) RETURNING *;
// a 執行：INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES (1, 2, 10) RETURNING *;
// b 執行：INSERT INTO entries (account_id, amount) VALUES (2, 10) RETURNING *;
// b 執行：SELECT * FROM accounts WHERE id = 1 FOR UPDATE; (這時就會lock，因為他正等著另個transaction commit 或 rollback 才能繼續執行)
// a 執行：INSERT INTO entries (account_id, amount) VALUES (1, -10) RETURNING *;
// a 執行：INSERT INTO entries (account_id, amount) VALUES (2, 10) RETURNING *;
// a 執行：SELECT * FROM accounts WHERE id = 1 FOR UPDATE; (這時就會得到deadlock，因為他要等待另一個transaction commit 或 rollback，但是另一個也在等他，互相等待形成"死鎖")

// =======================================================
// 若要解決此情形發生，得將原本 SELECT * FROM accounts WHERE id = $1 LIMIT 1; 該語句換成
// SELECT * FROM accounts WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;
// 還有另一種解法為將下述sql給註解起來，但是通常不建議這樣做
// ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
// ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
// ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
// 上述兩個解決方式詳解：
// 由於 postgres 擔心 transaction 1 會更改到 accounts表中的欄位 ID，這樣會影響到 transfers 表的外鍵約束，進而lock住
// 所以我們必須告訴 postgres 我們不會改變 id 的值也就是(FOR NO KEY UPDATE)，這樣就不會產生 deadlock 了
