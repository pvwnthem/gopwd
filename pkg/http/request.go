package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Client) Request(method, url string, body interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)

	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.ua)

	return req, nil
}
