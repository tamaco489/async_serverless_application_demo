resource "aws_lambda_function" "notification_batch" {
  function_name = "${var.env}-notification-batch"
  description   = "通知サービスバッチ"
  role          = aws_iam_role.notification_batch.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.notification_batch.url}:notification_v0.0.0"
  timeout       = 20
  memory_size   = 256

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      API_SERVICE_NAME = "notification-batch"
      API_ENV          = "dev"
      SQS_DLQ_URL      = data.terraform_remote_state.sqs.outputs.notification_batch_dlq.id
    }
  }

  tags = { Name = "${var.env}-notification-batch" }
}

# Lambdaのイベントソースマッピングの設定
# SQSキューからのメッセージ受信をトリガーに、Lambadaを起動するために必要な設定
resource "aws_lambda_event_source_mapping" "notification_batch" {
  event_source_arn = data.terraform_remote_state.sqs.outputs.notification_batch.arn
  function_name    = aws_lambda_function.notification_batch.function_name

  # Queueのトリガー設定を有効化
  enabled = true

  # Lambdaは20秒間待機し、その間に最大5件のメッセージを集める
  # 待機時間が経過する前に5件集まった場合即座に処理を開始するが、20秒経過したら集まったメッセージ数に関わらず処理が始まる
  maximum_batching_window_in_seconds = 20
  batch_size                         = 5


  # 送信されるメッセージにおけるフィルター処理を定義することができる
  # DLQへの受け渡しの挙動を検証したいためフィルタリング設定は無効化する
  # filter_criteria {
  #   filter {
  #     pattern = jsonencode({
  #       body = {
  #         status = ["COMPLETED"] # "status"が"COMPLETED"の場合にのみトリガー
  #       }
  #     })
  #   }
  # }
}
