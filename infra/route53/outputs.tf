output "host_zone" {
  value = {
    id   = aws_route53_zone.nautilus.id,
    name = aws_route53_zone.nautilus.name
  }
}
