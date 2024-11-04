package scenes

import (
	"simulador/models"
	"simulador/poison"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type Escenario struct {
	window    fyne.Window
	contenido *fyne.Container
}

func (e *Escenario) RenderAutomovil(n *models.Parking) {
	for autoImagen := range n.ColocarAutomovil {
		e.contenido.Add(autoImagen)
		e.window.Canvas().Refresh(e.contenido)
	}
}

func (e *Escenario) IniciarSimulacion() {
	n := models.CrearParking(20)
	go poison.MetodoPoison(100, n)
	go e.RenderAutomovil(n)
}

func (e *Escenario) inicializarContenido() {
	bgImagen := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.png"))
	bgImagen.Resize(fyne.NewSize(800, 600))
	bgImagen.Move(fyne.NewPos(0, 0))

	e.contenido = container.NewWithoutLayout(bgImagen)
	e.window.SetContent(e.contenido)
}

func (e *Escenario) Renderizado() {
	e.inicializarContenido()
	e.IniciarSimulacion()
}

func NuevaEscena(window fyne.Window) *Escenario {
	escena := &Escenario{window: window}
	escena.Renderizado()
	return escena
}
