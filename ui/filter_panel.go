package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewFilterPanel(onApply func(minYear, maxYear, memberCount int)) *fyne.Container {
	minYear := widget.NewSlider(1950, 2025)
	minYear.Value = 1950
	minYear.Step = 1
	minLabel := widget.NewLabel("Année min : 1950")
	minYear.OnChanged = func(v float64) {
		minLabel.SetText("Année min : " + fmt.Sprint(int(v)))
	}

	maxYear := widget.NewSlider(1950, 2025)
	maxYear.Value = 2025
	maxYear.Step = 1
	maxLabel := widget.NewLabel("Année max : 2025")
	maxYear.OnChanged = func(v float64) {
		maxLabel.SetText("Année max : " + fmt.Sprint(int(v)))
	}

	memberSlider := widget.NewSlider(1, 10)
	memberSlider.Value = 1
	memberSlider.Step = 1
	memberLabel := widget.NewLabel("Membres : 1")
	memberSlider.OnChanged = func(v float64) {
		memberLabel.SetText("Membres : " + fmt.Sprint(int(v)))
	}

	apply := widget.NewButton("Appliquer", func() {
		onApply(int(minYear.Value), int(maxYear.Value), int(memberSlider.Value))
	})

	reset := widget.NewButton("Réinitialiser", func() {
		minYear.Value = 1950
		maxYear.Value = 2025
		memberSlider.Value = 1

		minLabel.SetText("Année min : 1950")
		maxLabel.SetText("Année max : 2025")
		memberLabel.SetText("Membres : 1")

		onApply(1950, 2025, 1)
	})

	return container.NewVBox(
		widget.NewLabel("Filtres avancés"),
		minLabel,
		minYear,
		maxLabel,
		maxYear,
		memberLabel,
		memberSlider,
		container.NewHBox(apply, reset),
	)
}
