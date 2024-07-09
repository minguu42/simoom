resource "aws_security_group" "ecs_api" {
  name   = "${local.product}-${var.env}-ecs-api"
  vpc_id = aws_vpc.main.id
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
