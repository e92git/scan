package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func SendRequest(client *http.Client, method string, url string, jsonData string) (*int, *[]byte, error) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	statusCode := response.StatusCode

	return &statusCode, &body, nil
}