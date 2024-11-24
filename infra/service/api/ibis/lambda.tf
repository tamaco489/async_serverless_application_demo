resource "aws_lambda_function" "ibis_api" {
  function_name = "${var.env}-ibis-api"
  description   = "画像処理APIサーバ"
  role          = aws_iam_role.ibis_api.arn
  package_type  = "Image"
  image_uri     = "${data.terraform_remote_state.ecr.outputs.ibis_api.url}:ibis_v0.0.0"
  timeout       = 20
  memory_size   = 256

  lifecycle {
    ignore_changes = [image_uri]
  }

  environment {
    variables = {
      API_SERVICE_NAME              = "ibis-api"
      API_ENV                       = "dev"
      API_PORT                      = "8080"
      S3_ORIGINAL_IMAGE_BUCKET_NAME = data.terraform_remote_state.s3.outputs.original_image.id
    }
  }

  tags = { Name = "${var.env}-ibis-api" }
}

# Lambda 関数が API Gateway から呼び出せるようにする権限
resource "aws_lambda_permission" "api_gateway_invoke" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.ibis_api.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.ibis_api_http.execution_arn}/*/*"
}
