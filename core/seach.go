package core

import (
	"fmt"
	"strconv"
	"strings"

	"groupie-tracker/api"
)

type SearchResult struct {
	Label string
	Type  string
	ID    int // ID artiste
}

// Crée un index de recherche pour toutes les données
func BuildSearchIndex(artists []api.ArtistFull) []SearchResult {
	var index []SearchResult
	for _, a := range artists {
		// artiste
		index = append(index, SearchResult{
			Label: fmt.Sprintf("%s – artist/band", a.Artist.Name),
			Type:  "artist",
			ID:    a.Artist.ID,
		})

		// membres
		for _, m := range a.Artist.Members {
			index = append(index, SearchResult{
				Label: fmt.Sprintf("%s – member", m),
				Type:  "member",
				ID:    a.Artist.ID,
			})
		}

		// lieux
		for _, loc := range a.Locations.Locations {
			index = append(index, SearchResult{
				Label: fmt.Sprintf("%s – location", loc),
				Type:  "location",
				ID:    a.Artist.ID,
			})
		}

		// date de création
		index = append(index, SearchResult{
			Label: fmt.Sprintf("%d – creation date (%s)", a.Artist.CreationDate, a.Artist.Name),
			Type:  "creationDate",
			ID:    a.Artist.ID,
		})

		// date premier album
		index = append(index, SearchResult{
			Label: fmt.Sprintf("%s – first album (%s)", a.Artist.FirstAlbum, a.Artist.Name),
			Type:  "firstAlbum",
			ID:    a.Artist.ID,
		})
	}
	return index
}

func Search(index []SearchResult, query string, limit int) []SearchResult {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return nil
	}
	var res []SearchResult
	for _, r := range index {
		if strings.Contains(strings.ToLower(r.Label), query) {
			res = append(res, r)
			if limit > 0 && len(res) >= limit {
				break
			}
		}
	}
	return res
}

// utilitaire simple pour transformer une string année
func ParseYear(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 4 {
		n, _ := strconv.Atoi(s)
		return n
	}
	return 0
}