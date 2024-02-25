output "alb_api_logs_bucket_id" {
  value = aws_s3_bucket.alb_api_logs.id
}

output "api_repository_url" {
  value = aws_ecr_repository.api.repository_url
}

output "api_secrets_arn" {
  value = aws_secretsmanager_secret.api_secrets.arn
}

output "private_subnet_ids" {
  value = [aws_subnet.private_a.id, aws_subnet.private_c.id]
}

output "public_subnet_ids" {
  value = [aws_subnet.public_a.id, aws_subnet.public_c.id]
}

output "vpc_id" {
  value = aws_vpc.main.id
}
