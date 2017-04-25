package main

import (
	"log"
	"net/http"

	"github.com/jasonwinn/geocoder"
)

var configuration Configuration

func init() {
	configuration = LoadConfiguration()

	// MapQuest API Key
	geocoder.SetAPIKey(configuration.MapQuest.ConsumerKey)
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
