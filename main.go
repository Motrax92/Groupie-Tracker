package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	"groupie-tracker/services"
	"groupie-tracker/ui"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	artists, err := services.GetArtists()
	if err != nil {
		log.Fatal(err)
	}

	window := ui.NewAppWindow(myApp, artists)
	window.ShowAndRun()
}
