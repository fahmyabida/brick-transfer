package queue

type SNSTopicConfig struct {
	Name           string
	AcceptedAction SQSQueue
	ProceedAction  SQSQueue
	ReversalAction SQSQueue
}

type SQSQueue struct {
	ActionValue string
	QueueNames  []string
}

var BrickTransfer = SNSTopicConfig{
	Name: "brick-transfer",
	AcceptedAction: SQSQueue{
		ActionValue: "accepted",
		QueueNames:  []string{"transfer-deduct-balance-queue"},
	},
	ProceedAction: SQSQueue{
		ActionValue: "proceed",
		QueueNames:  []string{"proceed-transfer-queue"},
	},
	ReversalAction: SQSQueue{
		ActionValue: "reversal",
		QueueNames:  []string{"transfer-reversal-balance-queue"},
	},
}
