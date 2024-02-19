resource "aws_lb" "api" {
  name               = "${local.product}-${var.env}-api-public"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb_api.id]
  subnets            = [aws_subnet.public_a.id, aws_subnet.public_c.id]
  access_logs {
    bucket  = aws_s3_bucket.alb_api_logs.id
    enabled = true
  }
}

resource "aws_security_group" "alb_api" {
  name   = "${local.product}-${var.env}-alb-api"
  vpc_id = aws_vpc.main.id
}

resource "aws_vpc_security_group_ingress_rule" "alb_api_ingress" {
  security_group_id = aws_security_group.alb_api.id
  from_port         = 80
  to_port           = 80
  ip_protocol       = "tcp"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_vpc_security_group_egress_rule" "alb_api_egress" {
  security_group_id = aws_security_group.alb_api.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
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
