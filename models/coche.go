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
	// Bloquea hasta que hay un espacio disponible
	espacio := <-a.Parking.EspaciosParking // Espera un espacio

	// Ocupar el espacio
	a.Parking.LugaresParking[espacio.Index].OcuparLugar()
	a.modelo.Move(fyne.NewPos(a.Parking.LugaresParking[espacio.Index].PosicionX, a.Parking.LugaresParking[espacio.Index].PosicionY))
	a.modelo.Refresh()

	now := time.Now()
	horaLocal := now.Format("15:04:05")
	fechaLocal := now.Format("2006-01-02")
	fmt.Println("El vehículo", a.Identificador, "entró a las", horaLocal, "el", fechaLocal)
	time.Sleep(300 * time.Millisecond)

	// Simular tiempo en el parking
	TiempoEsperaTurno := rand.Intn(3) + 3
	time.Sleep(time.Duration(TiempoEsperaTurno) * time.Second)

	// Liberar el espacio
	a.Parking.LugaresParking[espacio.Index].LiberarLugar()
	a.modelo.Move(fyne.NewPos(410, 0))
	time.Sleep(300 * time.Millisecond)
	a.modelo.Move(fyne.NewPos(410, -150))
	a.modelo.Refresh()

	now = time.Now()
	horaSalida := now.Format("15:04:05")
	fmt.Println("El vehículo", a.Identificador, "salió a las", horaSalida, "el", fechaLocal)
	time.Sleep(300 * time.Millisecond)

	// Regresar el espacio al canal
	a.Parking.EspaciosParking <- espacio
}
