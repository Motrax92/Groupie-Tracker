package main

import (
	"log"

	"groupie-tracker/api"
	"groupie-tracker/ui"
)

func main() {
	all, err := api.GetAllArtistFull()
	if err != nil {
		log.Fatalf("failed to load data from API: %v", err)
	}

	state := ui.NewApp(all)
	state.Window.ShowAndRun()
}