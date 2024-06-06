output "aws_security_group_http_server_details" {
  value = aws_security_group.http_server_sg
}

output "http_server_public_dns" {
  value = aws_instance.http_server.public_dns
}

output "instance_key_name" {
  value = aws_instance.http_server.key_name
}

output "instance_public_ip" {
  value = aws_instance.http_server.public_ip
}
