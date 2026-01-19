package ui

import "fyne.io/fyne/v2/widget"

func NewSearchBar(onChange func(string)) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Rechercher un artiste ou un membre")
	entry.OnChanged = onChange
	return entry
}
