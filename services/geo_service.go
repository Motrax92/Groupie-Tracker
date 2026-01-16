package services

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

// OpenLocationMap convertit les lieux en coordonnées via une interface de carte
func OpenLocationMap(location string) {
	// Préparation de la requête pour l'API de carte
	address := fmt.Sprintf("https://www.google.com/maps/search/%s", url.QueryEscape(location))
	
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", address).Start()
	case "darwin":
		err = exec.Command("open", address).Start()
	default: // Linux
		err = exec.Command("xdg-open", address).Start()
	}
	if err != nil {
		fmt.Println("Erreur d'ouverture de la carte:", err)
	}
}