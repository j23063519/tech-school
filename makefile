include local.env

.PHONY : build up down start restart stop rm test rmimg rmsys execdb createdb dropdb migrateup migratedown sqlc server mock

build:
	docker compose -f docker-compose.${APP_ENV}.yaml --env-file ${APP_ENV}.env build

# -d: 表示背景執行, --build: 重新建立新的image
up:
	docker compose -f docker-compose.${APP_ENV}.yaml --env-file ${APP_ENV}.env up -d --build

# -v: 表示連資料也清空 --rmi type: 刪除image，後面的type是local或是all
down:
	docker compose -f docker-compose.${APP_ENV}.yaml down -v --rmi all

start:
	docker compose -f docker-compose.${APP_ENV}.yaml start

restart:
	docker compose -f docker-compose.${APP_ENV}.yaml restart

stop:
	docker compose -f docker-compose.${APP_ENV}.yaml stop

rm:
	docker compose -f docker-compose.${APP_ENV}.yaml rm --force

# -v: 顯示測試結果 -cover: 顯示測試覆蓋率
test:
	go test -v -cover ./...

# 刪除所有 unused images
rmimg:
	docker image prune -a -f

# 刪除所有 unused containers images networks cache
rmsys:
	docker system prune -a -f

# -it: 此選項表示我們要進入互動式（Interactive）模式 psql: 在容器內執行的命令，即PostgreSQL的命令行客戶端 
# -U: 使用指定的使用者（user）登入PostgreSQL資料庫 -d: 連接到指定的資料庫
execdb:
	docker exec -it ${POSTGRES_CONTAINER_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

# example_db只是範例名稱，可隨意更改
createdb:
	docker exec -it ${POSTGRES_CONTAINER_NAME} createdb -U ${POSTGRES_USER} --owner=${POSTGRES_USER} example_db

# example_db只是範例名稱，可隨意更改
dropdb:
	docker exec -it ${POSTGRES_CONTAINER_NAME} dropdb -U ${POSTGRES_USER} example_db

migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/j23063519/tech-school/db/sqlc Store
