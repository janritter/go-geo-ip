// Package geoip provides a wrapper for the freegeoip.net api, which returns geo information for
// a given IP address or domain.
package geoip

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

// This struct is filled with the information of the response json, received for the specified IP
// Not all values have to be set in the response struct
type IPGeoInfo struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

// ForIP takes a IP address as string and fetches the geo information for it.
// It returns a IPGeoInfo struct with the result information.
func ForIP(ip string) IPGeoInfo {
	return makeAPICall(ip)
}

// ForDomain takes a domain as string and fetches the geo information for it.
// The Domain gets first resolved to the corresponding IP.
// It returns a IPGeoInfo struct with the result information.
func ForDomain(domain string) IPGeoInfo {
	return makeAPICall(domain)
}

func makeAPICall(data string) IPGeoInfo {
	url := "https://freegeoip.net/json/"+data

	httpClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	geoInfo := IPGeoInfo{}
	jsonErr := json.Unmarshal(body, &geoInfo)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return geoInfo
}