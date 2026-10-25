package taller

const elevadoresTaller = 3

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
