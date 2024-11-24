resource "aws_cloudwatch_log_group" "rank_batch" {
  name              = "/aws/lambda/${aws_lambda_function.rank_batch.function_name}"
  retention_in_days = 3

  tags = { Name = "rank-batch-logs" }
}
