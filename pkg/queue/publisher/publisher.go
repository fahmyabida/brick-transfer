package publisher

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/fahmyabida/brick-transfer/cmd/config"
)

type Publisher struct {
	SNS       *sns.SNS
	configAWS config.AWS
}

func NewPublisher(sns *sns.SNS, configAWS config.AWS) Publisher {
	return Publisher{
		SNS:       sns,
		configAWS: configAWS,
	}
}

func (p Publisher) publish(actionValue, topicArn string, data interface{}) error {

	// Specify the message you want to publish
	jsonByte, _ := json.Marshal(data)
	message := string(jsonByte)

	// Specify message attributes for filtering
	attributes := map[string]*sns.MessageAttributeValue{
		"action": {
			DataType:    aws.String("String"),
			StringValue: aws.String(actionValue),
		},
	}

	// Publish a message to the specified topic with message attributes
	_, err := p.SNS.Publish(&sns.PublishInput{
		Message:           aws.String(message),
		MessageAttributes: attributes,
		TopicArn:          aws.String(topicArn),
	})

	if err != nil {
		log.Default().Println("Error publishing message:", err)
		return err
	}

	return nil
}
