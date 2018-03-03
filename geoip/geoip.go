package geoip

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type IpGeoInfo struct {
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

func ForIp(ip string) IpGeoInfo {
	return makeApiCall(ip)
}

func ForDomain(domain string) IpGeoInfo {
	return makeApiCall(domain)
}

func makeApiCall(data string) IpGeoInfo {
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

	geoInfo := IpGeoInfo{}
	jsonErr := json.Unmarshal(body, &geoInfo)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return geoInfo
}