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
    subnets         = [aws_subnet.private_a.id, aws_subnet.private_c.id]
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
      image = "${aws_ecr_repository.api.repository_url}:${var.api_image_tag}"
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      secrets = [
        {
          name      = "ACCESS_TOKEN_SECRET"
          valueFrom = data.aws_ssm_parameter.api_access_token_secret.arn
        },
        {
          name      = "REFRESH_TOKEN_SECRET"
          valueFrom = data.aws_ssm_parameter.api_refresh_token_secret.arn
        },
        {
          name      = "DB_HOST"
          valueFrom = data.aws_ssm_parameter.db_host.arn
        },
        {
          name      = "DB_PORT"
          valueFrom = data.aws_ssm_parameter.db_port.arn
        },
        {
          name      = "DB_DATABASE"
          valueFrom = data.aws_ssm_parameter.db_database.arn
        },
        {
          name      = "DB_USER"
          valueFrom = data.aws_ssm_parameter.db_user.arn
        },
        {
          name      = "DB_PASSWORD"
          valueFrom = data.aws_ssm_parameter.db_password.arn
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
