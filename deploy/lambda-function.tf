
resource "aws_lambda_function" "fun_facts_fetcher_lambda" {
  function_name    = var.LAMBDA_FUNCTION_NAME
  filename         = "main.zip"
  handler          = "main"
  source_code_hash = "data.archive_file.zip.output_base64sha256"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10

  environment {
      variables = {
        TABLE_NAME = var.DYNAMO_TABLE_NAME,
        SNS_TOPIC = var.SNS_TOPIC_NAME,
        FUN_FACT_API_URL = "https://asli-fun-fact-api.herokuapp.com/",
      }
  }

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs,
    aws_cloudwatch_log_group.fun_facts_fetcher_lambda_log_group,
    aws_dynamodb_table.fun_facts_fecher_dynamo_table,
    aws_sns_topic.fun_facts
  ]
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
        Effect = "Allow"
        Sid = ""
      }
    ]
  })
}