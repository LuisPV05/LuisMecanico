package taller

type Libreta struct {
	elevadores int
	citas      map[int]Cita
}

func NuevaLibreta() Libreta {
	return Libreta{elevadores: 3, citas: map[int]Cita{}}
}

func (l *Libreta) Apuntar(cita Cita) error {
	_, existe := l.citas[cita.id]
	if existe {
		return ErrCitaRepetida
	}
	l.citas[cita.id] = cita
	return nil
}
