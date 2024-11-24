resource "aws_security_group" "nautilus_api" {
  name        = "${var.env}-nautilus-api-sg"
  description = "Nautilus API Security Group"
  vpc_id      = data.terraform_remote_state.network.outputs.vpc.id
  tags        = { Name = "${var.env}-nautilus-api-sg" }
}

# Nautilus API におけるインバウンド通信のトラフィックルールを定義
resource "aws_vpc_security_group_ingress_rule" "nautilus_api" {
  security_group_id = aws_security_group.nautilus_api.id
  description       = "Allow TCP traffic on port 8080 from ALB security group"
  from_port         = 8080
  to_port           = 8080
  ip_protocol       = "TCP"

  # ALBセキュリティグループからのトラフィックのみを許可する
  referenced_security_group_id = data.terraform_remote_state.alb.outputs.alb.security_group_id
  # [補足]
  # Nautilus API に接続しようとするトラフィックが、指定された ALB セキュリティグループを経由している場合のみ通信を許可し、
  # 他のリソースやネットワークから直接APIにアクセスできないようにする)

  tags = { Name = "${var.env}-nautilus-api-sg-ingress" }
}

# Nautilus API におけるアウトバウンド通信のトラフィックルールを定義
resource "aws_vpc_security_group_egress_rule" "nautilus_api" {
  security_group_id = aws_security_group.nautilus_api.id
  description       = "Allow all outbound traffic"
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
  tags              = { Name = "${var.env}-nautilus-api-sg-egress" }
}
