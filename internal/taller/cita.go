package taller

import (
	"errors"
	"time"
)

var (
	ErrFueraDeHorario     = errors.New("La cita empieza fuera del horario del taller")
	ErrDomingo            = errors.New("La cita no puede realizarse en domingo")
	ErrDuracionNoPositiva = errors.New("La duración de la reparación tiene que ser mayor que cero")
)

type Cita struct {
	inicio time.Time
	tipo   TipoReparacion
}

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}
