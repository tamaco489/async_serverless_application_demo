resource "aws_cloudwatch_log_group" "coral_api" {
  name              = "/aws/lambda/${aws_lambda_function.coral_api.function_name}"
  retention_in_days = 3

  tags = { Name = "${var.env}-coral-api-logs" }
}
