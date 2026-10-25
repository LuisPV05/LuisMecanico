package taller

import "errors"

const elevadoresTaller = 3

var (
	ErrCitaRepetida    = errors.New("La cita ya está apuntada en la libreta")
)

type Libreta struct {
	elevadores int
	citas      map[int]Cita
}

func NuevaLibreta() Libreta {
	return Libreta{elevadores: elevadoresTaller, citas: map[int]Cita{}}
}

func (l *Libreta) Apuntar(cita Cita) error {
	_, tipoValido := tiposReparacion[cita.tipo.nombre]
	if !tipoValido {
		return ErrTipoDesconocido
	}
	if l.citas == nil {
		*l = NuevaLibreta()
	}
	_, existe := l.citas[cita.id]
	if existe {
		return ErrCitaRepetida
	}
	l.citas[cita.id] = cita
	return nil
}
