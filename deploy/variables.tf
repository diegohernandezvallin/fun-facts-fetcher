
variable "AWS_REGION" {
  default = "us-east-1"
}

variable "LAMBDA_FUNCTION_NAME" {
  default = "fun-facts-fetcher-lambda"
}

variable "DYNAMO_TABLE_NAME" {
  default = "fun-facts-fecther"
}

variable "SNS_TOPIC_NAME" {
  default = "fun-facts-fetcher-topic"
}
