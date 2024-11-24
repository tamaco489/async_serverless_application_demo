resource "aws_cloudwatch_log_group" "notification_batch" {
  name              = "/aws/lambda/${aws_lambda_function.notification_batch.function_name}"
  retention_in_days = 3

  tags = { Name = "${var.env}-notification-batch-logs" }
}
