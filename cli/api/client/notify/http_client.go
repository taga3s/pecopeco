package notify

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/Seiya-Tagami/pecopeco-cli/config"
	"github.com/spf13/viper"
)

func HttpClient(method string, message string, response interface{}) error {
	uri := os.Getenv("LINE_NOTIFY_API_URL")
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err
	}

	var reqBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&reqBody)
	err = multipartWriter.WriteField("message", message)
	if err != nil {
		return err
	}

	err = multipartWriter.Close()
	if err != nil {
		return err
	}

	accessToken := viper.GetString(config.LINE_NOTIFY_API_TOKEN)

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

	err = json.Unmarshal(byteArray, response)
	if err != nil {
		return err
	}

	return nil
}
