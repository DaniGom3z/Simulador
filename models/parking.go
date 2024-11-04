package models

import (
	"fyne.io/fyne/v2/canvas"
	"sync"
)

type Lugar struct {
	PosicionX float32
	PosicionY float32
	Ocupado   bool
}

type Parking struct {
	EspaciosParking  chan bool
	LugaresParking   []Lugar
	ColocarAutomovil chan *canvas.Image
	M                sync.Mutex
}

func CrearParking(nP int) *Parking {
	coordenadas := []struct {
		x float32
		y float32
	}{
		{170, 100}, {270, 100}, {370, 100}, {470, 100}, {570, 100},
		{170, 205}, {270, 205}, {370, 205}, {470, 205}, {570, 205},
		{170, 330}, {270, 330}, {370, 330}, {470, 330}, {570, 330},
		{170, 425}, {270, 425}, {370, 425}, {470, 425}, {570, 425},
	}

	lugares := make([]Lugar, len(coordenadas))
	for i, coord := range coordenadas {
		lugares[i] = Lugar{
			PosicionX: coord.x,
			PosicionY: coord.y,
			Ocupado:   false,
		}
	}

	return &Parking{
		EspaciosParking:  make(chan bool, nP+1),
		ColocarAutomovil: make(chan *canvas.Image, 100),
		LugaresParking:   lugares,
	}
}
