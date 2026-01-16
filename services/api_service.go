package services

import (
	"encoding/json"
	"net/http"
)

// GetArtists effectue l'appel API pour récupérer les groupes
func GetArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var artists []Artist
	json.NewDecoder(resp.Body).Decode(&artists)
	return artists, nil
}

// GetRelation récupère les liens entre artistes, dates et lieux
func GetRelation(url string) (Relation, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Relation{}, err
	}
	defer resp.Body.Close()
	var rel Relation
	json.NewDecoder(resp.Body).Decode(&rel)
	return rel, nil
}