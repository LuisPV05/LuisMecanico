package taller

import (
	"errors"
	"testing"
	"time"
)

func fecha(dia, hora, minuto int) time.Time {
	return time.Date(2026, time.November, dia, hora, minuto, 0, 0, time.UTC)
}

func TestValidarInicio(t *testing.T) {
	casos := []struct {
		nombre string
		inicio time.Time
		quiero error
	}{
		{"lunes 8:00", fecha(23, 8, 0), nil},
		{"sábado 20:59", fecha(28, 20, 59), nil},
		{"lunes 7:59", fecha(23, 7, 59), ErrFueraDeHorario},
		{"lunes 21:00", fecha(23, 21, 0), ErrFueraDeHorario},
		{"domingo 10:00", fecha(22, 10, 0), ErrDomingo},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := validarInicio(c.inicio)
			if !errors.Is(err, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", err, c.quiero)
			}
		})
	}
}
