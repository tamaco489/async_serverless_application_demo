output "acm" {
  value = {
    id   = aws_acm_certificate.nautilus.id
    arn  = aws_acm_certificate.nautilus.arn
    name = aws_acm_certificate.nautilus.domain_name
  }
}
