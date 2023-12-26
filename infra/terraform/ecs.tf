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
}

resource "aws_ecs_task_definition" "api" {
  family                   = "${local.product}-${var.env}-api"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 256 # .25vCPU
  memory                   = 512 # .5GB
  execution_role_arn       = aws_iam_role.ecs_api_task_execution.arn
  container_definitions    = jsonencode([
    {
      name         = "${local.product}-api"
      image        = "${aws_ecr_repository.api.repository_url}:d207d40"
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
    }
  ])
  runtime_platform {
    cpu_architecture        = "X86_64"
    operating_system_family = "LINUX"
  }
}
