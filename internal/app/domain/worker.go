package domain

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type ISubscribe interface {
	Subscribe(queueUrl string, ctx context.Context, fn func(ctx context.Context, m *sqs.Message) error)
}

type IWorker interface {
	GetHandler() func(ctx context.Context, m *sqs.Message) error
}
