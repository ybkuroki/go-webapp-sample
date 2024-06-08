# providers.tf

terraform {
  required_version = ">= 1.0.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.74.0"
    }
  }

  backend "s3" {
    bucket = "terraform-remote-backend-test04"
    key    = "global/mystatefile/terraform.tfstate"
    region = "us-east-1"
    # Optional: Uncomment this line if you want to enable DynamoDB table for state locking
    dynamodb_table = "statelockfile"
    encrypt        = true
  }
}

provider "aws" {
  region = var.region
}
