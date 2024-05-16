package worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Subscribe struct {
	SQS *sqs.SQS
}

func NewSubcribe(sqs *sqs.SQS) Subscribe {
	return Subscribe{SQS: sqs}
}

func (w Subscribe) Subscribe(queueURL string, ctx context.Context, fn func(ctx context.Context, m *sqs.Message) error) {

	// Create a new input for the ReceiveMessage API call
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(10), // Maximum number of messages to receive
		WaitTimeSeconds:     aws.Int64(20), // Wait time for long polling
	}

	// Create a wait group to wait for goroutines to finish
	var wg sync.WaitGroup
	// Add 1 to the wait group to indicate the main goroutine
	wg.Add(1)

	// Start a goroutine to continuously poll messages from the queue
	go func() {
		defer wg.Done() // Decrement the wait group when the goroutine exits
		for {
			// Make the API call to receive messages from the queue
			resp, err := w.SQS.ReceiveMessage(params)
			if err != nil {
				fmt.Println("Error receiving message:", err)
				continue
			}

			// Process received messages
			for _, msg := range resp.Messages {

				fmt.Println("Received message:", *msg.Body)

				fn(ctx, msg)

				// Once processed, delete the message from the queue
				_, err := w.SQS.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
					ReceiptHandle: msg.ReceiptHandle,
				})
				if err != nil {
					fmt.Println("Error deleting message:", err)
					continue
				}
				fmt.Println("Deleted message:", *msg.MessageId)
			}
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
