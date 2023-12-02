package util

type ApiCallInterface interface {
	PaystackApiCall(bankntAccount, bankCode string) error
}

type apiCallInterface struct {
	config Config
}

func NewApiCallInterface(config Config) ApiCallInterface {
	return &apiCallInterface{config: config}
}
