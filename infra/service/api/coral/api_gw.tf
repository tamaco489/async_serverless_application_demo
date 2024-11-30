# Lambda関数をターゲットとするAPI Gateway (HTTP API) を作成
resource "aws_apigatewayv2_api" "coral_api_http" {
  name          = "${var.env}-coral-http-api"
  description   = "ユーザ情報管理 API Gateway (HTTP API)"
  protocol_type = "HTTP"
  tags = {
    Name = "${var.env}-coral-api-gateway"
  }
}

# Lambda と API Gateway を統合するための統合リソース
resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id                 = aws_apigatewayv2_api.coral_api_http.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.coral_api.invoke_arn
  payload_format_version = "2.0"
}

# Lambda 関数に対するルーティング設定 (/coral/v3/{proxy+})
resource "aws_apigatewayv2_route" "coral_api_route" {
  api_id    = aws_apigatewayv2_api.coral_api_http.id
  route_key = "ANY /coral/v3/{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

# ステージ (デフォルトステージ)
resource "aws_apigatewayv2_stage" "default_stage" {
  api_id      = aws_apigatewayv2_api.coral_api_http.id
  name        = "$default"
  auto_deploy = true
}

# API Gateway とカスタムドメインのマッピング
resource "aws_apigatewayv2_api_mapping" "coral_api_mapping" {
  api_id      = aws_apigatewayv2_api.coral_api_http.id
  domain_name = data.terraform_remote_state.acm.outputs.coral_apigatewayv2_domain_name.id
  stage       = aws_apigatewayv2_stage.default_stage.id
}
