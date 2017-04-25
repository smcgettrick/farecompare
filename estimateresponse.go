package main

// EstimateResponse is the return JSON object to a front-end client
type EstimateResponse struct {
	Uber struct {
		Estimates []Estimate `json:"Estimates"`
	} `json:"Uber"`
	Lyft struct {
		Estimates []Estimate `json:"Estimates"`
	} `json:"Lyft"`
}

type Estimate struct {
	Type     string `json:"Type"`
	Estimate string `json:"Estimate"`
}

func (er *EstimateResponse) AddUberEstimate(e Estimate) {
	er.Uber.Estimates = append(er.Uber.Estimates, e)
}

func (er *EstimateResponse) AddLyftEstimate(e Estimate) {
	er.Lyft.Estimates = append(er.Lyft.Estimates, e)
}
