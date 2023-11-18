fmt: ## Protocol Buffersを整形する
	@buf format --write

lint: ## Protocol Buffersの静的解析を実行する
	@buf lint

gen: ## Connectのコードを生成する
	@buf generate

migrate: ## DBのスキーマの変更を確認する
	@mysqldef -u root -h 127.0.0.1 --dry-run --enable-drop-table simoomdb < ./mysql/schema.sql

migrate-apply: ## DBのスキーマの変更を適用する
	@mysqldef -u root -h 127.0.0.1 --enable-drop-table simoomdb < ./mysql/schema.sql
	@mysqldef -u root -h 127.0.0.1 --enable-drop-table simoomdb_test < ./mysql/schema.sql
