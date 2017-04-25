package main

import (
	"fmt"

	"github.com/jasonwinn/geocoder"
)

// Address holds search addresses
type Address struct {
	StreetAddress string  `json:"StreetAddress"`
	City          string  `json:"City"`
	State         string  `json:"State"`
	ZipCode       string  `json:"ZipCode"`
	Latitude      float64 `json:"Latitude"`
	Longitude     float64 `json:"Longitude"`
}

// Geocode converts an address into a lat/long pair
func (a *Address) Geocode() {
	lat, long, err := geocoder.Geocode(a.String())
	if err != nil {
		panic(err)
	}

	a.Latitude = lat
	a.Longitude = long
}

func (a Address) String() string {
	return fmt.Sprintf("%s %s, %s %s", a.StreetAddress, a.City, a.State, a.ZipCode)
}
