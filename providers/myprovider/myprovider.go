package myprovider

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type MyProvider struct {
	EC2 *ec2.EC2
}

func NewMyProvider(region string) *MyProvider {
	fmt.Printf("Creating new AWS session for region %s...\n", region)
	p := &MyProvider{}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}
	p.EC2 = ec2.New(sess)
	fmt.Println("AWS session created successfully")
	return p
}

func (p *MyProvider) CreateInstance(ami string, instanceType string) (*ec2.Reservation, error) {
	fmt.Printf("Creating instance with AMI %s and instance type %s\n", ami, instanceType)

	runResult, err := p.EC2.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(ami),
		InstanceType: aws.String(instanceType),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if err != nil {
		log.Printf("Failed to create instance: %v\n", err)
		return nil, err
	}

	fmt.Printf("Successfully created instance with reservation ID %s\n", *runResult.ReservationId)
	return runResult, nil
}

func (p *MyProvider) PlanInstance(ami string, instanceType string) {
	fmt.Printf("Plan to create instance with AMI %s and instance type %s\n", ami, instanceType)
}
