resource "aws_cloudwatch_log_group" "reward_batch" {
  name              = "/aws/lambda/${aws_lambda_function.reward_batch.function_name}"
  retention_in_days = 3

  tags = { Name = "reward-batch-logs" }
}
