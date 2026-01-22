package services

import (
	"encoding/json"
	"net/http"
)

type LocationsResponse struct {
	Locations []string `json:"locations"`
}

func GetArtistLocations(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data LocationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Locations, nil
}
