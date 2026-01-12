package ui

import (
	"sort"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"groupie-tracker/api"
	"groupie-tracker/core"
)

func (s *AppState) buildSearchBar() (*widget.Entry, *widget.List) {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Search artist, member, location, date...")

	results := []core.SearchResult{}
	list := widget.NewList(
		func() int { return len(results) },
		func() fyne.CanvasObject {
			return widget.NewLabel("result")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(results[i].Label)
		},
	)

	entry.OnChanged = func(sq string) {
		results = core.Search(s.SearchIndex, sq, 10)
		list.Refresh()
	}

	list.OnSelected = func(id widget.ListItemID) {
		if id < 0 || id >= len(results) {
			return
		}
		r := results[id]
		s.selectArtistByID(r.ID)
	}

	return entry, list
}

func (s *AppState) selectArtistByID(id int) {
	for i, a := range s.Filtered {
		if a.Artist.ID == id {
			s.ArtistList.Select(i)
			return
		}
	}

	for _, a := range s.AllArtists {
		if a.Artist.ID == id {
			s.Filtered = []api.ArtistFull{a}
			s.ArtistList.Refresh()
			s.ArtistList.Select(0)
			return
		}
	}
}

func (s *AppState) GetAllUniqueLocations() []string {
	set := make(map[string]bool)

	for _, a := range s.AllArtists {
		for _, loc := range a.Locations.Locations {
			clean := strings.TrimSpace(loc)
			if clean != "" {
				set[clean] = true
			}
		}
	}

	var list []string
	for k := range set {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

func (s *AppState) updateSelectedLocations(checks []*widget.Check) {
	var selected []string
	for _, c := range checks {
		if c.Checked {
			selected = append(selected, c.Text)
		}
	}
	s.Filters.SelectedLocations = selected
}

func (s *AppState) buildFilterControls() fyne.CanvasObject {
	creationMinEntry := widget.NewEntry()
	creationMinEntry.SetPlaceHolder("Creation min (year)")

	creationMaxEntry := widget.NewEntry()
	creationMaxEntry.SetPlaceHolder("Creation max (year)")

	albumMinEntry := widget.NewEntry()
	albumMinEntry.SetPlaceHolder("First album min (year)")

	albumMaxEntry := widget.NewEntry()
	albumMaxEntry.SetPlaceHolder("First album max (year)")

	membersEntry := widget.NewEntry()
	membersEntry.SetPlaceHolder("Members (ex: 3,4,5+)")

	locations := s.GetAllUniqueLocations()
	locationChecks := make([]*widget.Check, len(locations))

	for i, loc := range locations {
		l := loc
		locationChecks[i] = widget.NewCheck(l, func(b bool) {
			s.updateSelectedLocations(locationChecks)
		})
	}

	applyBtn := widget.NewButton("Apply filters", func() {
		s.Filters.CreationMin = parseInt(creationMinEntry.Text)
		s.Filters.CreationMax = parseInt(creationMaxEntry.Text)

		s.Filters.AlbumMinYear = parseInt(albumMinEntry.Text)
		s.Filters.AlbumMaxYear = parseInt(albumMaxEntry.Text)

		s.Filters.MembersCounts = parseMembers(membersEntry.Text)

		s.applyFilters()
	})

	resetBtn := widget.NewButton("Reset filters", func() {
		creationMinEntry.SetText("")
		creationMaxEntry.SetText("")
		albumMinEntry.SetText("")
		albumMaxEntry.SetText("")
		membersEntry.SetText("")

		for _, c := range locationChecks {
			c.SetChecked(false)
		}

		s.Filters = core.FilterOptions{}
		s.Filtered = s.AllArtists
		s.ArtistList.Refresh()
		s.StatusLabel.SetText("Filters reset")
	})

	return container.NewVBox(
		widget.NewLabel("Filters"),
		widget.NewSeparator(),
		widget.NewLabel("Creation date range"),
		creationMinEntry,
		creationMaxEntry,
		widget.NewLabel("First album date range"),
		albumMinEntry,
		albumMaxEntry,
		widget.NewLabel("Members"),
		membersEntry,
		widget.NewLabel("Locations"),
		container.NewVScroll(container.NewVBox(locationChecks...)),
		applyBtn,
		resetBtn,
	)
}

func parseInt(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func parseMembers(s string) []int {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	var res []int
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "5+" {
			res = append(res, 99)
			continue
		}
		n, err := strconv.Atoi(p)
		if err == nil {
			res = append(res, n)
		}
	}
	return res
}
