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

func (c Cita) ID() int {
	return c.id
}

func (c Cita) Inicio() time.Time {
	return c.inicio
}

func (c Cita) Tipo() string {
	return c.tipo
}
