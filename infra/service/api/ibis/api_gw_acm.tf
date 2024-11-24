# カスタムドメインの定義
resource "aws_apigatewayv2_domain_name" "ibis_api_http" {
  domain_name = "apiv2.${data.terraform_remote_state.route53.outputs.host_zone.name}"

  domain_name_configuration {
    certificate_arn = data.terraform_remote_state.acm.outputs.acm.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

# API Gateway とカスタムドメインのマッピング
resource "aws_apigatewayv2_api_mapping" "ibis_api_mapping" {
  api_id      = aws_apigatewayv2_api.ibis_api_http.id
  domain_name = aws_apigatewayv2_domain_name.ibis_api_http.id
  stage       = aws_apigatewayv2_stage.default_stage.id
}

# Route53 レコードを設定 (カスタムドメインを指す)
resource "aws_route53_record" "api" {
  zone_id = data.terraform_remote_state.route53.outputs.host_zone.id
  name    = "apiv2.${data.terraform_remote_state.route53.outputs.host_zone.name}"
  type    = "A"

  alias {
    name                   = aws_apigatewayv2_domain_name.ibis_api_http.domain_name_configuration.0.target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.ibis_api_http.domain_name_configuration.0.hosted_zone_id
    evaluate_target_health = false
  }
}
