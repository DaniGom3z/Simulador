package poison

import (
	"simulador/models"
	"time"
	"math/rand"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

func MetodoPoison(cantidad int, parking *models.Parking) {
	for i := 0; i < cantidad; i++ {
		automovilImagen := crearAutomovilImagen()
		automovilNuevo := models.CrearAutomovil(parking, automovilImagen)
		automovilNuevo.Identificador = i + 1

		// Agregar la imagen del automóvil al canal de colocación
		parking.ColocarAutomovil <- automovilImagen
		time.Sleep(2000 * time.Millisecond)

		// Movimiento del coche hasta el lugar del estacionamiento
		go automovilNuevo.MoverCarro()

		time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))

	}
}

func crearAutomovilImagen() *canvas.Image {
	automovilImagen := canvas.NewImageFromURI(storage.NewFileURI("./assets/coche.png"))
	automovilImagen.Resize(fyne.NewSize(50, 80))
	automovilImagen.Move(fyne.NewPos(350,0))
	return automovilImagen
}
