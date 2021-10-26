resource "aws_dynamodb_table" "fun_facts_fecher_dynamo_table" {
  name           = var.DYNAMO_TABLE_NAME
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "lastTimePolled"

  attribute {
    name = "lastTimePolled"
    type = "N"
  }

  tags = {
    Name        = "fun-facts-fetcher-dynamodb-table"
    Environment = "dev"
  }
}