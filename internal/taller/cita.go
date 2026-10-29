package taller

import "time"

type Cita struct {
	inicio time.Time
	tipo   TipoReparacion
}

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}
