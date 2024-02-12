resource "aws_db_subnet_group" "main" {
  name       = "${local.product}-${var.env}"
  subnet_ids = data.terraform_remote_state.network.outputs.private_subnet_ids
  tags       = {
    Name = "${local.product}-${var.env}"
  }
}

resource "aws_db_parameter_group" "main" {
  name   = "${local.product}-${var.env}"
  family = "mysql8.0"
  tags   = {
    Name = "${local.product}-${var.env}"
  }
}

resource "aws_db_instance" "main" {
  identifier             = "${local.product}-${var.env}-01"
  engine                 = "mysql"
  engine_version         = "8.0.32"
  instance_class         = "db.t4g.micro"
  storage_type           = "gp2"
  allocated_storage      = 20
  username               = local.credential["db_master_username"]
  password               = local.credential["db_master_password"]
  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.rds.id]
  parameter_group_name   = aws_db_parameter_group.main.name
  skip_final_snapshot    = true # 検証用のため
  tags                   = {
    Name = "${local.product}-${var.env}-01"
  }
}

resource "aws_security_group" "rds" {
  name   = "${local.product}-${var.env}-rds"
  vpc_id = data.terraform_remote_state.network.outputs.vpc_id
}

resource "aws_vpc_security_group_ingress_rule" "rds_ingress" {
  security_group_id = aws_security_group.rds.id
  from_port         = 3306
  to_port           = 3306
  ip_protocol       = "tcp"
  cidr_ipv4         = "10.0.0.0/16"
}

resource "aws_vpc_security_group_egress_rule" "rds_egress" {
  security_group_id = aws_security_group.rds.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}
