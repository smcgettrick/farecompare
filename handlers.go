package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func EstimateHandler(w http.ResponseWriter, r *http.Request) {
	estimateRequest := &EstimateRequest{}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, estimateRequest); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	startAddress := estimateRequest.StartAddress
	startAddress.Geocode()

	endAddress := estimateRequest.EndAddress
	endAddress.Geocode()

	estimateResponse := &EstimateResponse{}

	for _, u := range GetUberEstimates(&startAddress, &endAddress).Prices {
		e := Estimate{u.LocalizedDisplayName, u.Estimate}
		estimateResponse.AddUberEstimate(e)
	}

	for _, l := range GetLyftEstimates(&startAddress, &endAddress).CostEstimates {
		e := Estimate{l.DisplayName, formatLyftEstimate(l.EstimatedCostCentsMin, l.EstimatedCostCentsMax)}
		estimateResponse.AddLyftEstimate(e)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(estimateResponse); err != nil {
		panic(err)
	}
}
