package services

import "strings"

// SearchEngine identifie et affiche les suggestions par type (membre, artiste, etc.)
func SearchEngine(artists []Artist, query string) []string {
	var suggestions []string
	query = strings.ToLower(query)

	for _, a := range artists {
		// Recherche par Nom d'artiste
		if strings.Contains(strings.ToLower(a.Name), query) {
			suggestions = append(suggestions, a.Name + " - artist/band")
		}
		// Recherche par Membres
		for _, m := range a.Members {
			if strings.Contains(strings.ToLower(m), query) {
				suggestions = append(suggestions, m + " - member")
			}
		}
		// Recherche par Lieu ou Date d'album
		if strings.Contains(strings.ToLower(a.FirstAlbum), query) {
			suggestions = append(suggestions, a.FirstAlbum + " - first album date")
		}
	}
	return suggestions
}