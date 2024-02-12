resource "aws_instance" "bastion" {
  ami                    = "ami-0b5c74e235ed808b9" # Amazon Linux 2023 AMI
  instance_type          = "t2.micro"
  key_name               = "${local.product}-${var.env}-bastion"
  subnet_id              = data.terraform_remote_state.network.outputs.public_a_subnet_id
  vpc_security_group_ids = [aws_security_group.bastion.id]
  root_block_device {
    volume_type = "gp3"
    volume_size = 30
  }
  tags = {
    Name = "${local.product}-${var.env}-bastion"
  }
}

resource "aws_eip" "bastion" {
  instance = aws_instance.bastion.id
  domain   = "vpc"
  tags     = {
    Name = "${local.product}-${var.env}-bastion"
  }
}

resource "aws_security_group" "bastion" {
  name   = "${local.product}-${var.env}-bastion"
  vpc_id = data.terraform_remote_state.network.outputs.vpc_id
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
