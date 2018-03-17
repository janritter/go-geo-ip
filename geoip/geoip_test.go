package geoip

import "testing"

func TestForIP(t *testing.T) {
	geoInfo, err := ForIP("8.8.8.8")

	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if geoInfo.CountryName != "United States" {
		t.Error("Expected Unites States, got ", geoInfo.CountryName)
	}

	if geoInfo.IP != "8.8.8.8" {
		t.Error("Expected 8.8.8.8, got ", geoInfo.IP)
	}
}

func TestForIPInvalid(t *testing.T){
	_, err := ForIP("8.8.8")

	if err == nil {
		t.Error("Expected error to contain an error description, got ", err)
	}
}

func TestForDomain(t *testing.T) {
	geoInfo, err := ForDomain("google.com")

	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if geoInfo.CountryCode != "US" {
		t.Error("expected US, got ", geoInfo.CountryCode)
	}

	if geoInfo.CountryName != "United States" {
		t.Error("Expected Unites States, got ", geoInfo.CountryName)
	}
}
