package taller

import "time"

const (
	horaApertura     = 8
	horaCierre       = 21
	elevadoresTaller = 3
)

type Libreta struct {
	citas []Cita
}

func (l Libreta) elevadoresOcupados(inicio, fin time.Time) int {
	ocupados := 0
	for _, c := range l.citas {
		if c.tipo.usaElevador && c.solapaConTramo(inicio, fin) {
			ocupados++
		}
	}
	return ocupados
}