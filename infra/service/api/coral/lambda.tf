resource "aws_lambda_function" "coral_api" {
  function_name = "${var.env}-coral-api"
  description   = "ユーザ情報登録APIサーバ"
  role          = aws_iam_role.coral_api.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.coral_api.url}:coral_v0.0.0"
  timeout       = 20
  memory_size   = 256

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      API_SERVICE_NAME = "coral-api"
      API_ENV          = "dev"
      API_PORT         = "8080"
    }
  }

  tags = { Name = "${var.env}-coral-api" }
}
