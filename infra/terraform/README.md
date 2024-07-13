# インフラ管理

## コマンド

### init

```bash
terraform -chdir=./infra/terraform init -backend-config=stg.tfbackend -reconfigure
```

### apply

```bash
terraform -chdir=./infra/terraform apply
```

### destroy

```bash
terraform -chdir=./infra/terraform destroy \
  -target=aws_eip.nat_a \
  -target=aws_nat_gateway.a \
  -target=aws_route_table.private_c \
  -target=aws_lb.api \
  -target=aws_lb_listener.api_http \
  -target=aws_lb_listener_rule.api \
  -target=aws_lb_target_group.api \
  -target=aws_ecs_service.api
```

## Terraformの管理対象外

- AWS Systems Manager Parameter Storeのパラメータ
