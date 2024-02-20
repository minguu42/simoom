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

data "aws_secretsmanager_secret_version" "api_secrets" {
  secret_id = aws_secretsmanager_secret.api_secrets.id
}

variable "env" {
  type = string
  validation {
    condition     = contains(["prod", "stg"], var.env)
    error_message = "Only the following environment is allowed: prod | stg"
  }
}

variable "api_image_tag" {
  type = string
  validation {
    condition     = length(var.api_image_tag) == 40
    error_message = "The api image tag is a 40-character string"
  }
}

locals {
  product      = "simoom"
  isProduction = var.env == "prod"
  api_secrets  = jsondecode(data.aws_secretsmanager_secret_version.api_secrets.secret_string)
}
