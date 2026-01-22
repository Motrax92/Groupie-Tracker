package services

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/models"
)

const artistsURL = "https://groupietrackers.herokuapp.com/api/artists"

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get(artistsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}
