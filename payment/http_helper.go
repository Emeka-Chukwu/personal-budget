package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HttpCaller(method, linkUrl, token string, body interface{}, data interface{}) (interface{}, error) {
	var jsonData []byte
	if body != nil {
		jsonData, _ = json.Marshal(body)
	}
	request, err := http.NewRequest(strings.ToUpper(method), linkUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err, response.StatusCode)
		return nil, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if !(response.StatusCode == 200 || response.StatusCode == 201 || response.StatusCode == 202) {
		err = fmt.Errorf("Error occured: %v", data)
	}
	return data, err
}
