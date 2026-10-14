package taller

import "errors"

var (
	ErrFueraDeHorario  = errors.New("la cita empieza fuera del horario del taller")
	ErrDomingo         = errors.New("la cita no puede realizarse en domingo")
	ErrTipoDesconocido = errors.New("el tipo de reparación que se ha indicado no está en la lista")
	ErrCitaRepetida    = errors.New("la cita ya está apuntada en la libreta")
)
