package taller

type Libreta struct {
	elevadores int
	citas      []Cita
}

func NuevaLibreta(citas []Cita) (Libreta, error) {
	for i := 0; i < len(citas); i++ {
		for j := i + 1; j < len(citas); j++ {
			if citas[i].id == citas[j].id {
				return Libreta{}, ErrCitaRepetida
			}
		}
	}
	return Libreta{elevadores: 3, citas: citas}, nil
}
