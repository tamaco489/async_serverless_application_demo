resource "aws_lambda_function" "nautilus_api" {
  function_name = "${var.env}-nautilus-api"
  description   = "発注処理APIサーバ"
  role          = aws_iam_role.nautilus_api.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.nautilus_api.url}:nautilus_v0.0.0"
  timeout       = 20
  memory_size   = 256

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      API_SERVICE_NAME            = "nautilus-api"
      API_ENV                     = "dev"
      API_PORT                    = "8080"
      SQS_NOTIFICATIONS_QUEUE_URL = data.terraform_remote_state.sqs.outputs.notification_batch.id
    }
  }

  tags = { Name = "${var.env}-nautilus-api" }
}

# NOTE:「aws_lambda_permission」は特定のAWSリソースが、指定したLambda関数を呼び出すことを許可するための権限を設定する
# DOC: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission
resource "aws_lambda_permission" "alb" {
  # 権限の識別子を指定 (リソース内でユニークな名称にしておく必要がある)
  statement_id = "AllowExecutionFromALB"

  # 許可するアクションを指定 ('lambda:InvokeFunction'を設定することでLambda関数を呼び出すことができるようになる)
  action = "lambda:InvokeFunction"

  # アクセスを許可するLambda関数名
  function_name = aws_lambda_function.nautilus_api.function_name

  # 権限を付与するエンティティを指定
  # 'elasticloadbalancing.amazonaws.com'を指定することでALBのみがこのLambda関数を呼び出すことが可能になる
  # 補足として、仮に「principal = "*"」と設定することも可能だが、その場合全てのリソースからアクセスを許可することになりセキュリティリスクが高まってしまうため、可能な限り最小限のアクセス権限を設定することを推奨されている
  # DOC: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_target_group_attachment#lambda-target
  principal = "elasticloadbalancing.amazonaws.com"

  # 別途、alb.tfで指定したALBターゲットグループのARNを指定
  source_arn = aws_alb_target_group.nautilus_api.arn
}
