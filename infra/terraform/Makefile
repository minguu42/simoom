.DEFAULT_GOAL := help
.PHONY: fmt plan-stg apply-stg plan-prod apply-prod help

fmt: ## Terraformコードを整形する
	@terraform fmt -recursive

plan-datastore-stg: ## 検証環境でデータストア関連のリソースに対していterraform planコマンドを実行する
	terraform -chdir=./datastore init -backend-config=stg.tfbackend -reconfigure
	terraform -chdir=./datastore plan -var-file=../stg.tfvars

apply-datastore-stg: ## 検証環境でデータストア関連のリソースに対してterraform applyコマンドを実行する
	terraform -chdir=./datastore init -backend-config=stg.tfbackend -reconfigure
	terraform -chdir=./datastore apply -var-file=../stg.tfvars

plan-datastore-prod: ## 本番環境でデータストア関連のリソースに対していterraform planコマンドを実行する
	terraform -chdir=./datastore init -backend-config=prod.tfbackend -reconfigure
	terraform -chdir=./datastore plan -var-file=../prod.tfvars

apply-datastore-prod: ## 本番環境でデータストア関連のリソースに対してterraform applyコマンドを実行する
	terraform -chdir=./datastore init -backend-config=prod.tfbackend -reconfigure
	terraform -chdir=./datastore apply -var-file=../prod.tfvars

plan-network-stg: ## 検証環境でネットワーク関連のリソースに対してterraform planコマンドを実行する
	terraform -chdir=./network init -backend-config=stg.tfbackend -reconfigure
	terraform -chdir=./network plan -var-file=../stg.tfvars

apply-network-stg: ## 検証環境でネットワーク関連のリソースに対してterraform applyコマンドを実行する
	terraform -chdir=./network init -backend-config=stg.tfbackend -reconfigure
	terraform -chdir=./network apply -var-file=../stg.tfvars

plan-network-prod: ## 本番環境でネットワーク関連のリソースに対してterraform planコマンドを実行する
	terraform -chdir=./network init -backend-config=prod.tfbackend -reconfigure
	terraform -chdir=./network plan -var-file=../prod.tfvars

apply-network-prod: ## 本番環境でネットワーク関連のリソースに対してterraform applyコマンドを実行する
	terraform -chdir=./network init -backend-config=prod.tfbackend -reconfigure
	terraform -chdir=./network apply -var-file=../prod.tfvars

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
      | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
