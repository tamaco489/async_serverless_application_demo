output "original_image" {
  value = {
    id                          = aws_s3_bucket.original_image.id
    arn                         = aws_s3_bucket.original_image.arn
    bucket                      = aws_s3_bucket.original_image.bucket
    bucket_regional_domain_name = aws_s3_bucket.original_image.bucket_regional_domain_name
  }
}

output "thumbnail_image" {
  value = {
    id                          = aws_s3_bucket.thumbnail_image.id
    arn                         = aws_s3_bucket.thumbnail_image.arn
    bucket                      = aws_s3_bucket.thumbnail_image.bucket
    bucket_regional_domain_name = aws_s3_bucket.thumbnail_image.bucket_regional_domain_name
  }
}
