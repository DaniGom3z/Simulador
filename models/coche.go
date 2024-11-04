package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Automovil struct {
	Parking         *Parking
	Identificador   int
	posicionParking int
	modelo          *canvas.Image
}

func CrearAutomovil(p *Parking, m *canvas.Image) *Automovil {
	return &Automovil{
		Parking: p,
		modelo:  m,
	}
}

func (a *Automovil) MoverCarro() {
	a.Parking.EspaciosParking <- true
	a.Parking.M.Lock()

	// Buscar un lugar disponible
	for i := 0; i < len(a.Parking.EspaciosParking); i++ {
		if !a.Parking.LugaresParking[i].Ocupado {
			a.modelo.Move(fyne.NewPos(a.Parking.LugaresParking[i].PosicionX, a.Parking.LugaresParking[i].PosicionY))
			a.modelo.Refresh()

			// Ocupa el lugar
			a.posicionParking = i
			a.Parking.LugaresParking[i].Ocupado = true
			break
		}
	}

	now := time.Now()
	horaLocal := now.Format("15:04:05")
	fechaLocal := now.Format("2006-01-02")

	fmt.Println("El vehiculo", a.Identificador, "Entró a las", horaLocal, "el", fechaLocal)
	time.Sleep(300 * time.Millisecond)

	// Desbloquear Mutex
	a.Parking.M.Unlock()

	TiempoEsperaTurno := rand.Intn(6) + 5 
	time.Sleep(time.Duration(TiempoEsperaTurno) * time.Second)

	// Volver a bloquear el Mutex antes de salir
	a.Parking.M.Lock()

	// Liberar el espacio de automóviles
	<-a.Parking.EspaciosParking
	a.Parking.LugaresParking[a.posicionParking].Ocupado = false
	a.modelo.Move(fyne.NewPos(350, 0))
	a.modelo.Refresh()

	// Registrar la salida después de la espera
	now = time.Now() // Obtener la hora actual nuevamente para la salida
	horaSalida := now.Format("15:04:05")
	fmt.Println("El vehiculo", a.Identificador, "Salió a las", horaSalida, "el", fechaLocal)
	time.Sleep(300 * time.Millisecond)

	// Desbloquear el Mutex
	a.Parking.M.Unlock()
}
