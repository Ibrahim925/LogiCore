package main

import (
	"time"
)

//https://base_url.com/ResourceServer/api/v6/internal/Service
type ServicePostStruct struct {
	Name              string    `json:"name"`
	IsActive          bool      `json:"isActive"`
	ServiceTypeID     int       `json:"serviceTypeId"`
	ServiceCategoryID int       `json:"serviceCategoryId"`
	IsInclusiveTaxes  bool      `json:"isInclusiveTaxes"`
	IsTaxExempt       bool      `json:"isTaxExempt"`
	Created           time.Time `json:"created"`
	Details           struct {
		UsageBucket UsageBucket `json:"usageBuckets"`
	} `json:"details"`
	Description                string `json:"description"`
	UsageFrequency             int    `json:"useageFrequency"`
	UsageFrequencyTypeID       int    `json:"usageFrequencyTypeId"`
	DefaultServiceStatusTypeID int    `json:"defaultServiceStatusTypeId"`
}

type ServicePostResponse struct {
	TrackingID string `json:"trackingId"`
	Type       string `json:"type"`
	Results    struct {
		TotalCount int `json:"totalCount"`
		Items      struct {
			Identity   int    `json:"identity"`
			Action     string `json:"action"`
			DtoTypeKey string `json:"dtoTypeKey"`
			Instance   Item   `json:"instance"`
		} `json:"items"`
	} `json:"results"`
}

//GET https://base_url.com/ResourceServer/api/v6/internal/Service
type ServiceGetResponse struct {
	TrackingID string `json:"trackingId"`
	TotalCount int    `json:"totalCount"`
	Items      []Item
}

type UsageBucket struct {
	Items []Item `json:"items"` //CHECK WITH PHANI WHAT VALUES GO IN HERE
}

type Item struct {
	Identity                     string    `json:"identity"`
	Name                         string    `json:"name"`
	Description                  string    `json:"description"`
	ServiceTypeID                int       `"serviceTypeId"`
	ServiceTypeName              string    `"serviceTypeName"`
	Created                      time.Time `json:"created"`
	IsActive                     bool      `json:"isActive"`
	IsTaxExempt                  bool      `json:"isTaxExempt"`
	IsInclusiveTaxes             bool      `json:"isInclusiveTaxes"`
	DefaultServiceStatusTypeID   int       `json:"defaultServiceStatusTypeId"`
	DefaultServiceStatusTypeName string    `json:"defaultServiceStatusTypeName"`
	ServiceCategoryID            int       `json:"serviceCategoryId"`
	ServiceCategoryName          string    `json:"serviceCategoryName"`
	ServiceBaseID                int       `json:"serviceBaseTypeId"`
	ServiceBaseName              string    `json:"serviceBaseTypeName"`
	UsageFrequency               int       `json:"usageFrequency"`
	UseageFrequencyTypeID        int       `json:"usageFrequencyTypeId"`
	UsageFrequencyTypeName       string    `json:"usageFrequencyTypeName"`
}
