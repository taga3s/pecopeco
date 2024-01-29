package hotpepper

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

func HttpClient(method string, endpoint string, queryParams string, response interface{}) error {
	uri := os.Getenv("HOTPEPPER_API_URL")
	key := os.Getenv("HOTPEPPER_API_KEY")
	req, err := http.NewRequest(method, uri+endpoint+"/?key="+key+queryParams, nil)
	if err != nil {
		return err
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
