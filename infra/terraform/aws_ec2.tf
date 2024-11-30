resource "aws_eip" "nat_a" {
  domain    = "vpc"
  tags = {
    Name = "${local.product}-${var.env}-nat-a"
  }
}

resource "aws_eip" "nat_c" {
  count  = local.isProduction ? 1 : 0
  domain = "vpc"
  tags = {
    Name = "${local.product}-${var.env}-nat-c"
  }
}

resource "aws_instance" "bastion" {
  ami                    = "ami-0b5c74e235ed808b9" # Amazon Linux 2023 AMI
  instance_type          = "t2.micro"              # vCPU: 1, Memory: 1.0GiB
  key_name               = aws_key_pair.bastion.key_name
  subnet_id              = aws_subnet.private_a.id
  vpc_security_group_ids = [aws_security_group.bastion.id]
  root_block_device {
    volume_type = "gp3"
    volume_size = 30
  }
  tags = {
    Name = "${local.product}-${var.env}-bastion"
  }
}

resource "aws_key_pair" "bastion" {
  key_name   = "${local.product}-${var.env}-bastion"
  public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAVHKIFgKq+Gzyx/u1yczsSEzM7bl9TnpuZUF2+Tjr6D"
}

resource "aws_lb" "api" {
  name               = "${local.product}-${var.env}-api-public"
  internal           = false
  load_balancer_type = "application"
  subnets            = [aws_subnet.public_a.id, aws_subnet.public_c.id]
  security_groups    = [aws_security_group.alb_api.id]
  access_logs {
    bucket  = aws_s3_bucket.alb_api_logs.id
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
