
# main.tf

resource "aws_internet_gateway" "docker_igw" {
  vpc_id = aws_vpc.docker_vpc.id
}

resource "aws_route_table" "docker_route_table" {
  vpc_id = aws_vpc.docker_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.docker_igw.id
  }
}

resource "aws_route_table_association" "docker_route_table_association" {
  subnet_id      = aws_subnet.docker_subnet.id
  route_table_id = aws_route_table.docker_route_table.id
}

resource "aws_vpc" "docker_vpc" {
  cidr_block = var.vpc_cidr_block
  tags = {
    Name = var.vpc_name
  }
}

resource "aws_subnet" "docker_subnet" {
  vpc_id                  = aws_vpc.docker_vpc.id
  cidr_block              = var.subnet_cidr_block
  availability_zone       = "${var.region}a"
  map_public_ip_on_launch = true # Add this line

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
  ami                         = data.aws_ami.aws_linux_2_latest.id
  instance_type               = var.instance_type
  vpc_security_group_ids      = [aws_security_group.docker_server_sg.id]
  subnet_id                   = aws_subnet.docker_subnet.id
  associate_public_ip_address = true
  key_name                    = "default-ec2"


  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo amazon-linux-extras install docker -y",
      "sudo service docker start",
      "sudo usermod -a -G docker ec2-user",
      # Docker login
      "echo ${var.DOCKER_PASSWORD} | sudo docker login --username ${var.DOCKER_USERNAME} --password-stdin"
    ]
  }

  connection {
    type        = "ssh"
    host        = self.public_ip
    user        = "ec2-user"
    private_key = file(var.ssh_private_key_path)
  }

  tags = var.tags
}
