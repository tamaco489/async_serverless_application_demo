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

resource "aws_iam_role" "nautilus_api" {
  name               = "${var.env}-nautilus-api-iam-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_execution_assume_role.json

  inline_policy {
    name   = "${var.env}-nautilus-api-execution-inline-policy"
    policy = data.aws_iam_policy_document.nautilus_api.json
  }

  tags = { Name = "${var.env}-nautilus-api-iam-role" }
}

data "aws_iam_policy_document" "nautilus_api" {
  statement {
    effect = "Allow"
    actions = [
      "sqs:SendMessage"
    ]
    resources = [
      data.terraform_remote_state.sqs.outputs.notification_batch.arn
    ]
  }

  # ListQueuesでは特定のSQSリソースを指定しない
  # DOC: https://docs.aws.amazon.com/ja_jp/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-api-permissions-reference.html
  statement {
    effect = "Allow"
    actions = [
      "sqs:ListQueues",
    ]
    resources = [
      "arn:aws:sqs:ap-northeast-1:${var.account_id}:*"
    ]
  }
}

resource "aws_iam_role_policy_attachment" "nautilus_api_logs" {
  policy_arn = data.terraform_remote_state.lambda.outputs.iam.lambda_logging_policy_arn
  role       = aws_iam_role.nautilus_api.name
}
