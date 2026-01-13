package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"groupie-tracker/core"
)

type MapWidget struct {
	container *fyne.Container
	image     *canvas.Image
	markers   []core.Coordinates
}

func NewMapWidget() *MapWidget {
	img := canvas.NewImageFromFile("assets/world.png")
	img.FillMode = canvas.ImageFillContain

	c := container.NewWithoutLayout(img)

	return &MapWidget{
		container: c,
		image:     img,
		markers:   []core.Coordinates{},
	}
}

func (m *MapWidget) Object() fyne.CanvasObject {
	return m.container
}

func (m *MapWidget) SetMarkers(coords []core.Coordinates) {
	m.markers = coords
	m.container.Objects = []fyne.CanvasObject{m.image}

	for _, c := range m.markers {
		x, y := geoToPixel(
			c.Lat,
			c.Lon,
			float64(m.image.Size().Width),
			float64(m.image.Size().Height),
		)

		dot := canvas.NewCircle(color.RGBA{255, 0, 0, 255})
		dot.Resize(fyne.NewSize(8, 8))
		dot.Move(fyne.NewPos(float32(x), float32(y)))

		m.container.Add(dot)
	}

	m.container.Refresh()
}

func geoToPixel(lat, lon, width, height float64) (float64, float64) {
	x := (lon + 180.0) * (width / 360.0)
	y := (90.0 - lat) * (height / 180.0)
	return x, y
}
