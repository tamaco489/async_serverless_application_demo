resource "aws_route53_zone" "nautilus" {
  name    = var.domain
  comment = "nautilus API サーバの検証で利用"
  tags = { Name = "${local.fqn}-route53-zone" }
}
