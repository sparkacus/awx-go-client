package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Inventory struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Organization int    `json:"organization"`
}

func (s *Client) GetOrganizations() (*Inventory, error) {
	//url := fmt.Sprintf(s.BaseURL+"/%s/todos/%d", s.Username, id)
	url := fmt.Sprint(s.BaseURL + "/organizations/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Inventory
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
