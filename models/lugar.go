package models

type Lugar struct {
	PosicionX float32
	PosicionY float32
	Ocupado   bool
}

func (l *Lugar) EstaDisponible() bool {
	return !l.Ocupado
}

func (l *Lugar) LiberarLugar() {
	l.Ocupado = false
}

func (l *Lugar) OcuparLugar() {
	l.Ocupado = true
}
