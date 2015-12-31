package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	var hostname = flag.String("hostname", "None", "Hostname to look up")
	flag.Parse()

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})
	//using example from http://docs.aws.amazon.com/sdk-for-go/api/service/ec2/EC2.html#DescribeInstances-instance_method
	params := &ec2.DescribeInstancesInput{
		DryRun: aws.Bool(false),
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					hostname,
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)

	if len(resp.Reservations) < 1 {
		fmt.Println("No results found")
		return
	}

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
	fmt.Println(*resp.Reservations[0].Instances[0].InstanceId)
}
