package worker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	"github.com/go-playground/validator"
)

type ProceedTransferWorker struct {
	TransferUsecase domain.ITransferUsecase
	Validator          validator.Validate
}

func NewProceedTransferWorker(TransferUsecase domain.ITransferUsecase) ProceedTransferWorker {
	return ProceedTransferWorker{
		TransferUsecase: TransferUsecase,
		Validator:          *validator.New(),
	}
}

func (w ProceedTransferWorker) validateMessage(ctx context.Context, m *sqs.Message) (payload domain.Transfers, err error) {
	err = json.Unmarshal([]byte(*m.Body), &payload)
	if err != nil {
		log.Default().Println(err)
		return
	}

	err = w.Validator.Struct(payload)
	if err != nil {
		log.Default().Println(err)
		return
	}

	return
}

func (w ProceedTransferWorker) GetHandler() func(ctx context.Context, m *sqs.Message) error {
	return func(ctx context.Context, m *sqs.Message) (err error) {
		defer func() {
			if err := recover(); err != nil {
				log.Default().Println(map[string]interface{}{"error": err}, "Can't process this message. This message caused the worker crashed. Recovering..", err.(error))
				// move to handler for message that cant be procceed
			}
		}()

		payload, err := w.validateMessage(ctx, m)
		if err != nil {
			log.Default().Println(err)
			return
		}

		err = w.TransferUsecase.ProceedTransfer(ctx, &payload)
		if err != nil {
			// To Do: Should handle error that can be requeued
			log.Default().Println("Error when processing the message ", err, payload)
			return err
		}

		log.Default().Println(ctx, "Message processed", payload)

		return
	}
}
