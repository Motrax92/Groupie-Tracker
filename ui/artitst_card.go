package ui

import (
	"fmt"

	"groupie-tracker/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func NewArtistCard(a models.Artist, onClick func()) fyne.CanvasObject {
	var img *canvas.Image

	if a.Image != "" {
		uri := storage.NewURI(a.Image)
		img = canvas.NewImageFromURI(uri)
	} else {
		img = canvas.NewImageFromResource(nil)
	}

	img.SetMinSize(fyne.NewSize(64, 64))
	img.FillMode = canvas.ImageFillContain

	nameLabel := widget.NewLabelWithStyle(
		a.Name,
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	subLabel := widget.NewLabel(
		fmt.Sprintf("Création : %d  •  %d membres", a.CreationDate, len(a.Members)),
	)

	albumLabel := widget.NewLabel("Premier album : " + a.FirstAlbum)

	detailsBtn := widget.NewButton("Voir détails", func() {
		onClick()
	})

	textCol := container.NewVBox(
		nameLabel,
		subLabel,
		albumLabel,
		detailsBtn,
	)

	row := container.NewHBox(
		img,
		textCol,
	)

	card := widget.NewCard("", "", row)
	return card
}
