#!make
include .env


go-dev:
	go run .

go-test:
	go test -v ./...

go-build:
	go build -o main .


docker-up:
	docker-compose up -d

docker-db:
	docker-compose up db -d

docker-down:
	docker-compose down


migrate-up:
	migrate -path database/migration -database "${MYSQL_DIALEG}://${MYSQL_DSN}" -verbose up

migrate-down:
	migrate -path database/migration -database "${MYSQL_DIALEG}://${MYSQL_DSN}" -verbose down
