package main

import (
	"time"
)

//https://garmindev.dev.logisensebilling.com/ResourceServer/api/v6/internal/Service
type ServicePostBody struct {
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
}

type UsageBucket struct {
	Items []int `json:"items"` //CHECK WITH PHANI WHAT VALUES GO IN HERE
}
