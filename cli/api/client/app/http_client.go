package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

func HttpClient(method string, endpoint string, request interface{}, response interface{}) error {
	uri := os.Getenv("API_URI")
	req, err := http.NewRequest(method, uri+endpoint, nil)
	if err != nil {
		return err
	}

	if request != nil {
		req.Header.Set("Content-Type", "application/json")
		reqBody, err := json.Marshal(request)
		if err != nil {
			return err
		}
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}

	client := http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	byteArray, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteArray, response); err != nil {
		return err
	}

	return nil
}
