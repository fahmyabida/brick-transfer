package domain

type IPublisher interface {
	PublishAcceptedTransfer(data Transfers) error
	PublishProceedTransfer(data Transfers) error
	PublishReversalTransfer(data Transfers) error
}
