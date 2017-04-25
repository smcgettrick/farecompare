package main

type EstimateRequest struct {
	StartAddress Address `json:"StartAddress"`
	EndAddress   Address `json:"EndAddress"`
}
