resource "aws_dynamodb_table" "tfstate_lock" {
  name         = "${local.product}-${var.env}-tfstate-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"
  attribute {
    name = "LockID"
    type = "S"
  }
}
