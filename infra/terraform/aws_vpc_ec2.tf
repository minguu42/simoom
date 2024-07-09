resource "aws_security_group" "bastion" {
  name   = "${local.product}-${var.env}-bastion"
  vpc_id = aws_vpc.main.id
}

resource "aws_vpc_security_group_ingress_rule" "bastion_ingress" {
  security_group_id = aws_security_group.bastion.id
  from_port         = 22
  to_port           = 22
  ip_protocol       = "tcp"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_vpc_security_group_egress_rule" "bastion_egress" {
  security_group_id = aws_security_group.bastion.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_security_group" "eic" {
  name   = "${local.product}-${var.env}-eic"
  vpc_id = aws_vpc.main.id
}

resource "aws_vpc_security_group_egress_rule" "eic_egress" {
  security_group_id = aws_security_group.eic.id
  from_port         = 22
  to_port           = 22
  ip_protocol       = "tcp"
  cidr_ipv4         = aws_subnet.private_a.cidr_block
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
