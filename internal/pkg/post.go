package pkg

import (
	"bytes"
	"net/http"
)

// Post send http post
func Post(url string, request []byte, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(request))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}
	return res, nil
}
