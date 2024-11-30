resource "aws_ecr_repository" "api" {
  name                 = "${local.product}-${var.env}-api"
  image_tag_mutability  = "IMMUTABLE"
}

resource "aws_ecr_lifecycle_policy" "api" {
  repository = aws_ecr_repository.api.name
  policy = jsonencode({
    rules = [
      {
        rulePriority = 1,
        description  = "最新の3イメージのみを保持する",
        selection = {
          tagStatus   = "any",
          countType   = "imageCountMoreThan",
          countNumber = 3
        }
        action = {
          type = "expire"
        }
      }
    ]
  })
}
