package core

import (
	"strings"
	"time"

	"groupie-tracker/api"
)

type FilterOptions struct {
	CreationMin       int
	CreationMax       int
	AlbumMinYear      int
	AlbumMaxYear      int
	MembersCounts     []int
	SelectedLocations []string
}

func FilterArtists(artists []api.ArtistFull, opts FilterOptions) []api.ArtistFull {
	var res []api.ArtistFull

	for _, a := range artists {
		if !matchCreationDate(a, opts) {
			continue
		}
		if !matchAlbumDate(a, opts) {
			continue
		}
		if !matchMembers(a, opts) {
			continue
		}
		if !matchLocations(a, opts) {
			continue
		}
		res = append(res, a)
	}
	return res
}

func matchCreationDate(a api.ArtistFull, opts FilterOptions) bool {
	year := a.Artist.CreationDate

	if opts.CreationMin != 0 && year < opts.CreationMin {
		return false
	}
	if opts.CreationMax != 0 && year > opts.CreationMax {
		return false
	}
	return true
}

func matchAlbumDate(a api.ArtistFull, opts FilterOptions) bool {
	year := albumYear(a.Artist.FirstAlbum)
	if year == 0 {
		return true
	}

	if opts.AlbumMinYear != 0 && year < opts.AlbumMinYear {
		return false
	}
	if opts.AlbumMaxYear != 0 && year > opts.AlbumMaxYear {
		return false
	}
	return true
}

func albumYear(dateStr string) int {
	t, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return 0
	}
	return t.Year()
}

func matchMembers(a api.ArtistFull, opts FilterOptions) bool {
	if len(opts.MembersCounts) == 0 {
		return true
	}

	n := len(a.Artist.Members)

	for _, c := range opts.MembersCounts {
		if c == n {
			return true
		}
		if c == 99 && n >= 5 {
			return true
		}
	}
	return false
}

func matchLocations(a api.ArtistFull, opts FilterOptions) bool {
	if len(opts.SelectedLocations) == 0 {
		return true
	}

	artistLocations := make([]string, 0)
	for _, loc := range a.Locations.Locations {
		artistLocations = append(artistLocations, strings.ToLower(loc))
	}

	for _, selected := range opts.SelectedLocations {
		selected = strings.ToLower(selected)
		for _, loc := range artistLocations {
			if strings.Contains(loc, selected) {
				return true
			}
		}
	}

	return false
}