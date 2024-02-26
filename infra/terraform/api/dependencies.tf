data "terraform_remote_state" "main" {
  backend = "s3"
  config = {
    bucket = "${local.product}-${var.env}-tf-remote-state"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
  }
}

data "aws_ssm_parameter" "api_access_token_secret" {
  name = "/${local.product}/${var.env}/api/access_token_secret"
}

data "aws_ssm_parameter" "api_refresh_token_secret" {
  name = "/${local.product}/${var.env}/api/refresh_token_secret"
}

data "aws_ssm_parameter" "db_database" {
  name = "/${local.product}/${var.env}/db/database"
}

data "aws_ssm_parameter" "db_host" {
  name = "/${local.product}/${var.env}/db/host"
}

data "aws_ssm_parameter" "db_password" {
  name = "/${local.product}/${var.env}/db/password"
}

data "aws_ssm_parameter" "db_port" {
  name = "/${local.product}/${var.env}/db/port"
}

data "aws_ssm_parameter" "db_user" {
  name = "/${local.product}/${var.env}/db/user"
}
