package taller

type Libreta struct {
	elevadores int
	citas      []Cita
}

func NuevaLibreta(citas []Cita) Libreta {
	return Libreta{elevadores: 3, citas: citas}
}
