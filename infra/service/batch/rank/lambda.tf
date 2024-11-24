resource "aws_lambda_function" "rank_batch" {
  function_name = "${var.env}-rank-batch"
  description   = "ランキング集計バッチ"
  role          = aws_iam_role.rank_batch.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.rank_batch.url}:rank_v0.0.0"
  timeout       = 20
  memory_size   = 256

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      API_SERVICE_NAME = "rank_batch"
      API_ENV          = "dev"
      API_PORT         = "8080"
    }
  }

  tags = { Name = "${var.env}-rank-batch" }
}
