package services

// FiltrerArtistes gère les filtres de plage (dates) et de sélection (membres)
func ApplyFilters(artists []Artist, minDate, maxDate int, memberCount int) []Artist {
	var filtered []Artist
	for _, a := range artists {
		// Filtre par date de création
		matchDate := a.CreationDate >= minDate && a.CreationDate <= maxDate
		// Filtre par nombre de membres (checkbox/selection)
		matchMembers := memberCount == 0 || len(a.Members) == memberCount

		if matchDate && matchMembers {
			filtered = append(filtered, a)
		}
	}
	return filtered
}