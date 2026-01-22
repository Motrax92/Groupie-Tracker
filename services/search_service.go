package services

import (
	"fmt"
	"strings"

	"groupie-tracker/models"
)

func FilterArtistsByQuery(artists []models.Artist, query string) []models.Artist {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return artists
	}

	var result []models.Artist

	for _, a := range artists {
		match := false

		if strings.Contains(strings.ToLower(a.Name), query) {
			match = true
		}

		for _, m := range a.Members {
			if strings.Contains(strings.ToLower(m), query) {
				match = true
				break
			}
		}

		if strings.Contains(strings.ToLower(a.FirstAlbum), query) {
			match = true
		}

		if match {
			result = append(result, a)
		}
	}

	return result
}

func SearchSuggestions(artists []models.Artist, query string) []string {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return []string{}
	}

	var out []string

	for _, a := range artists {
		if strings.Contains(strings.ToLower(a.Name), query) {
			out = append(out, a.Name+" — artiste")
		}

		for _, m := range a.Members {
			if strings.Contains(strings.ToLower(m), query) {
				out = append(out, m+" — membre")
			}
		}

		if strings.Contains(strings.ToLower(a.FirstAlbum), query) {
			out = append(out, a.FirstAlbum+" — premier album")
		}

		year := fmt.Sprint(a.CreationDate)
		if strings.Contains(year, query) {
			out = append(out, year+" — année de création")
		}
	}

	return out
}
