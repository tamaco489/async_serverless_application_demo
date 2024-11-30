resource "aws_acm_certificate" "nautilus" {
  domain_name               = "*.${data.terraform_remote_state.route53.outputs.host_zone.name}"
  subject_alternative_names = [data.terraform_remote_state.route53.outputs.host_zone.name]
  validation_method         = "DNS"

  lifecycle {
    create_before_destroy = true
  }

  tags = { Name = "${local.fqn}-acm" }
}

# API Gatewayのカスタムドメインの定義（画像処理サービスAPI向け）
resource "aws_apigatewayv2_domain_name" "ibis_api_http_v2" {
  domain_name = "apiv2.${data.terraform_remote_state.route53.outputs.host_zone.name}"

  domain_name_configuration {
    certificate_arn = aws_acm_certificate.nautilus.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

# API V2向けのRoute53 Aレコードを設定（画像処理サービスAPI向け）
resource "aws_route53_record" "ibis_api_http_v2" {
  zone_id = data.terraform_remote_state.route53.outputs.host_zone.id
  name    = "apiv2.${data.terraform_remote_state.route53.outputs.host_zone.name}"
  type    = "A"

  alias {
    name                   = aws_apigatewayv2_domain_name.ibis_api_http_v2.domain_name_configuration.0.target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.ibis_api_http_v2.domain_name_configuration.0.hosted_zone_id
    evaluate_target_health = false
  }
}
