package observers

type Sujeto interface {
	AgregarObservador(o Observador)
	RemoverObservador(o Observador)
	NotificarObservadores()
}
