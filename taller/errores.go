package taller

import "errors"

var (
	ErrFueraDeHorario  = errors.New("la cita empieza fuera del horario del taller")
	ErrDomingo         = errors.New("la cita no puede realizarse en domingo")
)
