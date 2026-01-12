package ui

import (
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"groupie-tracker/core"
)

type MapWidget struct {
	fyne.CanvasObject
	img     *canvas.Image
	markers []core.Coordinates
}

func NewMapWidget() *MapWidget {
	img := canvas.NewImageFromFile("assets/world.png")
	img.FillMode = canvas.ImageFillContain

	return &MapWidget{
		CanvasObject: img,
		img:          img,
		markers:      []core.Coordinates{},
	}
}

func (m *MapWidget) SetMarkers(coords []core.Coordinates) {
	m.markers = coords
	canvas.Refresh(m.img)
}

func (m *MapWidget) DrawMarkers() []fyne.CanvasObject {
	var objs []fyne.CanvasObject

	for _, c := range m.markers {
		x, y := geoToPixel(c.Lat, c.Lon, float64(m.img.Size().Width), float64(m.img.Size().Height))

		dot := canvas.NewCircle(color.RGBA{255, 0, 0, 255})
		dot.Resize(fyne.NewSize(8, 8))
		dot.Move(fyne.NewPos(float32(x), float32(y)))

		objs = append(objs, dot)
	}

	return objs
}

// Projection simple (Plate Carrée)
func geoToPixel(lat, lon, width, height float64) (float64, float64) {
	x := (lon + 180.0) * (width / 360.0)
	y := (90.0 - lat) * (height / 180.0)
	return x, y
}
