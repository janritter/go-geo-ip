package geoip

import "testing"

func TestForIp(t *testing.T) {
	geoInfo := ForIp("8.8.8.8")

	if geoInfo.CountryName != "United States" {
		t.Error("Expected Unites States, got ", geoInfo.CountryName)
	}

	if geoInfo.IP != "8.8.8.8" {
		t.Error("Expected 8.8.8.8, got ", geoInfo.IP)
	}
}

func TestForDomain(t *testing.T) {
	geoInfo := ForDomain("google.com")

	if geoInfo.CountryCode != "US" {
		t.Error("expected US, got ", geoInfo.CountryCode)
	}

	if geoInfo.CountryName != "United States" {
		t.Error("Expected Unites States, got ", geoInfo.CountryName)
	}
}
