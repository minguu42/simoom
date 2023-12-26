resource "aws_lb" "api" {
  name                       = "${local.product}-${var.env}-api-public"
  internal                   = false
  load_balancer_type         = "application"
  security_groups            = [aws_security_group.lb_api.id]
  subnets                    = [aws_subnet.public_a.id, aws_subnet.public_c.id]
  enable_deletion_protection = false # 検証のため
  access_logs {
    bucket  = aws_s3_bucket.lb_api_logs.id
    enabled = true
  }
}

resource "aws_lb_listener" "api_http" {
  load_balancer_arn = aws_lb.api.arn
  port              = 80
  protocol          = "HTTP"
  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "404 Not Found"
      status_code  = 404
    }
  }
}

resource "aws_lb_listener_rule" "api" {
  listener_arn = aws_lb_listener.api_http.arn
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.api.arn
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_lb_target_group" "api" {
  name        = "${local.product}-${var.env}-api-public"
  port        = 8080
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_vpc.main.id
  health_check {
    path = "/simoompb.v1.SimoomService/CheckHealth?encoding=json&message=%7b%7d"
  }
}
