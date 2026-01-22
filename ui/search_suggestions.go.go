package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewSuggestionList() *fyne.Container {
	list := container.NewVBox()

	box := container.NewMax(list)
	box.Hide()

	return box
}

func UpdateSuggestionList(box *fyne.Container, suggestions []string, onSelect func(string)) {
	list := box.Objects[0].(*fyne.Container)
	list.Objects = nil

	if len(suggestions) == 0 {
		box.Hide()
		return
	}

	for _, s := range suggestions {
		text := s
		btn := widget.NewButton(text, func() {
			onSelect(text)
		})
		btn.Importance = widget.LowImportance
		list.Add(btn)
	}

	list.Refresh()
	box.Show()
}
