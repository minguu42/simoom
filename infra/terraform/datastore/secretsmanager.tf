resource "aws_secretsmanager_secret" "credential" {
  name = "${local.product}-${var.env}-credential"
  tags = {
    Name = "${local.product}-${var.env}-credential"
  }
}
