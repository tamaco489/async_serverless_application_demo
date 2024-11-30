output "nautilus_api" {
  value = {
    arn  = aws_ecr_repository.nautilus_api.arn
    id   = aws_ecr_repository.nautilus_api.id
    name = aws_ecr_repository.nautilus_api.name
    url  = aws_ecr_repository.nautilus_api.repository_url
  }
}

output "ibis_api" {
  value = {
    arn  = aws_ecr_repository.ibis_api.arn
    id   = aws_ecr_repository.ibis_api.id
    name = aws_ecr_repository.ibis_api.name
    url  = aws_ecr_repository.ibis_api.repository_url
  }
}

output "coral_api" {
  value = {
    arn  = aws_ecr_repository.coral_api.arn
    id   = aws_ecr_repository.coral_api.id
    name = aws_ecr_repository.coral_api.name
    url  = aws_ecr_repository.coral_api.repository_url
  }
}

output "image_maker_batch" {
  value = {
    arn  = aws_ecr_repository.image_maker_batch.arn
    id   = aws_ecr_repository.image_maker_batch.id
    name = aws_ecr_repository.image_maker_batch.name
    url  = aws_ecr_repository.image_maker_batch.repository_url
  }
}

output "notification_batch" {
  value = {
    arn  = aws_ecr_repository.notification_batch.arn
    id   = aws_ecr_repository.notification_batch.id
    name = aws_ecr_repository.notification_batch.name
    url  = aws_ecr_repository.notification_batch.repository_url
  }
}

output "rank_batch" {
  value = {
    arn  = aws_ecr_repository.rank_batch.arn
    id   = aws_ecr_repository.rank_batch.id
    name = aws_ecr_repository.rank_batch.name
    url  = aws_ecr_repository.rank_batch.repository_url
  }
}

output "reward_batch" {
  value = {
    arn  = aws_ecr_repository.reward_batch.arn
    id   = aws_ecr_repository.reward_batch.id
    name = aws_ecr_repository.reward_batch.name
    url  = aws_ecr_repository.reward_batch.repository_url
  }
}
