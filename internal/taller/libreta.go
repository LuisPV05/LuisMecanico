package taller

import "time"

const (
	horaApertura     = 8
	horaCierre       = 21
	elevadoresTaller = 3
)

type Libreta struct {
	elevadores [elevadoresTaller]map[time.Time]Cita
	sinElevador []Cita
}
