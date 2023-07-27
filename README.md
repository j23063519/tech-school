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

- db
  - migration
  - query
  - sqlc

### Folder Structure Explanation

- `db: 資料庫`:這裡會存放 database 資料庫相關的目錄，此資料夾底下會有:
  1. `migration: 資料庫遷移`: 存放 資料庫遷移等 sql檔案:
  2. `query: CRUD sql`: 存放 新刪修查等 sql檔案
  3. `sqlc: sqlc 所產生的檔案`: 存放 sqlc 所產生的 新刪修查等 go檔案

## Go Packages

### golang-migrate <https://github.com/golang-migrate/migrate#cli-usage>
```bash
brew install golang-migrate
```

### sqlc <https://github.com/kyleconroy/sqlc>
```bash
brew install sqlc
```

## command line

```bash
# 查詢 docker run 的歷史紀錄 (mac)
history | grep "docker run"
```