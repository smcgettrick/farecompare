FareCompare
===========

Description
-----------
A Go WebAPI that takes in a start and finish address and returns all the Uber and Lyft fare estimates for available ride types.

Installation
------------
Install the package:

	go get github.com/smcgettrick/farecompare

Install dependencies:

	godep restore

If you do not have godep then you will need to install the dependencies manually:

	go get github.com/gorilla/mux
	go get github.com/jasonwinn/geocoder

Build the application

	go build
	
Configuration
-------------
This application requires you to provide API keys for Lyft, MapQuest (for geocoding), and Uber.  Create a <b>conf.json</b> file as follows:

    {
	    "Lyft" : {
	        "ClientId" : "",
	        "ClientSecret" : ""
	    },
	    "MapQuest" : {
	        "ConsumerKey" : ""
	    },
	    "Uber" : {
	        "ServerToken" : ""
	    }
	}

Usage
-----

The default endpoint is <b>http://[server]:8080/estimate</b>.  This endpoint takes a POST request with the following JSON body:

	{
		"StartAddress" : {
			"StreetAddress": "1401 John F Kennedy Blvd",
			"City": "Philadelphia",
			"State": "PA",
			"ZipCode": "19102",
			"Latitude": 0.0,
			"Longitude": 0.0
		},
		"EndAddress" : {
			"StreetAddress": "2955 Market St",
			"City": "Philadelphia",
			"State": "PA",
			"ZipCode": "19104",
			"Latitude": 0.0,
			"Longitude": 0.0
		}
	}

Sample response:

	{
		"Uber": {
			"Estimates": [
				{
					"Type": "uberPOOL",
					"Estimate": "$3-5"
				},
				{
					"Type": "uberX",
					"Estimate": "$7-9"
				},
				{
					"Type": "uberXL",
					"Estimate": "$12-16"
				},
				{
					"Type": "UberBLACK",
					"Estimate": "$14-19"
				},
				{
					"Type": "UberBLACK with Car Seat",
					"Estimate": "$23-30"
				},
				{
					"Type": "UberSUV",
					"Estimate": "$22-28"
				},
				{
					"Type": "Family 6",
					"Estimate": "$31-39"
				},
				{
					"Type": "uberWAV",
					"Estimate": "$5-8"
				}
			]
		},
		"Lyft": {
			"Estimates": [
				{
					"Type": "Lyft Plus",
					"Estimate": "$14-26"
				},
				{
					"Type": "Lyft Line",
					"Estimate": "$6"
				},
				{
					"Type": "Lyft",
					"Estimate": "$10-17"
				}
			]
		}
	}

Credits
-------

Thanks to <a href="https://thenewstack.io/make-a-restful-json-api-go/">https://thenewstack.io/make-a-restful-json-api-go/</a> for assitance in API organization and adding logging decorator to HTTP handler functions.



























































































































































































































































































































































































































































































































































































































































































