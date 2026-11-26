package elevadores_test

import (
	"errors"
	"testing"

	"github.com/LuisPV05/LuisMecanico/internal/taller/cita"
	"github.com/LuisPV05/LuisMecanico/internal/taller/elevadores"
	"github.com/LuisPV05/LuisMecanico/internal/taller/pruebas"
)

func TestOcupados(t *testing.T) {
	inicio, fin := pruebas.Fecha(23, 9, 0), pruebas.Fecha(23, 10, 0)
	casos := []struct {
		nombre string
		citas  []cita.Cita
		quiero int
	}{
		{"sin citas", nil, 0},
		{"una cita en el tramo", []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 9, 15), 30)}, 1},
		{"tres citas a la vez", []cita.Cita{
			pruebas.CitaDe(pruebas.Fecha(23, 9, 0), 60),
			pruebas.CitaDe(pruebas.Fecha(23, 8, 30), 60),
			pruebas.CitaDe(pruebas.Fecha(23, 9, 15), 60),
		}, 3},
		{"tres citas seguidas, una detrás de otra", []cita.Cita{
			pruebas.CitaDe(pruebas.Fecha(23, 9, 0), 20),
			pruebas.CitaDe(pruebas.Fecha(23, 9, 20), 20),
			pruebas.CitaDe(pruebas.Fecha(23, 9, 40), 20),
		}, 1},
		{"cita que empieza antes del tramo y sigue dentro", []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 8, 30), 60)}, 1},
		{"cita anterior al tramo", []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 8, 0), 30)}, 0},
		{"cita contigua por delante", []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 8, 0), 60)}, 0},
		{"cita contigua por detrás", []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 10, 0), 30)}, 0},
		{"cita posterior al tramo", []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 11, 0), 30)}, 0},
		{"cita que no usa elevador", []cita.Cita{pruebas.CitaSinElevador(pruebas.Fecha(23, 9, 0), 30)}, 0},
		{"mezcla de citas", []cita.Cita{
			pruebas.CitaDe(pruebas.Fecha(23, 9, 0), 60),
			pruebas.CitaDe(pruebas.Fecha(23, 9, 30), 30),
			pruebas.CitaSinElevador(pruebas.Fecha(23, 9, 15), 15),
			pruebas.CitaDe(pruebas.Fecha(23, 11, 0), 30),
		}, 2},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenido := elevadores.Ocupados(c.citas, inicio, fin)
			if obtenido != c.quiero {
				t.Errorf("obtenido %d, quiero %d", obtenido, c.quiero)
			}
		})
	}
}

func TestOcupadosDeNoche(t *testing.T) {
	citas := []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 20, 30), 45)}
	obtenido := elevadores.Ocupados(citas, pruebas.Fecha(24, 8, 0), pruebas.Fecha(24, 9, 0))
	if obtenido != 1 {
		t.Errorf("obtenido %d, quiero 1", obtenido)
	}
}

func TestComprobar(t *testing.T) {
	nueva := pruebas.CitaDe(pruebas.Fecha(23, 9, 0), 60)
	casos := []struct {
		nombre string
		citas  []cita.Cita
		nueva  cita.Cita
		quiero error
	}{
		{"sin citas", nil, nueva, nil},
		{"dos citas a la vez", pruebas.CitasALaVez(2, pruebas.Fecha(23, 9, 0), 60), nueva, nil},
		{"tres citas a la vez", pruebas.CitasALaVez(3, pruebas.Fecha(23, 9, 0), 60), nueva, elevadores.ErrSinElevadorLibre},
		{"tres citas a la vez y la nueva no usa elevador", pruebas.CitasALaVez(3, pruebas.Fecha(23, 9, 0), 60), pruebas.CitaSinElevador(pruebas.Fecha(23, 9, 0), 60), nil},
		{"tres citas que acaban justo cuando empieza la nueva", pruebas.CitasALaVez(3, pruebas.Fecha(23, 8, 0), 60), nueva, nil},
		{"tres citas en otro momento", pruebas.CitasALaVez(3, pruebas.Fecha(23, 11, 0), 60), nueva, nil},
		{"tres citas a la vez que empiezan antes y siguen dentro", pruebas.CitasALaVez(3, pruebas.Fecha(23, 8, 30), 60), nueva, elevadores.ErrSinElevadorLibre},
		{"tres citas seguidas dentro de la nueva", []cita.Cita{
			pruebas.CitaDe(pruebas.Fecha(23, 9, 0), 20),
			pruebas.CitaDe(pruebas.Fecha(23, 9, 20), 20),
			pruebas.CitaDe(pruebas.Fecha(23, 9, 40), 20),
		}, nueva, nil},
		{"tres citas de la noche anterior", pruebas.CitasALaVez(3, pruebas.Fecha(23, 20, 30), 45), pruebas.CitaDe(pruebas.Fecha(24, 8, 0), 30), elevadores.ErrSinElevadorLibre},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := elevadores.Comprobar(c.citas, c.nueva)
			if !errors.Is(err, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", err, c.quiero)
			}
		})
	}
}
