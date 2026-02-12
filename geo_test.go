package geo

import (
	"testing"
)

func TestGetCoords(t *testing.T) {
	// A known valid US zip code
	zip := "14623"
	expectedCity := "Rochester"

	loc, err := GetCoords(zip)

	// Test 1: Did we get an error?
	if err != nil {
		t.Fatalf("GetCoords failed: %v", err)
	}

	// Test 2: Is the city correct?
	if loc.City != expectedCity {
		t.Errorf("Expected city %s, but got %s", expectedCity, loc.City)
	}

	// NEW Test 3: Is the Zip field actually populated now?
	// This confirms your normalization logic (result.Zip = data.PostCode) is working.
	if loc.Zip != zip {
		t.Errorf("Expected zip %s in struct, but got %s", zip, loc.Zip)
	}

	// Test 4: Are coordinates populated?
	if loc.Latitude == 0 || loc.Longitude == 0 {
		t.Error("Latitude or Longitude returned as 0")
	}
}

func TestZipCityMismatch(t *testing.T) {
	zip := "90210"
	wrongCity := "Rochester"

	loc, err := GetCoords(zip)
	if err != nil {
		t.Fatalf("Failed to get coordinates: %v", err)
	}

	// Negative Test
	if loc.City == wrongCity {
		t.Errorf("Data mismatch: Zip %s returned %s, but that zip belongs to Beverly Hills!", zip, wrongCity)
	}
}

func TestInvalidZip(t *testing.T) {
	// An intentionally invalid zip code
	_, err := GetCoords("00000")

	if err == nil {
		t.Error("Expected an error for invalid zip code 00000, but got none")
	}
}
