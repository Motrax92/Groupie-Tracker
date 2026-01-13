package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"groupie-tracker/api"
	"groupie-tracker/core"
)

type AppState struct {
	App        fyne.App
	Window     fyne.Window
	AllArtists []api.ArtistFull
	Filtered   []api.ArtistFull

	Filters     core.FilterOptions
	SearchIndex []core.SearchResult

	SearchEntry *widget.Entry
	SearchList  *widget.List
	ArtistList  *widget.List
	Detail      *widget.Label
	StatusLabel *widget.Label
}

func NewApp(all []api.ArtistFull) *AppState {
	a := app.New()
	w := a.NewWindow("Groupie Tracker")

	state := &AppState{
		App:        a,
		Window:     w,
		AllArtists: all,
		Filtered:   all,
		Filters:    core.FilterOptions{},
	}

	state.SearchIndex = core.BuildSearchIndex(all)
	state.buildUI()

	return state
}

func (s *AppState) buildUI() {
	s.StatusLabel = widget.NewLabel("Ready")

	s.ArtistList = widget.NewList(
		func() int {
			return len(s.Filtered)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(s.Filtered[i].Artist.Name)
		},
	)

	s.ArtistList.OnSelected = func(id widget.ListItemID) {
		if id < 0 || id >= len(s.Filtered) {
			return
		}
		s.showArtistDetail(s.Filtered[id])
	}

	s.Detail = widget.NewLabel("Select an artist")
	s.Detail.Wrapping = fyne.TextWrapWord

	searchBar, searchList := s.buildSearchBar()
	s.SearchEntry = searchBar
	s.SearchList = searchList

	filtersBox := s.buildFilterControls()

	left := container.NewBorder(
		container.NewVBox(
			searchBar,
			searchList,
			filtersBox,
		),
		s.StatusLabel,
		nil,
		nil,
		s.ArtistList,
	)

	content := container.NewHSplit(left, s.Detail)
	content.SetOffset(0.35)

	s.Window.SetContent(content)
	s.Window.Resize(fyne.NewSize(1200, 700))
}

func (s *AppState) showArtistDetail(a api.ArtistFull) {
	text := fmt.Sprintf(
		"Name: %s\nCreation date: %d\nFirst album: %s\n\nMembers:\n",
		a.Artist.Name,
		a.Artist.CreationDate,
		a.Artist.FirstAlbum,
	)

	for _, m := range a.Artist.Members {
		text += " - " + m + "\n"
	}

	text += "\nLocations:\n"
	for _, loc := range a.Locations.Locations {
		text += " - " + loc + "\n"
	}

	text += "\nConcerts:\n"
	for loc, dates := range a.Relation.DatesLocations {
		text += " " + loc + ":\n"
		for _, d := range dates {
			text += "   * " + d + "\n"
		}
	}

	s.Detail.SetText(text)
}

func (s *AppState) applyFilters() {
	s.Filtered = core.FilterArtists(s.AllArtists, s.Filters)
	s.ArtistList.Refresh()
	s.StatusLabel.SetText(fmt.Sprintf("%d artists shown", len(s.Filtered)))
}
