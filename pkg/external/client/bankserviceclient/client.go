package bankserviceclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/fahmyabida/brick-transfer/cmd/config"
)

// IBankServiceClient is interface that defines the methods that the client must implement.
type IBankServiceClient interface {
	ValidateBankAccount(ctx context.Context, payload ValidateBankAccountRequest) (ValidateBankAccountResponse, error)
	TransferMoney(ctx context.Context, payload TransferMoneyRequest) (TransferMoneyResponse, error)
}

type BankServiceClient struct {
	hostUrl    string
	httpClient *http.Client
}

func NewBankServiceClient(httpClient *http.Client, bankClientConfig config.BankClient) IBankServiceClient {
	return &BankServiceClient{
		hostUrl:    bankClientConfig.Host,
		httpClient: httpClient,
	}
}

func (c *BankServiceClient) ValidateBankAccount(ctx context.Context, payload ValidateBankAccountRequest) (
	response ValidateBankAccountResponse, err error) {

	reqPayload, err := json.Marshal(payload)
	if err != nil {
		log.Default().Println("unable to marshal request payload", err)
		return
	}

	url := c.hostUrl + "/bank-account-validate"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqPayload))

	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Default().Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Default().Println(err)
		return
	}

	json.Unmarshal(body, &response)
	if err != nil {
		log.Default().Println(err)
		return
	}

	return
}

func (c *BankServiceClient) TransferMoney(ctx context.Context, payload TransferMoneyRequest) (response TransferMoneyResponse, err error) {
	reqPayload, err := json.Marshal(payload)
	if err != nil {
		log.Default().Println("unable to marshal request payload", err)
		return
	}

	url := c.hostUrl + "/transfer"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqPayload))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Default().Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Default().Println(err)
		return
	}

	json.Unmarshal(body, &response)
	if err != nil {
		log.Default().Println(err)
		return
	}

	return

}
