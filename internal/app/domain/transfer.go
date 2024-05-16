package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"gorm.io/datatypes"
)

type TransferStatus string

const (
	Accepted TransferStatus = "ACCEPTED"
	Rejected TransferStatus = "REJECTED"
	Deducted TransferStatus = "DEDUCTED"
	Proceed  TransferStatus = "PROCEED"
	Success  TransferStatus = "SUCCESS"
	Reversal TransferStatus = "REVERSAL"
	Failed   TransferStatus = "FAILED"
)

type Transfers struct {
	ID                 string           `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4();PRIMARY_KEY"`
	UserID             string           `json:"user_id" gorm:"column:user_id;NOT NULL"`
	DestinationAccount string           `json:"destination_account" gorm:"column:destination_account;NOT NULL"`
	BankCode           string           `json:"bank_code" gorm:"column:bank_code;NOT NULL"`
	Currency           string           `json:"currency" gorm:"column:currency;NOT NULL"`
	Amount             float64          `json:"amount" gorm:"column:amount;NOT NULL"`
	Notes              string           `json:"notes" gorm:"column:notes;NOT NULL"`
	ReferenceID        string           `json:"reference_id" gorm:"column:reference_id;NOT NULL"`
	Status             string           `json:"status" gorm:"column:status;default:ACCEPTED"`
	BankTransferID     string           `json:"-" gorm:"column:bank_transfer_id;default:null"`
	TransferMetadata   TransferMetadata `json:"metadata" gorm:"-"`
	Metadata           datatypes.JSON   `json:"-" gorm:"column:metadata;type:JSONB;NOT NULL;default:null"`
	CreatedAt          *time.Time       `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt          *time.Time       `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt          *time.Time       `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type TransferMetadata struct {
	Recipient Recipient `json:"recipient"`
}

type Recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Serialize converts struct fields into JSON raw data, so it can be saved in database
func (d *Transfers) Serialize() error {
	metadataBytes := new(bytes.Buffer)
	err := json.NewEncoder(metadataBytes).Encode(d.TransferMetadata)
	if err != nil {
		return err
	}

	d.Metadata = metadataBytes.Bytes()
	return nil
}

// Deserialize converts JSON raw data into struct fields, so it can be used programmatically
func (d *Transfers) Deserialize() error {
	err := json.Unmarshal(d.Metadata, &d.TransferMetadata)
	if err != nil {
		return err
	}

	return nil
}

type ITransferUsecase interface {
	CreateTransfers(ctx context.Context, data *Transfers) error
	ProceedTransfer(ctx context.Context, data *Transfers) error
}

type ITransferRepo interface {
	Create(ctx context.Context, data *Transfers) error
	FindByID(ctx context.Context, id string) (Transfers, error)
	FindByBankTransferID(ctx context.Context, bankTransferId string) (Transfers, error)
	UpdateByID(ctx context.Context, data *Transfers) error
}
