# Tech School

## Project Introduction

I'm following the "tech school" video tutorial to learn golang、docker、postgres、k8s、aws、grpc、redis etc knowledges.

reference: https://github.com/techschool/simplebank

## Project Excution

<!-- create table -->
1. make migrateup
<!-- drop table -->
2. make migratedown

## Folder Structure

- db
  - migration

### Folder Structure Explanation

- `db: 資料庫`:這裡會存放 database 資料庫相關的目錄，此資料夾底下會有:
  1. `migration: 資料庫遷移`: 存放 資料庫遷移等 sql檔案:

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