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

func ShowArtistDetails(app fyne.App, a models.Artist) {

	win := app.NewWindow(a.Name)
	win.Resize(fyne.NewSize(600, 500))

	var img *canvas.Image
	if a.Image != "" {
		img = canvas.NewImageFromURI(storage.NewURI(a.Image))
	} else {
		img = canvas.NewImageFromResource(nil)
	}
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(300, 300))

	name := widget.NewLabelWithStyle(
		a.Name,
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	creation := widget.NewLabel(fmt.Sprintf("Année de création : %d", a.CreationDate))
	album := widget.NewLabel("Premier album : " + a.FirstAlbum)

	membersTitle := widget.NewLabel("Membres :")
	memberList := container.NewVBox()
	for _, m := range a.Members {
		memberList.Add(widget.NewLabel("• " + m))
	}

	closeBtn := widget.NewButton("Fermer", func() {
		win.Close()
	})

	content := container.NewVBox(
		img,
		name,
		creation,
		album,
		membersTitle,
		memberList,
		closeBtn,
	)

	win.SetContent(container.NewVScroll(content))
	win.Show()
}
