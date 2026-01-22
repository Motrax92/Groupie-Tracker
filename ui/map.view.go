package ui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewMapView() fyne.CanvasObject {
	u, _ := url.Parse("https://www.openstreetmap.org")
	web := widget.NewHyperlink("Carte des concerts", u)

	return web
}
