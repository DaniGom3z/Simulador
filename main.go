package main

import (
	"simulador/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	interfaz := app.New()
	ventana := interfaz.NewWindow("Estacioanmiento    10 por hora o fracción")

	ventana.CenterOnScreen()
	ventana.SetFixedSize(true)
	ventana.Resize(fyne.NewSize(800, 600))
	scenes.NuevaEscena(ventana)
	ventana.ShowAndRun()
}
