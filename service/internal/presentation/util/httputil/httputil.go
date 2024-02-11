package httputil

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func ParseJSONRequestBody(r *http.Request, params interface{}) error {
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return err
	}

	err = json.Unmarshal(body[:length], &params)
	if err != nil {
		return err
	}

	return nil
}
