package horario_test

import (
	"errors"
	"testing"
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/horario"
	"github.com/LuisPV05/LuisMecanico/internal/taller/pruebas"
)

func TestValidarInicio(t *testing.T) {
	casos := []struct {
		nombre string
		inicio time.Time
		quiero error
	}{
		{"lunes 8:00", pruebas.Fecha(23, 8, 0), nil},
		{"sábado 20:59", pruebas.Fecha(28, 20, 59), nil},
		{"lunes 7:59", pruebas.Fecha(23, 7, 59), horario.ErrFueraDeHorario},
		{"lunes 21:00", pruebas.Fecha(23, 21, 0), horario.ErrFueraDeHorario},
		{"domingo 10:00", pruebas.Fecha(22, 10, 0), horario.ErrDomingo},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := horario.ValidarInicio(c.inicio)
			if !errors.Is(err, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", err, c.quiero)
			}
		})
	}
}