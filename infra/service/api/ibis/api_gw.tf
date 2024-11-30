# Lambda関数をターゲットとするAPI Gateway (HTTP API) を作成
resource "aws_apigatewayv2_api" "ibis_api_http" {
  name          = "${var.env}-ibis-http-api"
  description   = "画像処理APIのAPI Gateway (HTTP API)"
  protocol_type = "HTTP"
  tags = {
    Name = "${var.env}-ibis-api-gateway"
  }
}

# Lambda と API Gateway を統合するための統合リソース
resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id                 = aws_apigatewayv2_api.ibis_api_http.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.ibis_api.invoke_arn
  payload_format_version = "2.0"
}

# Lambda 関数に対するルーティング設定 (/ibis/v2/{proxy+})
resource "aws_apigatewayv2_route" "ibis_api_route" {
  api_id    = aws_apigatewayv2_api.ibis_api_http.id
  route_key = "ANY /ibis/v2/{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

# ステージ (デフォルトステージ)
resource "aws_apigatewayv2_stage" "default_stage" {
  api_id      = aws_apigatewayv2_api.ibis_api_http.id
  name        = "$default"
  auto_deploy = true
}

# API Gateway とカスタムドメインのマッピング
resource "aws_apigatewayv2_api_mapping" "ibis_api_mapping" {
  api_id      = aws_apigatewayv2_api.ibis_api_http.id
  domain_name = data.terraform_remote_state.acm.outputs.apigatewayv2_domain_name.id
  stage       = aws_apigatewayv2_stage.default_stage.id
}
