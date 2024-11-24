# 定期実行するCloudWatch Event Ruleを定義
resource "aws_cloudwatch_event_rule" "reward_batch_schedule" {
  name        = "reward-batch-schedule"
  description = "毎日AM1:30に報酬を付与"
  schedule_expression = "cron(30 16 * * ? *)"  # UTC基準のため時刻設定はUTC+9で設定する必要がある
}

# Event RuleをLambda関数に紐づけるターゲットを定義
resource "aws_cloudwatch_event_target" "reward_batch_lambda_target" {
  rule      = aws_cloudwatch_event_rule.reward_batch_schedule.name
  arn       = aws_lambda_function.reward_batch.arn
}

# LambdaがEventBridgeからのイベントを受け取る権限を付与
resource "aws_lambda_permission" "allow_eventbridge_reward_batch" {
  statement_id  = "AllowExecutionFromEventBridge"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.reward_batch.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.reward_batch_schedule.arn
}