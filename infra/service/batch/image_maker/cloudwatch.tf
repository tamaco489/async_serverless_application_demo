resource "aws_cloudwatch_log_group" "image_maker_batch" {
  name              = "/aws/lambda/${aws_lambda_function.image_maker_batch.function_name}"
  retention_in_days = 3

  tags = { Name = "${var.env}-image-maker-batch-logs" }
}
