package poison

import (
	"simulador/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"math/rand"
	"time"
)

// MetodoPoison simula la llegada de automóviles al estacionamiento.
func MetodoPoison(cantidad int, parking *models.Parking) {
	parking.EspaciosParking <- true

	for i := 0; i < cantidad; i++ {
		automovilImagen := crearAutomovilImagen()
		automovilNuevo := models.CrearAutomovil(parking, automovilImagen)
		automovilNuevo.Identificador = i + 1

		// Agregar la imagen del automóvil al canal de colocación
		parking.ColocarAutomovil <- automovilImagen
		time.Sleep(200 * time.Millisecond)

		go automovilNuevo.MoverCarro()

		// Espera exponencial entre llegadas de automóviles
		time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
	}
}

// crearAutomovilImagen crea y configura la imagen del automóvil.
func crearAutomovilImagen() *canvas.Image {
	automovilImagen := canvas.NewImageFromURI(storage.NewFileURI("./assets/coche.png"))
	automovilImagen.Resize(fyne.NewSize(50, 80)) // Tamaño de la imagen
	automovilImagen.Move(fyne.NewPos(410, 0))    // Posición inicial
	return automovilImagen
}
