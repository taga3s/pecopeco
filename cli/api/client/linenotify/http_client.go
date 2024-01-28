package linenotify

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/viper"
)

func HttpClient(method string, message string, response interface{}) error {
	uri := os.Getenv("LINE_NOTIFY_API_URL")
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err
	}

	accessToken := viper.GetString("line_notify_api_token")

	formData := url.Values{}
	formData.Set("message", message)

	var reqBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&reqBody)

	for key, values := range formData {
		for _, value := range values {
			multipartWriter.WriteField(key, value)
		}
	}

	multipartWriter.Close()

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Body = io.NopCloser(&reqBody)

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
