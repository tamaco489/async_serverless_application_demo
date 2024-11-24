resource "aws_lambda_function" "reward_batch" {
  function_name = "${var.env}-reward-batch"
  description   = "報酬管理バッチ"
  role          = aws_iam_role.reward_batch.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.reward_batch.url}:reward_v0.0.0"
  timeout       = 20
  memory_size   = 256

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      API_SERVICE_NAME = "reward_batch"
      API_ENV          = "dev"
      API_PORT         = "8080"
    }
  }

  tags = { Name = "${var.env}-reward-batch" }
}
