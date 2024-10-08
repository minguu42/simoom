resource "aws_db_instance" "main" {
  identifier     = "${local.product}-${var.env}-01"
  engine         = "mysql"
  engine_version = "8.0.32"
  instance_class = "db.t4g.micro"
  # vCPU: 2, Memory: 1GiB, EBS Burst Bandwidth: Up to 2085Mbps, Network Performance: Up to 5Gbps
  storage_type           = "gp2"
  allocated_storage      = 20
  username               = data.aws_ssm_parameter.db_master_username.value
  password               = data.aws_ssm_parameter.db_master_password.value
  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.rds.id]
  parameter_group_name   = aws_db_parameter_group.main.name
  skip_final_snapshot    = true # 検証用のため
  tags = {
    Name = "${local.product}-${var.env}-01"
  }
}

resource "aws_db_subnet_group" "main" {
  name       = "${local.product}-${var.env}"
  subnet_ids = [aws_subnet.private_a.id, aws_subnet.private_c.id]
  tags = {
    Name = "${local.product}-${var.env}"
  }
}

resource "aws_db_parameter_group" "main" {
  name   = "${local.product}-${var.env}"
  family = "mysql8.0"
  tags = {
    Name = "${local.product}-${var.env}"
  }
}
