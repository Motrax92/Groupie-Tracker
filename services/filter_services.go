package services

import "groupie-tracker/models"

func ApplyFilters(artists []models.Artist, minYear, maxYear, memberCount int) []models.Artist {
	var filtered []models.Artist

	for _, a := range artists {
		if minYear != 0 && a.CreationDate < minYear {
			continue
		}
		if maxYear != 0 && a.CreationDate > maxYear {
			continue
		}
		if memberCount != 0 && len(a.Members) != memberCount {
			continue
		}
		filtered = append(filtered, a)
	}

	return filtered
}
