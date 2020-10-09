# Set Provider as AWS and region
provider "aws" {
    region = var.region
    version = "~> 2"
}

terraform {
  backend "s3" {}
}

locals {
  name_prefix = "${var.common_name_prefix}-validation-${var.environment_name}"
  vpc_prefix = "${var.common_name_prefix}-${var.environment_name}"
}


# Importing existing resources for lambdas to use
data "aws_security_group" "vpc-private" {
  filter {
    name = "tag:Name"
    values = ["${local.vpc_prefix}-vpc-private"]
  }
}

data "aws_subnet" "private-subnet" {
  filter {
    name = "tag:Name"
    values = ["${local.vpc_prefix}-vpc-private-subnet"]
  }
}

data "aws_subnet" "private-subnet2" {
  filter {
    name = "tag:Name"
    values = ["${local.vpc_prefix}-vpc-private-subnet2"]
  }
}

data "aws_lb" "business-layer-lb" {
    name = "${local.vpc_prefix}-bl"
}

data "aws_caller_identity" "current" {}

# input
resource "aws_sqs_queue" "db-export-input" {
  name = "${local.name_prefix}-db-export-input"
  redrive_policy = "{\"deadLetterTargetArn\":\"${aws_sqs_queue.dlq.arn}\",\"maxReceiveCount\":3}"

  tags = merge(
    var.common_tags,
    {
    Name = "${local.name_prefix}-db-export-input",
    "ons:name" = "${local.name_prefix}-db-export-input"
    }
  )
}

# output
resource "aws_sqs_queue" "db-export-output" {
  name = "${local.name_prefix}-db-export-output"
  redrive_policy = "{\"deadLetterTargetArn\":\"${aws_sqs_queue.dlq.arn}\",\"maxReceiveCount\":3}"

  tags = merge(
    var.common_tags,
    {
    Name = "${local.name_prefix}-db-export-output",
    "ons:name" = "${local.name_prefix}-db-export-output"
    }
  )
}