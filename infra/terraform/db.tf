resource "aws_db_instance" "main" {
  identifier             = "${local.product}-${var.env}-01"
  engine                 = "mysql"
  engine_version         = "8.0.32"
  instance_class         = "db.t4g.micro" # vCPU: 2, Memory: 1GiB, EBS Burst Bandwidth: Up to 2085Mbps, Network Performance: Up to 5Gbps
  storage_type           = "gp2"
  allocated_storage      = 20
  username               = local.api_secrets["db_master_username"]
  password               = local.api_secrets["db_master_password"]
  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.rds.id]
  parameter_group_name   = aws_db_parameter_group.main.name
  skip_final_snapshot    = true # 検証用のため
  tags                   = {
    Name = "${local.product}-${var.env}-01"
  }
}

resource "aws_db_subnet_group" "main" {
  name       = "${local.product}-${var.env}"
  subnet_ids = [aws_subnet.private_a.id, aws_subnet.private_c.id]
  tags       = {
    Name = "${local.product}-${var.env}"
  }
}

resource "aws_security_group" "rds" {
  name   = "${local.product}-${var.env}-rds"
  vpc_id = aws_vpc.main.id
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

resource "aws_db_parameter_group" "main" {
  name   = "${local.product}-${var.env}"
  family = "mysql8.0"
  tags   = {
    Name = "${local.product}-${var.env}"
  }
}
