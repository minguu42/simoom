resource "aws_secretsmanager_secret" "api_secrets" {
  name = "${local.product}-${var.env}-api-secrets"
  tags = {
    Name = "${local.product}-${var.env}-api-secrets"
  }
}
