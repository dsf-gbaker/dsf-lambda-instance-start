package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2SetStatusEvent struct {
	InstanceId string `json:"InstanceId"`
	Status     string `json:"Status"`
}

type Response struct {
	Message string `json:"Message:"`
}

func StartInstance(ev EC2SetStatusEvent) string {

	res := ""
	region := os.Getenv("AWS_REGION")

	// Load session from shared config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}

	// Create new EC2 client
	svc := ec2.New(sess)

	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(ev.InstanceId),
		},
		DryRun: aws.Bool(true),
	}

	result, err := svc.StartInstances(input)
	awsErr, ok := err.(awserr.Error)

	if ok && awsErr.Code() == "DryRunOperation" {
		// Let's now set dry run to be false. This will allow us to start the instances
		input.DryRun = aws.Bool(false)
		result, err = svc.StartInstances(input)
		if err != nil {
			res = fmt.Sprintf("Error: %s", err)
		} else {
			res = fmt.Sprintf("Success: %s", result.GoString())
		}

	} else { // This could be due to a lack of permissions
		res = fmt.Sprintf("Error: %s", err)
	}

	return res
}

func StopInstance(ev EC2SetStatusEvent) string {

	res := "There was an "
	region := os.Getenv("AWS_REGION")

	// Load session from shared config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}

	// Create new EC2 client
	svc := ec2.New(sess)

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(ev.InstanceId),
		},
		DryRun: aws.Bool(true),
	}

	result, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)

	if ok && awsErr.Code() == "DryRunOperation" {
		// Let's now set dry run to be false. This will allow us to start the instances
		input.DryRun = aws.Bool(false)
		result, err = svc.StopInstances(input)
		if err != nil {
			res = fmt.Sprintf("Error: %s", err)
		} else {
			res = fmt.Sprintf("Success: %s", result.GoString())
		}

	} else { // This could be due to a lack of permissions
		res = fmt.Sprintf("Error: %s", err)
	}

	return res
}

func HandleEvent(ev EC2SetStatusEvent) (Response, error) {

	res := ""

	fmt.Printf("Received instance id %s and desired status of %s\n", ev.InstanceId, ev.Status)

	switch ev.Status {
	case "START":
		res = StartInstance(ev)
	case "STOP":
		res = StopInstance(ev)
	case "REBOOT":
		fallthrough
	default:
		res = fmt.Sprintf("No support for desired status %s", ev.Status)
	}

	fmt.Printf("Result: %s\n", res)
	return Response{Message: res}, nil
}

func main() {
	lambda.Start(HandleEvent)
}
