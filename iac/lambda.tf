resource "aws_cloudwatch_event_rule" "on-instance-stop" {
  name              = "on-instance-stop"
  description       = "Capture when an EC2 Instance goes into the STOPPED state"
  event_bus_name    = "default"

  event_pattern     = jsonencode({
    detail = {
        state = [
            "stopped",
        ]
    }
    detail-type = [
        "EC2 Instance State-change Notification",
    ]
    source      = [
        "aws.ec2",
    ]
  })
}

resource "aws_cloudwatch_event_target" "start-instance" {
  rule  = aws_cloudwatch_event_rule.on-instance-stop.name
  arn   = aws_lambda_function.start-instance.arn

  retry_policy {
    maximum_event_age_in_seconds  = 3600
    maximum_retry_attempts        = 5
  }

  input_transformer {
    input_paths     = {
      instanceid  = "$.detail.instance-id"
    }
    
    input_template = <<EOF
{
  "InstanceId" : <instanceid>,
  "Status": "START"
}
EOF
  } 
}

resource "aws_lambda_function" "start-instance" {
  filename          = "../dist/${var.filename}"
  function_name     = "StartEC2Instance"
  role              = data.terraform_remote_state.dsf.outputs.role_lambda_arn
  source_code_hash  = filebase64sha256("../dist/${var.filename}")
  runtime           = "go1.x"
  handler           = "main"
  timeout           = 180
}