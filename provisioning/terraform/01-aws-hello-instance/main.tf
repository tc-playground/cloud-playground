provider "aws" {
    version = "2.12.0"
    region = "eu-west-2"
    profile = "trjl-terraform"
}

# ami-ubuntu-18.04-1.13.2-00-1548795114
resource "aws_instance" "hello-instance" {
    ami = "ami-0521c4dc863d418da"
    instance_type = "t2.micro"
}