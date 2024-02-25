resource "aws_ecs_cluster" "api" {
  name = "${local.product}-${var.env}-api"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "aws_ecs_service" "api" {
  name            = "${local.product}-${var.env}-api"
  cluster         = aws_ecs_cluster.api.id
  task_definition = aws_ecs_task_definition.api.arn
  launch_type     = "FARGATE"
  desired_count   = 2
  network_configuration {
    subnets         = data.terraform_remote_state.main.outputs.private_subnet_ids
    security_groups = [aws_security_group.ecs_api.id]
  }
  load_balancer {
    target_group_arn = aws_lb_target_group.api.arn
    container_name   = "${local.product}-api"
    container_port   = 8080
  }
  deployment_circuit_breaker {
    enable   = true
    rollback = true
  }
}

resource "aws_security_group" "ecs_api" {
  name   = "${local.product}-${var.env}-ecs-api"
  vpc_id = data.terraform_remote_state.main.outputs.vpc_id
}

resource "aws_vpc_security_group_ingress_rule" "ecs_api_ingress" {
  security_group_id = aws_security_group.ecs_api.id
  from_port         = 8080
  to_port           = 8080
  ip_protocol       = "tcp"
  cidr_ipv4         = "10.0.0.0/16"
}

resource "aws_vpc_security_group_egress_rule" "ecs_api_egress" {
  security_group_id = aws_security_group.ecs_api.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_ecs_task_definition" "api" {
  family                   = "${local.product}-${var.env}-api"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 256 # .25vCPU
  memory                   = 512 # .5GB
  execution_role_arn       = aws_iam_role.ecs_api_execution.arn
  container_definitions = jsonencode([
    {
      name  = "${local.product}-api"
      image = "${data.terraform_remote_state.main.outputs.api_repository_url}:${var.api_image_tag}"
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      secrets = [
        {
          name      = "ACCESS_TOKEN_SECRET"
          valueFrom = "${data.terraform_remote_state.main.outputs.api_secrets_arn}:access_token_secret::"
        },
        {
          name      = "REFRESH_TOKEN_SECRET"
          valueFrom = "${data.terraform_remote_state.main.outputs.api_secrets_arn}:refresh_token_secret::"
        },
        {
          name      = "DB_HOST"
          valueFrom = "${data.terraform_remote_state.main.outputs.api_secrets_arn}:db_host::"
        },
        {
          name      = "DB_PORT"
          valueFrom = "${data.terraform_remote_state.main.outputs.api_secrets_arn}:db_port::"
        },
        {
          name      = "DB_DATABASE"
          valueFrom = "${data.terraform_remote_state.main.outputs.api_secrets_arn}:db_database::"
        },
        {
          name      = "DB_USER"
          valueFrom = "${data.terraform_remote_state.main.outputs.api_secrets_arn}:db_user::"
        },
        {
          name      = "DB_PASSWORD"
          valueFrom = "${data.terraform_remote_state.main.outputs.api_secrets_arn}:db_password::"
        },
      ]
      logConfiguration = {
        "logDriver" : "awslogs",
        "options" : {
          "awslogs-create-group" : "true",
          "awslogs-group" : "/ecs/simoom-stg-api",
          "awslogs-region" : "ap-northeast-1",
          "awslogs-stream-prefix" : "ecs"
        },
        "secretOptions" : []
      }
    }
  ])
  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }
}

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
          "secretsmanager:GetSecretValue"
        ],
        "Resource" : [
          data.terraform_remote_state.main.outputs.api_secrets_arn
        ]
      }
    ]
  })
}
