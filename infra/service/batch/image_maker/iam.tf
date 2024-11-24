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
resource "aws_iam_role" "image_maker_batch" {
  name               = "${var.env}-image-maker-batch-iam-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_execution_assume_role.json
  tags               = { Name = "${var.env}-image-maker-batch-iam-role" }
}

# S3へのアクセス権を定義するポリシードキュメントを定義する
data "aws_iam_policy_document" "image_maker_batch" {
  # オリジナル画像コンテンツのS3バケット参照権限
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
    ]
    resources = ["${data.terraform_remote_state.s3.outputs.original_image.arn}/*"]
  }

  # サムネイル画像コンテンツを配置するS3バケットアップロード権限
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
    ]
    resources = ["${data.terraform_remote_state.s3.outputs.thumbnail_image.arn}/*"]
  }
}

# S3アクセス権限をIAM Role に付与する
resource "aws_iam_role_policy" "image_maker_batch" {
  name   = "${var.env}-image-maker-batch-lambda-s3-access-policy"
  role   = aws_iam_role.image_maker_batch.id
  policy = data.aws_iam_policy_document.image_maker_batch.json
}

# 予め定義しておいた、Lambda共通のポリシー（ログ関連）をアタッチする
resource "aws_iam_role_policy_attachment" "notification_batch_logs" {
  policy_arn = data.terraform_remote_state.lambda.outputs.iam.lambda_logging_policy_arn
  role       = aws_iam_role.image_maker_batch.name
}
