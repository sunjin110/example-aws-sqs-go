package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	endpoint := "http://localhost:9324"

	region := "elasticmq"

	cfg, err := config.LoadDefaultConfig(context.Background())
	checkErr(err)
	cfg.Region = region

	sqsClient := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = &endpoint
	})

	pollingForReceivingMessages(sqsClient, "http://localhost:9324/query/test")

}

func pollingForReceivingMessages(sqsClient *sqs.Client, queueURL string) {

	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     3,
	}

	for {
		output, err := sqsClient.ReceiveMessage(context.Background(), input)
		if err != nil {
			fmt.Printf("failed receive message. err: %+v", err)
			continue
		}
		for _, message := range output.Messages {
			if message.Body != nil {
				fmt.Println("message is ", *message.Body)
				deleteMessageQueue(sqsClient, queueURL, *message.ReceiptHandle)
			}
		}
	}
}

func deleteMessageQueue(sqsClient *sqs.Client, queueURL string, receiptHandle string) {
	input := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: &receiptHandle,
	}

	_, err := sqsClient.DeleteMessage(context.Background(), input)
	if err != nil {
		fmt.Printf("failed delete message. err: %+v", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
