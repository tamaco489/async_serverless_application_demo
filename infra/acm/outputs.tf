output "acm" {
  value = {
    id   = aws_acm_certificate.nautilus.id
    arn  = aws_acm_certificate.nautilus.arn
    name = aws_acm_certificate.nautilus.domain_name
  }
}

output "apigatewayv2_domain_name" {
  value = {
    id              = aws_apigatewayv2_domain_name.api_http_v2.id
    domain_name     = aws_apigatewayv2_domain_name.api_http_v2.domain_name
    endpoint_type   = aws_apigatewayv2_domain_name.api_http_v2.domain_name_configuration[0].endpoint_type
    security_policy = aws_apigatewayv2_domain_name.api_http_v2.domain_name_configuration[0].security_policy
    certificate_arn = aws_apigatewayv2_domain_name.api_http_v2.domain_name_configuration[0].certificate_arn
  }
  description = "Details of the API Gateway v2 custom domain name configuration."
}
