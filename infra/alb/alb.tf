resource "aws_alb" "alb" {
  name               = "${local.fqn}-alb"
  internal           = false
  load_balancer_type = "application"
  subnets            = data.terraform_remote_state.network.outputs.vpc.public_subnet_ids
  security_groups    = [aws_security_group.alb.id]

  # NOTE: ALB向けログ格納バケットを作成した後に有効化
  # access_logs {
  #   enabled = true
  #   bucket  = ""
  #   prefix  = "alb/nautilus-api/access_log"
  # }

  tags = { Name = "${local.fqn}-alb" }
}

resource "aws_route53_record" "api" {
  zone_id = data.terraform_remote_state.route53.outputs.host_zone.id
  name    = "api.${data.terraform_remote_state.route53.outputs.host_zone.name}"
  type    = "A"

  alias {
    name                   = aws_alb.alb.dns_name
    zone_id                = aws_alb.alb.zone_id
    evaluate_target_health = true
  }
}

resource "aws_alb_listener" "http_listener" {
  load_balancer_arn = aws_alb.alb.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"
    redirect {
      port        = 443
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_alb_listener" "https_listener" {
  load_balancer_arn = aws_alb.alb.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-TLS13-1-2-2021-06"
  certificate_arn   = data.terraform_remote_state.acm.outputs.acm.arn

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      status_code  = "404"
    }
  }
}
