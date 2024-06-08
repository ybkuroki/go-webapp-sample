
# outputs.tf

output "instance_public_ip" {
  description = "The public IP address of the EC2 instance"
  value       = aws_instance.docker_server.public_ip
}

output "instance_id" {
  description = "The ID of the EC2 instance"
  value       = aws_instance.docker_server.id
}

output "vpc_id" {
  description = "The ID of the VPC"
  value       = aws_vpc.docker_vpc.id
}

output "subnet_id" {
  description = "The ID of the subnet"
  value       = aws_subnet.docker_subnet.id
}

output "security_group_id" {
  description = "The ID of the security group"
  value       = aws_security_group.docker_server_sg.id
}

output "allowed_ports" {
  description = "List of ingress ports allowed"
  value       = var.allowed_ports
}

output "egress_ports" {
  description = "List of egress ports allowed"
  value       = var.egress_ports
}

output "cidr_blocks" {
  description = "CIDR blocks allowed for access"
  value       = var.cidr_blocks
}
