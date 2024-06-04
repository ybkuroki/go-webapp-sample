# main.tf

resource "aws_vpc" "docker_vpc" {
  cidr_block = var.vpc_cidr_block
  tags = {
    Name = var.vpc_name
  }
}

resource "aws_subnet" "docker_subnet" {
  vpc_id            = aws_vpc.docker_vpc.id
  cidr_block        = var.subnet_cidr_block
  availability_zone = "${var.region}a"

  tags = {
    Name = var.subnet_name
  }
}

resource "aws_security_group" "docker_server_sg" {
  name   = var.security_group_name
  vpc_id = aws_vpc.docker_vpc.id

  dynamic "ingress" {
    for_each = var.allowed_ports
    content {
      from_port   = ingress.value
      to_port     = ingress.value
      protocol    = "tcp"
      cidr_blocks = var.cidr_blocks
    }
  }

  dynamic "egress" {
    for_each = var.egress_ports
    content {
      from_port   = egress.value
      to_port     = egress.value
      protocol    = -1
      cidr_blocks = var.cidr_blocks
    }
  }

  tags = {
    Name = var.security_group_name
  }
}

resource "aws_instance" "docker_server" {
  ami                    = data.aws_ami.aws_linux_2_latest.id
  instance_type          = var.instance_type
  vpc_security_group_ids = [aws_security_group.docker_server_sg.id]
  subnet_id              = aws_subnet.docker_subnet.id

  tags = var.tags
}
