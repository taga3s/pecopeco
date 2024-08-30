package notify

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/taga3s/pecopeco-cli/config"
)

func HttpClient(method string, message string, response interface{}) error {
	uri := os.Getenv("LINE_NOTIFY_API_URL")
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("message", message)
	reqBody := bytes.NewReader([]byte(form.Encode()))

	accessToken := viper.GetString(config.LINE_NOTIFY_API_TOKEN)

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(reqBody)

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

	err = json.Unmarshal(byteArray, response)
	if err != nil {
		return err
	}

	return nil
}
