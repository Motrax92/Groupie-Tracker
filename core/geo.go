package core

import (
	"errors"
	"strings"
)

type Coordinates struct {
	Lat float64
	Lon float64
}

// Stub de géocodage : à remplacer par un appel HTTP vers un vrai service
func Geocode(location string) (Coordinates, error) {
	location = strings.TrimSpace(location)
	if location == "" {
		return Coordinates{}, errors.New("empty location")
	}

	// TODO: remplacer par une vraie implémentation :
	// - construire une requête HTTP vers une API de géocodage
	// - parser le JSON, etc.
	// Ici on renvoie tout au même endroit pour ne pas casser le code.
	return Coordinates{Lat: 0, Lon: 0}, nil
}