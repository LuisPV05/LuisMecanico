package taller

import (
	"time"
)

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}
