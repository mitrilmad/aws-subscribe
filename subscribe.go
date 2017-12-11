package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"fmt"
	"os"
)

const endPoint = "your endpoint"
const protocol = "http"
const topicARN = "your arn"

func main(){
	httpEndpoint := getEnv("HTTPENDPOINT", endPoint)
	httpProtocol := getEnv("HTTPPROTOCOL", protocol)
	topicARN := getEnv("SNS_TOPIC_ARN", topicARN)

	Subscribe(httpEndpoint, httpProtocol, topicARN)
}

func Subscribe(endPoint string, protocol string, topicARN string){

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(),
		},
	)

	if err != nil {
		fmt.Printf("unable to initialize a session")
	}


	svc := sns.New(sess)

	input := &sns.SubscribeInput{
		Endpoint: &endPoint,
		Protocol: &protocol,
		TopicArn: &topicARN,

	}

	out, err := svc.Subscribe(input)
	if err != nil {
		fmt.Printf("unable to subscribe")
	}

	fmt.Printf(*out.SubscriptionArn)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

