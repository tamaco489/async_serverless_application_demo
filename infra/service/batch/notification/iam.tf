# AWS Lambdaが指定されたロールを引き受けるためのポリシードキュメントを生成する
# そのロールに付与された権限を使用できるようにするために必要な設定
# これにより、Lambda関数が特定のAWSリソースにアクセスするために必要な権限を持つことができるようになる
data "aws_iam_policy_document" "lambda_execution_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

# IAM Roleの作成
resource "aws_iam_role" "notification_batch" {
  name               = "${var.env}-notification-batch-iam-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_execution_assume_role.json
  tags               = { Name = "notification-batch-iam-role" }
}

# SQSへのアクセス権を定義するポリシードキュメントを生成する
data "aws_iam_policy_document" "notification_batch" {
  statement {
    effect = "Allow"
    actions = [
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes"
    ]
    resources = [
      data.terraform_remote_state.sqs.outputs.notification_batch.arn
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "sqs:SendMessage"
    ]
    resources = [
      data.terraform_remote_state.sqs.outputs.notification_batch_dlq.arn
    ]
  }
}

# SQSへのアクセス権を定義するポリシードキュメントをIAM Policyとして定義し、IAM Roleに関連付ける
resource "aws_iam_role_policy" "notification_batch" {
  name   = "${var.env}-notification-batch-execution-policy"
  role   = aws_iam_role.notification_batch.id
  policy = data.aws_iam_policy_document.notification_batch.json
}

# 予め定義しておいた、Lambda共通のポリシー（ログ関連）をアタッチする
resource "aws_iam_role_policy_attachment" "notification_batch_logs" {
  policy_arn = data.terraform_remote_state.lambda.outputs.iam.lambda_logging_policy_arn
  role       = aws_iam_role.notification_batch.name
}
