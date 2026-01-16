package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"groupie-tracker/models"
	"groupie-tracker/services"
	"groupie-tracker/ui"
)

func main() {

	// --- Création de l'application Fyne ---
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Trackers")
	myWindow.Resize(fyne.NewSize(1200, 800))

	// --- Services ---
	apiService := services.NewAPIService()
	searchService := services.NewSearchService()
	filterService := services.NewFilterService()
	geoService := services.NewGeoService()

	// --- Récupération des artistes depuis l'API ---
	artists, err := apiService.GetArtists()
	if err != nil {
		log.Fatal("Erreur lors de la récupération des artistes :", err)
	}

	// --- UI Components ---
	searchBar := ui.NewSearchBar(searchService, artists)
	filterPanel := ui.NewFilterPanel(filterService, artists)

	artistList := ui.NewArtistList(artists, geoService)

	// --- Layout principal ---
	topBar := container.NewVBox(
		searchBar.Render(),
		filterPanel.Render(),
	)

	content := container.NewBorder(
		topBar,  // top
		nil,     // bottom
		nil,     // left
		nil,     // right
		artistList.Render(), // center
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
