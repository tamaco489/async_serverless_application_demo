# オリジナルコンテンツ画像管理バケット
resource "aws_s3_bucket" "original_image" {
  bucket = "${var.env}-nautilus-original-image"

  # terraformによる削除処理の防止
  lifecycle {
    prevent_destroy = true
  }

  tags = { Name = "${var.env}-nautilus-original-image" }
}

# オリジナルコンテンツ画像管理バケットのCORSポリシー設定
resource "aws_s3_bucket_cors_configuration" "original_image" {
  bucket = aws_s3_bucket.original_image.bucket

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT"]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }
}

# サムネイルコンテンツ画像管理バケット
resource "aws_s3_bucket" "thumbnail_image" {
  bucket = "${var.env}-nautilus-thumbnail-image"

  # terraformによる削除処理の防止
  lifecycle {
    prevent_destroy = true
  }

  tags = { Name = "${var.env}-nautilus-thumbnail-image" }
}

# サムネイルコンテンツ画像管理バケットのCORSポリシー設定
resource "aws_s3_bucket_cors_configuration" "thumbnail_image" {
  bucket = aws_s3_bucket.thumbnail_image.bucket

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT"]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }
}
