output "notification_batch" {
  value = {
    id  = aws_sqs_queue.to_notifications_queue.id,
    arn = aws_sqs_queue.to_notifications_queue.arn
  }
}

output "notification_batch_dlq" {
  value = {
    id  = aws_sqs_queue.to_notifications_dlq.id,
    arn = aws_sqs_queue.to_notifications_dlq.arn
  }
}
