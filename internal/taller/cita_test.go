package taller

import (
	"errors"
	"reflect"
	"testing"
)

func TestNuevaCita(t *testing.T) {
	catalogo := []TipoReparacion{
		tipo("Mantenimiento básico", 45, true),
		tipo("Cambio de neumáticos y alineación", 60, true),
		tipo("Sin duración", 0, true),
		tipo("Duración negativa", -10, true),
	}
	inicio := fecha(23, 8, 0)
	casos := []struct {
		nombre    string
		quiero    Cita
		quieroErr error
	}{
		{"Mantenimiento básico", Cita{inicio: inicio, tipo: tipo("Mantenimiento básico", 45, true)}, nil},
		{"Cambio de neumáticos y alineación", Cita{inicio: inicio, tipo: tipo("Cambio de neumáticos y alineación", 60, true)}, nil},
		{"Lavado de coche", Cita{}, ErrTipoDesconocido},
		{"Sin duración", Cita{}, ErrDuracionInvalida},
		{"Duración negativa", Cita{}, ErrDuracionInvalida},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenida, err := nuevaCita(catalogo, c.nombre, inicio)
			if !errors.Is(err, c.quieroErr) {
				t.Errorf("error obtenido %v, quiero %v", err, c.quieroErr)
			}
			if !reflect.DeepEqual(obtenida, c.quiero) {
				t.Errorf("obtenida %v, quiero %v", obtenida, c.quiero)
			}
		})
	}
}
