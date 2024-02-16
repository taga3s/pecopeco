package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ayanami77/pecopeco-cli/api/client/util"
	"github.com/ayanami77/pecopeco-cli/config"
)

func HttpClient(method string, endpoint string, request interface{}, response interface{}) error {
	uri := os.Getenv("API_URI")
	req, err := http.NewRequest(method, uri+endpoint, nil)
	if err != nil {
		return err
	}

	// jwtトークンをセットする
	accessToken := config.Get(config.PECOPECO_API_TOKEN)
	req.Header.Set("Authorization", accessToken)

	if request != nil {
		req.Header.Set("Content-Type", "application/json")
		reqBody, err := json.Marshal(request)
		if err != nil {
			return err
		}
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		req.ContentLength = int64(len(reqBody))
	}

	client := http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = util.CheckStatus(res.StatusCode); err != nil {
		return err
	}

	// No Contentの場合は早期リターン
	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	byteArray, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteArray, response); err != nil {
		return err
	}

	// jwtトークンを保存する
	accessToken = res.Header.Get("Authorization")
	config.Save(config.PECOPECO_API_TOKEN, accessToken)

	return nil
}
