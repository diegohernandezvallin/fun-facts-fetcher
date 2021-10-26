resource "aws_cloudwatch_event_rule" "daily" {
  name                = "daily"
  description         = "Fires daily at midday"
  //schedule_expression = "cron(30 12 * * ? *)" uncomment after testing
  schedule_expression = "rate(1 minute)" //remove once deploy to prod
}

resource "aws_cloudwatch_event_target" "fun_facts_fetcher_lambda_daily" {
    rule      = aws_cloudwatch_event_rule.daily.name
    target_id = "fun_facts_fetcher_lambda"
    arn       = aws_lambda_function.fun_facts_fetcher_lambda.arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_fun_facts_fetcher_lambda" {
    statement_id = "AllowExecutionFromCloudWatch"
    action = "lambda:InvokeFunction"
    function_name = var.LAMBDA_FUNCTION_NAME
    principal = "events.amazonaws.com"
    source_arn = aws_cloudwatch_event_rule.daily.arn
}