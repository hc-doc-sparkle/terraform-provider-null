resource "aws_instance" "cluster" {
  count = 3

  # ...
}