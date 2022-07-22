package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateService(service ServicePostStruct) (*ServicePostResponse, error) {
	requestBody, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/ResourceServer/api/v6/internal/Service", c.HostURL),
		strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(request, &c.AccessToken)
	if err != nil {
		return nil, err
	}

	newService := ServicePostResponse{}
	if err := json.Unmarshal(body, &newService); err != nil {
		return nil, err
	}

	return &newService, err
}

func (c *Client) GetService() (*ServiceGetResponse, error) {
		
}