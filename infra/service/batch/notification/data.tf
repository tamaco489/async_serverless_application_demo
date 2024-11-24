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

data "terraform_remote_state" "sqs" {
  backend = "s3"
  config = {
    bucket = "${var.env}-nautilus-tfstate"
    key    = "sqs/terraform.tfstate"
  }
}
