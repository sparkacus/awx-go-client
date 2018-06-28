package awx

import (
	"net/http"
)

type Client struct {
	Username string
	Password string
	BaseURL  string
}

func (s *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(s.Username, s.Password)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
