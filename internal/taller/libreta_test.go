package taller

import (
	"testing"
	"time"
)

func citaSinElevador(inicio time.Time, minutos int) Cita {
	return Cita{inicio: inicio, tipo: tipo("Reemplazo de la batería", minutos, false)}
}

func TestElevadoresOcupados(t *testing.T) {
	inicio, fin := fecha(23, 9, 0), fecha(23, 10, 0)
	casos := []struct {
		nombre  string
		libreta Libreta
		quiero  int
	}{
		{"libreta vacía", Libreta{}, 0},
		{"una cita en el tramo", Libreta{citas: []Cita{citaDe(fecha(23, 9, 15), 30)}}, 1},
		{"tres citas en el tramo", Libreta{citas: []Cita{
			citaDe(fecha(23, 9, 0), 60),
			citaDe(fecha(23, 8, 30), 60),
			citaDe(fecha(23, 9, 30), 120),
		}}, 3},
		{"cita anterior al tramo", Libreta{citas: []Cita{citaDe(fecha(23, 8, 0), 30)}}, 0},
		{"cita contigua por delante", Libreta{citas: []Cita{citaDe(fecha(23, 8, 0), 60)}}, 0},
		{"cita contigua por detrás", Libreta{citas: []Cita{citaDe(fecha(23, 10, 0), 30)}}, 0},
		{"cita posterior al tramo", Libreta{citas: []Cita{citaDe(fecha(23, 11, 0), 30)}}, 0},
		{"cita que no usa elevador", Libreta{citas: []Cita{citaSinElevador(fecha(23, 9, 0), 30)}}, 0},
		{"mezcla de citas", Libreta{citas: []Cita{
			citaDe(fecha(23, 9, 0), 60),
			citaDe(fecha(23, 9, 30), 30),
			citaSinElevador(fecha(23, 9, 15), 15),
			citaDe(fecha(23, 11, 0), 30),
		}}, 2},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenido := c.libreta.elevadoresOcupados(inicio, fin)
			if obtenido != c.quiero {
				t.Errorf("obtenido %d, quiero %d", obtenido, c.quiero)
			}
		})
	}
}

func TestElevadoresOcupadosDeNoche(t *testing.T) {
	libreta := Libreta{citas: []Cita{citaDe(fecha(23, 20, 30), 45)}}
	obtenido := libreta.elevadoresOcupados(fecha(24, 8, 0), fecha(24, 9, 0))
	if obtenido != 1 {
		t.Errorf("obtenido %d, quiero 1", obtenido)
	}
}
