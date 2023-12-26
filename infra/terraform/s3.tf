resource "aws_s3_bucket" "tfstate" {
  bucket = "${local.product}-${var.env}-tfstate"
  lifecycle {
    prevent_destroy = true
  }
}

# ステートファイルの完全な履歴が見れるように、バージョニングを有効化する
resource "aws_s3_bucket_versioning" "tfstate" {
  bucket = aws_s3_bucket.tfstate.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket" "lb_api_logs" {
  bucket = "${local.product}-${var.env}-lb-api-logs"
  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_s3_bucket_policy" "lb_api_logs" {
  bucket = aws_s3_bucket.lb_api_logs.id
  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Effect    = "Allow"
        Principal = { AWS = "arn:aws:iam::582318560864:root" }
        Action    = "s3:PutObject"
        Resource  = "${aws_s3_bucket.lb_api_logs.arn}/*"
      }
    ]
  })
}
