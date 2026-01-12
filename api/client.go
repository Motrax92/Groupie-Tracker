package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://groupietrackers.herokuapp.com" // exemple courant

func fetchJSON[T any](endpoint string, target *T) error {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return fmt.Errorf("http get %s: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status %d for %s", resp.StatusCode, endpoint)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decode json %s: %w", endpoint, err)
	}
	return nil
}

func GetArtists() ([]Artist, error) {
	var artists []Artist
	err := fetchJSON("/artists", &artists)
	return artists, err
}

func GetLocations() ([]Locations, error) {
	var locations []Locations
	err := fetchJSON("/locations", &locations)
	return locations, err
}

func GetDates() ([]Dates, error) {
	var dates []Dates
	err := fetchJSON("/dates", &dates)
	return dates, err
}

func GetRelations() ([]Relation, error) {
	var relations []Relation
	err := fetchJSON("/relation", &relations)
	return relations, err
}

// Construit la liste fusionnée ArtistFull
func GetAllArtistFull() ([]ArtistFull, error) {
	artists, err := GetArtists()
	if err != nil {
		return nil, err
	}
	locations, err := GetLocations()
	if err != nil {
		return nil, err
	}
	dates, err := GetDates()
	if err != nil {
		return nil, err
	}
	relations, err := GetRelations()
	if err != nil {
		return nil, err
	}

	locMap := make(map[int]Locations)
	for _, l := range locations {
		locMap[l.ID] = l
	}
	dateMap := make(map[int]Dates)
	for _, d := range dates {
		dateMap[d.ID] = d
	}
	relMap := make(map[int]Relation)
	for _, r := range relations {
		relMap[r.ID] = r
	}

	var full []ArtistFull
	for _, a := range artists {
		full = append(full, ArtistFull{
			Artist:    a,
			Locations: locMap[a.ID],
			Dates:     dateMap[a.ID],
			Relation:  relMap[a.ID],
		})
	}
	return full, nil
}
