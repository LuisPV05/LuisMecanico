package taller

import (
	"testing"
	"time"
)

func TestFinDeCita(t *testing.T) {
	casos := []struct {
		nombre   string
		inicio   time.Time
		duracion time.Duration
		quiero   time.Time
	}{
		{"lunes 8:00 + 45 min", fecha(23, 8, 0), 45 * time.Minute, fecha(23, 8, 45)},
		{"lunes 20:30 + 45 min", fecha(23, 20, 30), 45 * time.Minute, fecha(24, 8, 15)},
		{"viernes 20:30 + 45 min", fecha(27, 20, 30), 45 * time.Minute, fecha(28, 8, 15)},
		{"sábado 20:30 + 45 min", fecha(28, 20, 30), 45 * time.Minute, fecha(30, 8, 15)},
		{"lunes 8:00 + 13 h", fecha(23, 8, 0), 13 * time.Hour, fecha(23, 21, 0)},
		{"lunes 8:00 + 39 h", fecha(23, 8, 0), 39 * time.Hour, fecha(25, 21, 0)},
		{"viernes 8:00 + 39 h", fecha(27, 8, 0), 39 * time.Hour, fecha(30, 21, 0)},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenido := finDeCita(c.inicio, c.duracion)
			if !obtenido.Equal(c.quiero) {
				t.Errorf("obtenido %v, quiero %v", obtenido, c.quiero)
			}
		})
	}
}
