output "mailer_config" {
  value = {
    name = aws_secretsmanager_secret.mailer_config.name
    arn  = aws_secretsmanager_secret.mailer_config.arn
  }
}