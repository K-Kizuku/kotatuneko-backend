ENV_FILE := .env
ENV = $(shell grep -v "^\#" $(ENV_FILE))
include .env

.PHONY: lint lint-fix run logs migrate migrate-down gen-migrate cloud-sql-proxy setup-proxy wire

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./...  --fix

run:
	docker compose up --build -d

logs:
	docker compose logs -f

migrate:
	migrate -path db/migrations -database "postgres://user:password@localhost:5432/db?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://user:password@localhost:5432/db?sslmode=disable" down

gen-migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

# gcloud auth loginを実行ずみであること
# gcloud set project PROJECT_IDを実行ずみであること
setup-proxy:
	curl -o cloud-sql-proxy https://storage.googleapis.com/cloud-sql-connectors/cloud-sql-proxy/v2.1.2/cloud-sql-proxy.darwin.amd64
	chmod 744 cloud-sql-proxy
	gcloud components install cloud_sql_proxy --quiet
	rm cloud-sql-proxy

cloud-sql-proxy:
	cloud_sql_proxy hacku-416915:asia-northeast1:hacku-postgres=tcp:0.0.0.0:5432 --credential-file=key.json

wire:
	wire gen ./internal/di/wire.go

commit:
	npx git-cz