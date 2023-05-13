package pkg

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Post send http post
func Post(url string, request []byte, header map[string]string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(request))
	if err != nil {
		return nil, nil, err
	}
	var body []byte
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	// 这里的顺序很有讲究，不要改动
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return res, body, err
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return res, body, err
	}
	return res, body, nil
}
