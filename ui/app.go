package ui

import (
	"groupie-tracker/models"
	"groupie-tracker/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewAppWindow(app fyne.App, artists []models.Artist) fyne.Window {
	window := app.NewWindow("Groupie Tracker")
	window.Resize(fyne.NewSize(1000, 700))

	artistList := container.NewVBox()

	updateArtists := func(list []models.Artist) {
		artistList.Objects = nil
		for _, a := range list {
			artistList.Add(NewArtistCard(a))
		}
		artistList.Refresh()
	}


	updateArtists(artists)

	
	searchBar := NewSearchBar(func(query string) {
		_ = services.SearchEngine(artists, query)
		
	})


	filterPanel := NewFilterPanel(func() {
		filtered := services.ApplyFilters(
			artists,
			1980, 
			2025, 
			0,    
		)
		updateArtists(filtered)
	})

	left := container.NewVBox(
		searchBar,
		filterPanel,
		container.NewVScroll(artistList),
	)

	content := container.NewHSplit(left, NewMapView())
	content.SetOffset(0.4)

	window.SetContent(content)
	return window
}