package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const lyftAuthenticationURL = "https://api.lyft.com/oauth/token"
const lyftEstimateURL = "https://api.lyft.com/v1/cost?start_lat=%f&start_lng=%f&end_lat=%f&end_lng=%f"

type lyftAuthenticationRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
}

type lyftAuthenticationResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// LyftEstimateResponse holds the JSON response object from Lyft
type LyftEstimateResponse struct {
	CostEstimates []struct {
		RideType                   string      `json:"ride_type"`
		EstimatedDurationSeconds   int         `json:"estimated_duration_seconds"`
		EstimatedDistanceMiles     float64     `json:"estimated_distance_miles"`
		EstimatedCostCentsMax      int         `json:"estimated_cost_cents_max"`
		PrimetimePercentage        string      `json:"primetime_percentage"`
		IsValidEstimate            bool        `json:"is_valid_estimate"`
		Currency                   string      `json:"currency"`
		CostToken                  interface{} `json:"cost_token"`
		EstimatedCostCentsMin      int         `json:"estimated_cost_cents_min"`
		DisplayName                string      `json:"display_name"`
		PrimetimeConfirmationToken interface{} `json:"primetime_confirmation_token"`
		CanRequestRide             bool        `json:"can_request_ride"`
	} `json:"cost_estimates"`
}

// GetLyftEstimates returns fare estimates for a given start end end address
func GetLyftEstimates(startAddress *Address, endAddress *Address) *LyftEstimateResponse {
	lyftAuthenticationRequest := &lyftAuthenticationRequest{}
	lyftAuthenticationRequest.GrantType = "client_credentials"
	lyftAuthenticationRequest.Scope = "public"
	requestBody, err := json.Marshal(lyftAuthenticationRequest)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", lyftAuthenticationURL, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(configuration.Lyft.ClientID, configuration.Lyft.ClientSecret)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	lyftAuthenticationResponse := &lyftAuthenticationResponse{}
	err = json.Unmarshal(body, lyftAuthenticationResponse)
	if err != nil {
		panic(err)
	}

	req, err = http.NewRequest("GET", fmt.Sprintf(lyftEstimateURL, startAddress.Latitude, startAddress.Longitude, endAddress.Latitude, endAddress.Longitude), nil)
	if err != nil {
		panic(err)
	}

	lyftAccessToken := lyftAuthenticationResponse.AccessToken
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", lyftAccessToken))

	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	lyftEstimateResponse := &LyftEstimateResponse{}
	err = json.Unmarshal(body, lyftEstimateResponse)
	if err != nil {
		panic(err)
	}

	return lyftEstimateResponse
}

func formatLyftEstimate(minCents int, maxCents int) string {
	minDollars := minCents / 60
	maxDollars := maxCents / 60

	if minDollars == maxDollars {
		return fmt.Sprintf("$%d", minDollars)
	}

	return fmt.Sprintf("$%d-%d", minDollars, maxDollars)

}
