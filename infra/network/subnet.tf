resource "aws_subnet" "public" {
  for_each                = var.public_subnet
  vpc_id                  = aws_vpc.vpc.id
  cidr_block              = each.value["cidr"]
  availability_zone       = "${var.region}${each.value["az"]}"
  map_public_ip_on_launch = true

  tags = { Name = "${local.fqn}-public-${each.key}" }
}

# NOTE: privateリソースは現時点では使用する予定はないため一旦コメントアウト
# resource "aws_subnet" "private" {
#   for_each                = var.private_subnet
#   vpc_id                  = aws_vpc.vpc.id
#   cidr_block              = each.value["cidr"]
#   availability_zone       = "${var.region}${each.value["az"]}"
#   map_public_ip_on_launch = false

#   tags = { Name = "${local.fqn}-private-${each.key}" }
# }
