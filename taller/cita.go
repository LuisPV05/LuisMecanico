package taller

import "time"

type Cita struct {
	id     int
	inicio time.Time
	tipo   string
}

func NuevaCita(id int, inicio time.Time, tipo string) Cita {
	return Cita{id: id, inicio: inicio, tipo: tipo}
}
