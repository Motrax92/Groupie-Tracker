package services

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

func OpenLocationMap(location string) {
	if location == "" {
		return
	}

	address := fmt.Sprintf("https://www.google.com/maps/search/%s", url.QueryEscape(location))

	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", address).Start()
	case "darwin":
		err = exec.Command("open", address).Start()
	default:
		err = exec.Command("xdg-open", address).Start()
	}

	if err != nil {
		fmt.Println("Erreur d'ouverture de la carte:", err)
	}
}
