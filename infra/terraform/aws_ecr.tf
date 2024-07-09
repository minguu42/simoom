resource "aws_ecr_repository" "api" {
  name                 = "${local.product}-${var.env}-api"
  image_tag_mutability = "IMMUTABLE"
}
