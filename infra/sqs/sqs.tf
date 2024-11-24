# 注文受付サービス（Lambda）より送信されるキュー
# このSQSがキューを受信したことをトリガーにメッセージ通知サービス（Lambda）を起動する
resource "aws_sqs_queue" "to_notifications_queue" {
  name                       = "${var.env}-to-notifications-queue"
  fifo_queue                 = false  # 標準Queueとして定義（順序を考慮しない）
  max_message_size           = 262144 # 最大メッセージサイズ (256 KB)
  visibility_timeout_seconds = 20     # キューの可視性タイムアウト。秒単位で設定。（左記は20秒で設定、指定した期間中は他のコンシューマーはキューを参照することができない）
  message_retention_seconds  = 345600 # メッセージの保持時間、秒単位で設定（左記は4日間で設定）
  delay_seconds              = 30     # キュー内の全メッセージの配送を遅延させる時間、秒単位で設定。（左記は30秒で設定、つまりキューに送信してもLambda30秒間はこのキューを見ることができない。※0の場合は即時配信となる）

  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.to_notifications_dlq.arn # DLQのARNを指定
    maxReceiveCount     = 5                                      # リトライ回数を5回に設定
  })

  tags = { Name = "${var.env}-to-notifications-queue" }
}

# 注文受付サービス --> メッセージ通知サービスのキューイング処理失敗時に退避させるDLQ
resource "aws_sqs_queue" "to_notifications_dlq" {
  name                      = "${var.env}-to-notification-dlq"
  max_message_size          = 262144
  message_retention_seconds = 604800 # メッセージ保持期間を7日間に設定

  tags = {
    Name = "${var.env}-to-notification-dlq"
  }
}
