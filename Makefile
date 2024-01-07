.DEFAULT_GOAL := help
.PHONY: setup gen build run migrate migrate-apply fmt check-style-go lint-go lint-protobuf lint test help

setup: ## 開発に必要なツールをインストールする
	brew install sqldef/sqldef/mysqldef
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

gen: ## コードを生成する
	@buf generate
	@sqlc generate
	@$(MAKE) fmt

build: ## 本番用APIサーバのコンテナイメージをビルドする
	@docker build \
            --tag=simoom-api:latest \
            --target=prod .

run: ## 本番用APIサーバを実行する
	@docker compose up -d db
	@docker container run \
            --env-file .env \
            --name simoom-api \
            --network=simoom_default \
            -p 8080:8080 \
            --rm \
            simoom-api

migrate: ## DBのスキーマの変更を確認する
	@mysqldef -u root -h 127.0.0.1 --dry-run --enable-drop-table simoomdb < ./infra/mysql/schema.sql

migrate-apply: ## DBのスキーマの変更を適用する
	@mysqldef -u root -h 127.0.0.1 --enable-drop-table simoomdb < ./infra/mysql/schema.sql
	@mysqldef -u root -h 127.0.0.1 --enable-drop-table testdb < ./infra/mysql/schema.sql

fmt: ## コードを整形する
	@buf format --write
	@goimports -w .

check-style-go:
	@if [ $(shell goimports -l . | wc -l) -gt 0 ]; then exit 1; fi

lint-go: # Goファイルの静的解析を実行する
	@go vet $$(go list ./... | grep -v -e /simoompb -e /sqlc)
	@staticcheck $$(go list ./... | grep -v -e /simoompb -e /sqlc)

lint-protobuf: # Protocol Buffersファイルの静的解析を実行する
	@buf lint

lint: lint-go lint-protobuf ## 静的解析を実行する

test: ## テストを実行する
	@go test $$(go list ./... | grep -v -e /simoompb -e /sqlc)

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
      | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
