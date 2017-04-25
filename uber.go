package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const uberEstimateURL = "https://api.uber.com/v1.2/estimates/price?start_latitude=%f&start_longitude=%f&end_latitude=%f&end_longitude=%f"

// UberEstimateResponse holds the JSON response object from Uber
type UberEstimateResponse struct {
	Prices []struct {
		LocalizedDisplayName string  `json:"localized_display_name"`
		Distance             float64 `json:"distance"`
		DisplayName          string  `json:"display_name"`
		ProductID            string  `json:"product_id"`
		HighEstimate         float64 `json:"high_estimate"`
		LowEstimate          float64 `json:"low_estimate"`
		Duration             int     `json:"duration"`
		Estimate             string  `json:"estimate"`
		CurrencyCode         string  `json:"currency_code"`
	} `json:"prices"`
}

// GetUberEstimates returns fare estimates for a given start end end address
func GetUberEstimates(startAddress *Address, endAddress *Address) *UberEstimateResponse {
	req, err := http.NewRequest("GET", fmt.Sprintf(uberEstimateURL, startAddress.Latitude, startAddress.Longitude, endAddress.Latitude, endAddress.Longitude), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", configuration.Uber.ServerToken))
	req.Header.Set("Accept-Language", "en_US")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	uberEstimate := &UberEstimateResponse{}
	err = json.Unmarshal(body, uberEstimate)
	if err != nil {
		panic(err)
	}

	return uberEstimate
}
