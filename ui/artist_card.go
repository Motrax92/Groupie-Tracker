package ui

import (
	"fmt"

	"groupie-tracker/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewArtistCard(a models.Artist) *fyne.Container {
	return container.NewVBox(
		widget.NewLabelWithStyle(
			a.Name,
			fyne.TextAlignLeading,
			fyne.TextStyle{Bold: true},
		),
		widget.NewLabel("Membres : "+join(a.Members)),
		widget.NewLabel(fmt.Sprintf("Création : %d", a.CreationDate)),
		widget.NewLabel("Premier album : "+a.FirstAlbum),
	)
}

func join(list []string) string {
	result := ""
	for i, s := range list {
		if i > 0 {
			result += ", "
		}
		result += s
	}
	return result
}
