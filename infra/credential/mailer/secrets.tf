resource "aws_secretsmanager_secret" "mailer_config" {
  # NOTE: 一旦ローカル環境と同じキー名で指定
  name        = "supply-chain-demo/dev/notifications/mailer-config"
  description = "Mail関連の秘匿情報を取り扱うSecret Manager"
}

resource "aws_secretsmanager_secret_version" "mailer_config_secret" {
  secret_id     = aws_secretsmanager_secret.mailer_config.id
  secret_string = jsonencode(var.mailer_config)
}
