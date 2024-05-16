package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type AWS struct {
	Endpoint             string `envconfig:"ENDPOINT" required:"true"`
	Region               string `envconfig:"REGION" required:"true"`
	AccessKeyId          string `envconfig:"ACCESS_KEY_ID" required:"true"`
	SecretAccessKey      string `envconfig:"SECRET_ACCESS_KEY" required:"true"`
	DeductBalanceQueue   string `envconfig:"SQS_DEDUCT_BALANCE"`
	ProceedTransferQueue string `envconfig:"SQS_PROCEED_TRANSFER"`
	ReversalBalanceQueue string `envconfig:"SQS_REVERSAL_BALANCE"`
	TopicArn             string `envconfig:"SNS_TOPIC_ARN"`
}

func InitAWS_SNS_SQS(config *AWS) (SNS *sns.SNS, SQS *sqs.SQS) {

	// Create a new AWS session
	sessSNS := session.Must(session.NewSession(&aws.Config{
		Endpoint: &config.Endpoint,
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyId,
			config.SecretAccessKey,
			""),
		Region: aws.String(config.Region),
	}))

	sessSQS := session.Must(session.NewSession(&aws.Config{
		Endpoint: &config.Endpoint,
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyId,
			config.SecretAccessKey,
			""),
		Region: aws.String(config.Region),
	}))

	SNS = sns.New(sessSNS)
	SQS = sqs.New(sessSQS)

	return SNS, SQS

}

// LoadForAWS loads AWS configuration and returns it
func LoadForAWS() (awsConfig *AWS) {
	awsConfig = &AWS{}

	mustLoad("AWS", awsConfig)

	return
}
