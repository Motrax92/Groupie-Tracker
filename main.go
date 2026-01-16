package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	"groupie-tracker/services"
	"groupie-tracker/ui"
)

func main() {
	// --- Création de l'application Fyne ---
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme()) // Bonus : Mode sombre par défaut
	
	myWindow := myApp.NewWindow("Groupie Trackers Pro")
	myWindow.Resize(fyne.NewSize(1200, 800))

	// --- Récupération des données ---
	// On appelle directement la fonction GetArtists du package services
	artists, err := services.GetArtists()
	if err != nil {
		log.Fatal("Erreur lors de la récupération des artistes :", err)
	}

	// --- UI Components ---
	// Note : Assurez-vous que vos fichiers dans le dossier ui/ 
	// acceptent bien ces arguments dans leurs constructeurs.
	searchBar := ui.NewSearchBar(artists)
	filterPanel := ui.NewFilterPanel(artists)
	artistList := ui.NewArtistList(artists)

	// --- Layout principal ---
	// On empile la barre de recherche et les filtres en haut
	topBar := container.NewVBox(
		searchBar.Render(),
		filterPanel.Render(),
	)

	// Border layout : le topBar en haut, la liste au centre
	content := container.NewBorder(
		topBar,              // Top
		nil,                 // Bottom
		nil,                 // Left
		nil,                 // Right
		artistList.Render(), // Center (Scrollable)
	)

	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}