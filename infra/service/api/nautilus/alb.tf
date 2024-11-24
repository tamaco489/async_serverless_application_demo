# =================================================================
# ターゲットグループの設定
# =================================================================
resource "aws_alb_target_group" "nautilus_api" {
  name        = "${var.env}-nautilus-api"
  target_type = "lambda"

  # ALBのマルチバリューヘッダーの有効化
  # NOTE: ALB→Lambda構成において、同一クエリパラメータで複数の値を指定してリクエストする場合に必要（falseの場合、複数指定してもALBで握り潰され1つしか設定されないため）
  # DOC: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lb_target_group
  lambda_multi_value_headers_enabled = true

  tags = { Name = "${var.env}-nautilus-api" }
}

# =================================================================
# リスナールールの設定
# =================================================================
resource "aws_alb_listener_rule" "nautilus_api" {
  listener_arn = data.terraform_remote_state.alb.outputs.alb.https_listener_arn

  # 他リスナールールとの重複を避けるため優先度を設定
  priority = 100

  # ALBがリクエストを受け取った時の振る舞いを定義
  # 「forward」アクションとして、指定したターゲットグループにリクエストを転送する
  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.nautilus_api.arn
  }

  condition {
    # リクエストされたパスが、以下の指定したパスに合致している場合にのみこのルールを適用する
    # e.g. https://api.<domain-name>/nautilus/v1/products/purchase
    path_pattern {
      values = ["/nautilus/*"]
    }
  }

  tags = { Name = "${var.env}-nautilus-api" }
}

# =================================================================
# ALBのターゲットグループにLambdaをアタッチ
# =================================================================
resource "aws_lb_target_group_attachment" "nautilus_api" {
  # ALBターゲットグループ
  target_group_arn = aws_alb_target_group.nautilus_api.arn

  # Lambda関数
  target_id = aws_lambda_function.nautilus_api.arn
}
