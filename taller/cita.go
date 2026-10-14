package taller

import "time"

type Cita struct {
	id     int
	inicio time.Time
	tipo   TipoReparacion
}

func NuevaCita(id int, inicio time.Time, nombreTipo string) (Cita, error) {
	tipo, ok := tiposReparacion[nombreTipo]
	if !ok {
		return Cita{}, ErrTipoDesconocido
	}
	return Cita{id: id, inicio: inicio, tipo: tipo}, nil
}
