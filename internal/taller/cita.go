package taller

import (
	"errors"
	"time"
)

var (
	ErrFueraDeHorario = errors.New("La cita empieza fuera del horario del taller")
	ErrDomingo        = errors.New("La cita no puede realizarse en domingo")
	ErrDuracionNoPositiva = errors.New("La duración de la reparación tiene que ser mayor que cero")
)

type Cita struct {
	inicio time.Time
	tipo   TipoReparacion
	id     int
}

const (
	horaApertura = 8
	horaCierre   = 21
)

func NuevaCita(id int, inicio time.Time, tipo TipoReparacion) (Cita, error) {
	if inicio.Weekday() == time.Sunday {
		return Cita{}, ErrDomingo
	}
	if inicio.Hour() < horaApertura || inicio.Hour() >= horaCierre {
		return Cita{}, ErrFueraDeHorario
	}
	if tipo.duracion <= 0 {
		return Cita{}, ErrDuracionNoPositiva
	}
	
	return Cita{id: id, inicio: inicio, tipo: tipo}, nil
}
