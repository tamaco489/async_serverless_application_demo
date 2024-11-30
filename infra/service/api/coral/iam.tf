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
resource "aws_iam_role" "coral_api" {
  name               = "${var.env}-coral-api-iam-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_execution_assume_role.json
  tags               = { Name = "${var.env}-coral-api-iam-role" }
}

# 予め定義しておいた、Lambda共通のポリシー（ログ関連）をアタッチする
resource "aws_iam_role_policy_attachment" "coral_api_logs" {
  policy_arn = data.terraform_remote_state.lambda.outputs.iam.lambda_logging_policy_arn
  role       = aws_iam_role.coral_api.name
}
