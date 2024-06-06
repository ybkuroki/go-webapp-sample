variable "aws_key_pair" {
  default = "~/aws/aws_keys/default-ec2.pem"
}

variable "DOCKER_USERNAME" {
  description = "Username for Docker Hub login"
}

variable "DOCKER_PASSWORD" {
  description = "Password for Docker Hub login"
}

variable "ssh_private_key_path" {
  description = "Path to the SSH private key"
  #default     = "~/.ssh/default-ec2.pem"
}
