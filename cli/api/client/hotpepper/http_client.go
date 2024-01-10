package hotpepper

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func HttpClient(method string, queryParams string, response interface{}) error {
	uri := os.Getenv("HOTPEPPER_API_URL")
	key := os.Getenv("HOTPEPPER_API_KEY")
	req, _ := http.NewRequest(method, uri+"/?key="+key+queryParams, nil)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	byteArray, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal(byteArray, &response); err != nil {
		return err
	}

	return nil
}
