# variables.tf

variable "region" {
  description = "The AWS region to deploy to"
  default     = "us-east-1"
}

variable "vpc_name" {
  description = "Name for the VPC"
  default     = "docker_vpc"
}

variable "vpc_cidr_block" {
  description = "CIDR block for the VPC"
  default     = "10.0.0.0/16"
}

variable "subnet_name" {
  description = "Name for the subnet"
  default     = "docker_subnet"
}

variable "subnet_cidr_block" {
  description = "CIDR block for the subnet"
  default     = "10.0.1.0/24"
}

variable "security_group_name" {
  description = "Name for the security group"
  default     = "docker_server_sg"
}

variable "instance_type" {
  description = "EC2 instance type"
  default     = "t3.micro"
}

variable "ami_name_pattern" {
  description = "Pattern to search for the latest Amazon Linux 2 AMI"
  default     = "amzn2-ami-hvm-*"
}

variable "allowed_ports" {
  description = "List of ingress ports to allow"
  type        = list(number)
  default     = [22, 80, 443]
}

variable "egress_ports" {
  description = "List of egress ports to allow"
  type        = list(number)
  default     = [0]
}

variable "cidr_blocks" {
  description = "CIDR blocks to allow access"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "tags" {
  description = "Tags to apply to resources"
  type        = map(string)
  default     = {
    Name = "docker_server"
  }
}
