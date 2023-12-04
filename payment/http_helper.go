package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HttpCaller[T any](method, linkUrl, token string, body interface{}) (T, error) {
	var jsonData []byte
	var resp T
	if body != nil {
		jsonData, _ = json.Marshal(body)
	}
	request, err := http.NewRequest(strings.ToUpper(method), linkUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return resp, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err, response.StatusCode)
		return resp, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&resp)
	if !(response.StatusCode == 200 || response.StatusCode == 201 || response.StatusCode == 202) {
		err = fmt.Errorf("Error occured: %v", resp)
	}
	return resp, err
}
