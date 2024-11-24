resource "aws_lambda_function" "image_maker_batch" {
  function_name = "${var.env}-image-maker-batch"
  description   = "サムネイル画像生成サービスバッチ"
  role          = aws_iam_role.image_maker_batch.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.image_maker_batch.url}:image_maker_v0.0.0"
  timeout       = 20
  memory_size   = 256

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      SERVICE_NAME = "image-maker-batch"
      ENV          = "dev"
      S3_THUMBNAIL_IMAGE_BUCKET_NAME = data.terraform_remote_state.s3.outputs.thumbnail_image.id
    }
  }

  tags = { Name = "${var.env}-image-maker-batch" }
}

# LambdaがS3からPUTイベントを受け取る権限を付与
resource "aws_lambda_permission" "allow_s3_invoke_lambda" {
  statement_id  = "AllowExecutionFromS3"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.image_maker_batch.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = data.terraform_remote_state.s3.outputs.original_image.arn
}

# オリジナルコンテンツ画像へのファイルアップロードをトリガーとしてPUTイベントをLambdaに通知する設定
resource "aws_s3_bucket_notification" "original_image_notification" {
  bucket = data.terraform_remote_state.s3.outputs.original_image.id
  lambda_function {
    lambda_function_arn = aws_lambda_function.image_maker_batch.arn
    events              = ["s3:ObjectCreated:Put"]
  }
  depends_on = [aws_lambda_permission.allow_s3_invoke_lambda]
}
