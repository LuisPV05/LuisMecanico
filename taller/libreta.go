package taller

type Libreta struct {
	horaApertura int
	horaCierre   int
	elevadores   int
	citas        []Cita
}

func NuevaLibreta(citas []Cita) Libreta {
	return Libreta{horaApertura: 8, horaCierre: 21, elevadores: 3, citas: citas}
}
