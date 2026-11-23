package taller

import (
	"errors"
	"time"
)

var (
	ErrDomingo        = errors.New("el taller no abre en domingo")
	ErrFueraDeHorario = errors.New("la hora está fuera del horario del taller")
)

func validarInicio(inicio time.Time) error {
	if inicio.Weekday() == time.Sunday {
		return ErrDomingo
	}
	if inicio.Hour() < horaApertura || inicio.Hour() >= horaCierre {
		return ErrFueraDeHorario
	}
	return nil
}
