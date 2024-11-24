resource "aws_acm_certificate" "nautilus" {
  domain_name               = "*.${data.terraform_remote_state.route53.outputs.host_zone.name}"
  subject_alternative_names = [data.terraform_remote_state.route53.outputs.host_zone.name]
  validation_method         = "DNS"

  lifecycle {
    create_before_destroy = true
  }

  tags = { Name = "${local.fqn}-acm" }
}
