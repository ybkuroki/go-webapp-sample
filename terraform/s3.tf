resource "aws_s3_bucket" "example" {
  bucket = "terraform-remote-backend-test04"
  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }

  }


  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_dynamodb_table" "statelockfile" {
  name = "statelockfile"

  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"
  #stream_enabled   = true
  #stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "LockID"
    type = "S"
  }
}
