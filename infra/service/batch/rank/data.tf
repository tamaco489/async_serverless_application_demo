data "terraform_remote_state" "ecr" {
  backend = "s3"
  config = {
    bucket = "dev-nautilus-tfstate"
    key    = "ecr/terraform.tfstate"
  }
}

data "terraform_remote_state" "lambda" {
  backend = "s3"
  config = {
    bucket = "dev-nautilus-tfstate"
    key    = "lambda/terraform.tfstate"
  }
}
