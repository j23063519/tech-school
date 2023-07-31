# Learning Golang from Tech School

## Project Introduction

I'm following the "tech school" video tutorial to learn golang、docker、postgres、k8s、aws、grpc、redis etc knowledges.

reference: https://github.com/techschool/simplebank

## Project Excution

<!-- excution all docker -->
1. make up
<!-- create table -->
2. make migrateup
<!-- generate CRUD go file from sqlc -->
3. make sqlc

## Folder Structure

- .github
  - workflows
- api
- db
  - migration
  - query
  - sqlc
- util

### Folder Structure Explanation

- `.github: git`: 這裡存放 git 版本控制相關的目錄:
  1. `workflows: github action`: 存放 ci/cd 相關設定xxx.yml檔案
- `api: restful api`: 這裡存放 api 相關檔案包含middleware及validation
- `db: 資料庫`:這裡會存放 database 資料庫相關的目錄，此資料夾底下會有:
  1. `migration: 資料庫遷移`: 存放 資料庫遷移等 sql檔案
  2. `query: CRUD sql`: 存放 新刪修查等 sql檔案
  3. `sqlc: sqlc 所產生的檔案`: 存放 sqlc 所產生的 新刪修查等 go檔案
- `util: 輔助套件`:這裡會存放一些第三方套件或自訂套件

## Go Packages

### golang-migrate <https://github.com/golang-migrate/migrate#cli-usage>
```bash
brew install golang-migrate
```

### sqlc <https://github.com/kyleconroy/sqlc>
```bash
brew install sqlc
```

### lib/pq <https://github.com/lib/pq>
```bash
go get github.com/lib/pq
```

### testify <https://github.com/stretchr/testify>
```bash
go get github.com/stretchr/testify
```

### gin <https://github.com/gin-gonic/gin>
```bash
go get -u github.com/gin-gonic/gin
```

## command line

```bash
# 查詢 docker run 的歷史紀錄 (mac)
history | grep "docker run"

# 建立 image
make build

# 建立新的 image 並在背景執行容器(應用程式)
make up

# 結束應用程式、清除所有容器包含 image 及 掛載的資料
make down

# 啟動應用程式
make start

# 重啟應用程式
make restart

# 暫停應用程式
make stop

# 刪除且停止容器
make rm

# 執行測試
make test

# 刪除所有未使用的 image
make rmimg 

# 刪除所有未使用的 container image network cache
make rmsys

# 執行並進入db容器
make execdb

# 建立資料庫 資料庫名稱須自行更改
make createdb

# 刪除資料庫 資料庫名稱須自行更改
make dropdb

# 執行新增資料表、欄位、索引
make migrateup

# 回朔所新增的資料表、欄位、索引
make migratedown

# 根據 sql檔案 而產生 go檔案
make sqlc
```