resource "aws_lambda_function" "coral_api" {
  function_name = "${var.env}-coral-api"
  description   = "ユーザ情報管理APIサーバ"
  role          = aws_iam_role.coral_api.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.coral_api.url}:coral_v0.0.0"
  timeout       = 20
  memory_size   = 128

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      API_SERVICE_NAME = "coral-api"
      API_ENV          = "stg"
      API_PORT         = "8080"
    }
  }

  tags = { Name = "${var.env}-coral-api" }
}

# Lambda 関数が API Gateway から呼び出せるようにする権限
resource "aws_lambda_permission" "api_gateway_invoke" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.coral_api.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.coral_api_http.execution_arn}/*/*"
}
