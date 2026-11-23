package taller

import (
	"errors"
	"testing"
	"time"
)

func citasALaVez(cantidad int, inicio time.Time, minutos int) []Cita {
	citas := make([]Cita, cantidad)
	for i := range citas {
		citas[i] = citaDe(inicio, minutos)
	}
	return citas
}

func TestComprobarElevador(t *testing.T) {
	nueva := citaDe(fecha(23, 9, 0), 60)
	casos := []struct {
		nombre  string
		libreta Libreta
		nueva   Cita
		quiero  error
	}{
		{"libreta vacía", Libreta{}, nueva, nil},
		{"dos citas a la vez", Libreta{citas: citasALaVez(2, fecha(23, 9, 0), 60)}, nueva, nil},
		{"tres citas a la vez", Libreta{citas: citasALaVez(3, fecha(23, 9, 0), 60)}, nueva, ErrSinElevadorLibre},
		{"tres citas a la vez y la nueva no usa elevador", Libreta{citas: citasALaVez(3, fecha(23, 9, 0), 60)}, citaSinElevador(fecha(23, 9, 0), 60), nil},
		{"tres citas que acaban justo cuando empieza la nueva", Libreta{citas: citasALaVez(3, fecha(23, 8, 0), 60)}, nueva, nil},
		{"tres citas en otro momento", Libreta{citas: citasALaVez(3, fecha(23, 11, 0), 60)}, nueva, nil},
		{"tres citas a la vez que empiezan antes y siguen dentro", Libreta{citas: citasALaVez(3, fecha(23, 8, 30), 60)}, nueva, ErrSinElevadorLibre},
		{"tres citas seguidas dentro de la nueva", Libreta{citas: []Cita{
			citaDe(fecha(23, 9, 0), 20),
			citaDe(fecha(23, 9, 20), 20),
			citaDe(fecha(23, 9, 40), 20),
		}}, nueva, nil},
		{"tres citas de la noche anterior", Libreta{citas: citasALaVez(3, fecha(23, 20, 30), 45)}, citaDe(fecha(24, 8, 0), 30), ErrSinElevadorLibre},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := c.libreta.comprobarElevador(c.nueva)
			if !errors.Is(err, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", err, c.quiero)
			}
		})
	}
}
