package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Location holds the coordinates and metadata for a zip code.
// Fields are capitalized (Exported) so they are visible to the main app.
type Location struct {
	Zip       string  `json:"post code"`
	City      string  `json:"place name"`
	State     string  `json:"state abbreviation"`
	Latitude  float64 `json:"latitude,string"` // API returns strings; ,string directive converts to float
	Longitude float64 `json:"longitude,string"`
}

// zippoResponse reflects the top-level structure of api.zippopotam.us
type zippoResponse struct {
	PostCode string     `json:"post code"`
	Places   []Location `json:"places"`
}

func init() {
	// This prints once when the package is imported
	fmt.Printf("[GEO PACKAGE] Author: Joe O'Neil | Loaded at: %s\n", time.Now().Format("15:04:05"))
}

// GetCoords fetches city, state, lat, and lon for a US zip code.
func GetCoords(zip string) (Location, error) {
	url := fmt.Sprintf("http://api.zippopotam.us/us/%s", zip)

	resp, err := http.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("zip code %s not found (status %d)", zip, resp.StatusCode)
	}

	var data zippoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return Location{}, fmt.Errorf("json decode error: %w", err)
	}

	if len(data.Places) == 0 {
		return Location{}, fmt.Errorf("no data found for zip code: %s", zip)
	}

	// NORMALIZATION:
	// The API puts the zip code at the top level ("post code").
	// We grab the first place found and inject that zip code into our struct.
	result := data.Places[0]
	result.Zip = data.PostCode

	return result, nil
}
