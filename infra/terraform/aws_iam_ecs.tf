resource "aws_iam_role" "ecs_api_execution" {
  name = "${local.product}-${var.env}-ecs-api-task-execution"
  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "",
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "ecs-tasks.amazonaws.com"
        },
        "Action" : "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy" "ecs_api_execution" {
  name = "${local.product}-${var.env}-ecs-api-execution"
  role = aws_iam_role.ecs_api_execution.id
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : [
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
        ],
        "Resource" : "*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "ssm:GetParameters"
        ],
        "Resource" : [
          data.aws_ssm_parameter.api_access_token_secret.arn,
          data.aws_ssm_parameter.api_refresh_token_secret.arn,
          data.aws_ssm_parameter.db_database.arn,
          data.aws_ssm_parameter.db_host.arn,
          data.aws_ssm_parameter.db_password.arn,
          data.aws_ssm_parameter.db_port.arn,
          data.aws_ssm_parameter.db_user.arn,
        ]
      }
    ]
  })
}
