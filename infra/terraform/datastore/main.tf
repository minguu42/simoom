terraform {
  required_version = ">= 1.0.0, < 2.0.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.36.0"
    }
  }
  backend "s3" {}
}

provider "aws" {
  region = "ap-northeast-1"
  default_tags {
    tags = {
      Product   = local.product
      Env       = var.env
      Owner     = "minguu42"
      ManagedBy = "terraform"
    }
  }
}

data "terraform_remote_state" "network" {
  backend = "s3"
  config = {
    bucket = "${local.product}-${var.env}-tfstate"
    key    = "network/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

data "aws_secretsmanager_secret_version" "credential" {
  secret_id = aws_secretsmanager_secret.credential.id
}

variable "env" {
  type = string
  validation {
    condition     = contains(["prod", "stg"], var.env)
    error_message = "Only the following environment is allowed: prod | stg"
  }
}

locals {
  product      = "simoom"
  isProduction = var.env == "prod"
  credential = jsondecode(data.aws_secretsmanager_secret_version.credential.secret_string)
}

output "ecr_repository_api_repository_url" {
  value = aws_ecr_repository.api.repository_url
}

#output "s3_bucket_lb_api_logs_id" {
#  value = aws_s3_bucket.lb_api_logs.id
#}
