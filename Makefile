.DEFAULT_GOAL := help
.PHONY: setup migrate migrate-apply gen fmt fmt-go fmt-proto fmt-tf
.PHONY: lint lint-go lint-proto test build help

setup: ## 開発に必要なツールをインストールする
	brew install sqldef/sqldef/mysqldef
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/matryer/moq@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

migrate: ## DBのスキーマの変更を確認する
	@mysqldef -u root -h 127.0.0.1 --dry-run --enable-drop-table --file=./infra/mysql/schema.sql simoomdb

migrate-apply: ## DBのスキーマの変更を適用する
	@mysqldef -u root -h 127.0.0.1 --enable-drop-table --file=./infra/mysql/schema.sql simoomdb

gen: ## コードを生成する
	@buf generate
	@rm -rf ./api/infra/mysql/sqlc && sqlc generate
	@$(MAKE) fmt

fmt: fmt-go fmt-proto fmt-tf ## コードを整形する

fmt-go: # Goコードを整形する
	@goimports -w .

fmt-proto: # protoコードを整形する
	@buf format --write

fmt-tf: # Terraformコードを整形する
	@terraform fmt -recursive

lint: lint-go lint-proto ## 静的解析を実行する

lint-go: # Goファイルの静的解析を実行する
	@go vet $$(go list ./... | grep -v -e /simoompb -e /sqlc)
	@staticcheck $$(go list ./... | grep -v -e /simoompb -e /sqlc)

lint-proto: # protoファイルの静的解析を実行する
	@buf lint

test: ## テストを実行する
	@go test $$(go list ./... | grep -v -e /simoompb -e /sqlc)

build: ## 本番用コンテナイメージをビルドする
	@docker image build \
            --tag=simoom-api:latest \
            --target=prod .

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
      | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
