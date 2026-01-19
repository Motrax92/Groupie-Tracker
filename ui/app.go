package ui

import (
	"groupie-tracker/api"
	"groupie-tracker/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewAppWindow(app fyne.App, artists []api.Artist) fyne.Window {
	window := app.NewWindow("Groupie Tracker")
	window.Resize(fyne.NewSize(1000, 700))

	artistList := container.NewVBox()

	updateArtists := func(list []api.Artist) {
		artistList.Objects = nil
		for _, a := range list {
			artistList.Add(NewArtistCard(a))
		}
		artistList.Refresh()
	}

	updateArtists(artists)

	searchBar := NewSearchBar(func(query string) {
		filtered := core.SearchArtists(artists, query)
		updateArtists(filtered)
	})

	filterPanel := NewFilterPanel(func() {
		filtered := core.FilterByCreationDate(artists, 1950, 2025)
		updateArtists(filtered)
	})

	mapView := NewMapView()

	left := container.NewVBox(
		searchBar,
		filterPanel,
		container.NewVScroll(artistList),
	)

	content := container.NewHSplit(left, mapView)
	content.SetOffset(0.4)

	window.SetContent(content)
	return window
}
