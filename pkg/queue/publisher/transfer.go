package publisher

import (
	"github.com/fahmyabida/brick-transfer/internal/app/domain"
	"github.com/fahmyabida/brick-transfer/pkg/queue"
)

func (p Publisher) PublishAcceptedTransfer(data domain.Transfers) error {
	return p.publish(
		queue.BrickTransfer.AcceptedAction.ActionValue,
		p.configAWS.TopicArn,
		data,
	)
}

func (p Publisher) PublishProceedTransfer(data domain.Transfers) error {
	return p.publish(
		queue.BrickTransfer.ProceedAction.ActionValue,
		p.configAWS.TopicArn,
		data,
	)
}

func (p Publisher) PublishReversalTransfer(data domain.Transfers) error {
	return p.publish(
		queue.BrickTransfer.ReversalAction.ActionValue,
		p.configAWS.TopicArn,
		data,
	)
}
