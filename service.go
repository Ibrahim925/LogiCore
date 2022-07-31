package logicore

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

	return &newService, nil
}

func (c *Client) GetServices() (*ServicesGetResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/ResourceServer/api/v6/internal/Service", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(request, &c.AccessToken)
	if err != nil {
		return nil, err
	}

	services := ServicesGetResponse{}
	if err := json.Unmarshal(body, &services); err != nil {
		return nil, err
	}

	return &services, nil
}

func (c *Client) GetService(id int) (*ServiceGetResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/ResourceServer/api/v6/internal/Service/%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(request, &c.AccessToken)
	if err != nil {
		return nil, err
	}

	service := ServiceGetResponse{}
	if err := json.Unmarshal(body, &service); err != nil {
		return nil, err
	}

	return &service, nil
}
