.DEFAULT_GOAL := help
.PHONY: plan-stg apply-stg plan-prod apply-prod help

plan-stg: ## 検証環境でterraform planコマンドを実行する
	terraform init -backend-config=stg.tfbackend -reconfigure
	terraform plan -var-file=stg.tfvars

apply-stg: ## 検証環境でterraform applyコマンドを実行する
	terraform init -backend-config=stg.tfbackend -reconfigure
	terraform apply -var-file=stg.tfvars

plan-prod: ## 本番環境でterraform planコマンドを実行する
	terraform init -backend-config=prod.tfbackend -reconfigure
	terraform plan -var-file=prod.tfvars

apply-prod: ## 本番環境でterraform applyコマンドを実行する
	terraform init -backend-config=prod.tfbackend -reconfigure
	terraform apply -var-file=prod.tfvars

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
          | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
ifeq ($(words $(MAKEFILE_LIST)), 2)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(lastword $(MAKEFILE_LIST)) \
          | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
endif
