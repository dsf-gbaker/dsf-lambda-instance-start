# dsf-launch-ec2
A GO lambda function for launching an ec2 instance

## PreRequisite(s)

- go v1.17.x
- *build-lambda-zip* tool
  - `go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip`
- Terraform

## Compile & Deploy Steps

- Open Powershell prompt in the `src` folder
  - `./build.ps1`
- In powershell, navigate to the `iac` folder
  - `terraform apply`