package config

type BankClient struct {
	Host string `envconfig:"HOST_URL" required:"true"`
}

// LoadForAWS loads AWS configuration and returns it
func LoadForBankClient() (config *BankClient) {
	config = &BankClient{}

	mustLoad("BANK", config)

	return
}
