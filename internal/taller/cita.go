package taller

import "time"

type Cita struct {
	inicio time.Time
	tipo   TipoReparacion
}

type TipoReparacion struct {
	nombreEnCatalogo string
	duracion         time.Duration
	usaElevador      bool
}
