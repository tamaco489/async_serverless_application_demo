resource "aws_cloudwatch_log_group" "nautilus_api" {
  name              = "/aws/lambda/${aws_lambda_function.nautilus_api.function_name}"
  retention_in_days = 3

  tags = { Name = "${var.env}-nautilus-api-logs" }
}
