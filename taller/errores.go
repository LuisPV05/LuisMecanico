package taller

import "errors"

var (
	ErrFueraDeHorario  = errors.New("La cita empieza fuera del horario del taller")
	ErrDomingo         = errors.New("La cita no puede realizarse en domingo")
	ErrTipoDesconocido = errors.New("El tipo de reparación que se ha indicado no está en La lista")
	ErrCitaRepetida    = errors.New("La cita ya está apuntada en la libreta")
)
