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

// SendRequest - отправить запрос на удаленный сервер
// jsonData -  struct{..} OR []byte(`{..}`) (string с json не работает)
func SendRequest(client *http.Client, method string, url string, jsonData any, authHeader string) (*int, *[]byte, error) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

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