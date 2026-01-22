package ui

import "fyne.io/fyne/v2/widget"

func NewSearchBar(onChange func(string)) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Rechercher un artiste, un membre, une date…")
	entry.OnChanged = onChange
	return entry
}
