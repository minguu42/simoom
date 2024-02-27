## 料金のかかるリソースを削除するコマンド

### main

```bash
terraform -chdir=./infra/terraform init -backend-config=stg.tfbackend -reconfigure
```

```bash
terraform -chdir=./infra/terraform destroy \
  -target=aws_eip.nat_a \
  -target=aws_nat_gateway.a \
  -target=aws_route_table.private_c
```

### api

```bash
terraform -chdir=./infra/terraform/api init -backend-config=stg.tfbackend -reconfigure
```

```bash
terraform -chdir=./infra/terraform/api destroy
```

## Terraformの管理対象外

- AWS Systems Manager Parameter Storeのパラメータ
