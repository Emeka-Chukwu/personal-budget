package util

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (api apiCallInterface) PaystackApiCall(bankntAccount, bankCode string) error {
	client := resty.New()
	resp, err := client.R().SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+api.config.PaystackKey).
		Get(api.config.PaystackBaseURL + fmt.Sprintf("/bank/resolve?account_number=%s&bank_code=%s", bankntAccount, bankCode))
	if err != nil {
		return err
	}
	var res map[string]interface{}
	if err := json.Unmarshal([]byte(resp.String()), &res); err != nil {
		return err
	}
	if !res["status"].(bool) {
		// message := res["message"]
		return errors.New(fmt.Sprint("error: %v", "not valid date"))
	}
	return nil
}
