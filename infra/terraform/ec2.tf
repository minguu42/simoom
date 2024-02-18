resource "aws_eip" "bastion" {
  instance = aws_instance.bastion.id
  domain   = "vpc"
  tags     = {
    Name = "${local.product}-${var.env}-bastion"
  }
}

resource "aws_instance" "bastion" {
  ami                    = "ami-0b5c74e235ed808b9" # Amazon Linux 2023 AMI
  instance_type          = "t2.micro"              # vCPU: 1, Memory: 1.0GiB
  key_name               = aws_key_pair.bastion.key_name
  subnet_id              = aws_subnet.public_a.id
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
  key_name = "${local.product}-${var.env}-bastion"
  public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAVHKIFgKq+Gzyx/u1yczsSEzM7bl9TnpuZUF2+Tjr6D"
}

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
