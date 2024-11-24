resource "aws_cloudwatch_log_group" "ibis_api" {
  name              = "/aws/lambda/${aws_lambda_function.ibis_api.function_name}"
  retention_in_days = 3

  tags = { Name = "${var.env}-ibis-api-logs" }
}
