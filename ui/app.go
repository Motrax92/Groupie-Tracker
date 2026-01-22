package ui

import (
	"groupie-tracker/models"
	"groupie-tracker/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewAppWindow(app fyne.App, artists []models.Artist) fyne.Window {
	window := app.NewWindow("Groupie Tracker")
	window.Resize(fyne.NewSize(1100, 700))

	artistList := container.NewVBox()

	updateArtists := func(list []models.Artist) {
		artistList.Objects = nil
		for _, a := range list {
			aa := a
			artistList.Add(NewArtistCard(aa, func() {
				ShowArtistDetails(app, aa)
			}))
		}
		artistList.Refresh()
	}

	currentQuery := ""
	currentMin := 0
	currentMax := 0
	currentMembers := 0

	applyAll := func() {
		list := services.FilterArtistsByQuery(artists, currentQuery)
		list = services.ApplyFilters(list, currentMin, currentMax, currentMembers)
		updateArtists(list)
	}

	updateArtists(artists)

	search := NewSearchBar(func(q string) {
		currentQuery = q
		applyAll()
	})

	filters := NewFilterPanel(func(min, max, mem int) {
		currentMin = min
		currentMax = max
		currentMembers = mem
		applyAll()
	})

	scroll := container.NewVScroll(artistList)
	scroll.SetMinSize(fyne.NewSize(0, 500))

	left := container.NewVBox(
		search,
		filters,
		scroll,
	)

	content := container.NewHSplit(left, NewMapView())
	content.SetOffset(0.6)

	window.SetContent(content)
	return window
}
