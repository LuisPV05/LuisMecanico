package catalogo

import "time"

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}

func NuevoTipo(nombre string, duracion time.Duration, usaElevador bool) TipoReparacion {
	return TipoReparacion{nombre: nombre, duracion: duracion, usaElevador: usaElevador}
}

func (t TipoReparacion) Nombre() string {
	return t.nombre
}

func (t TipoReparacion) Duracion() time.Duration {
	return t.duracion
}

func (t TipoReparacion) UsaElevador() bool {
	return t.usaElevador
}
