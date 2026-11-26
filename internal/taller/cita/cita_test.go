package cita_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/catalogo"
	"github.com/LuisPV05/LuisMecanico/internal/taller/cita"
	"github.com/LuisPV05/LuisMecanico/internal/taller/pruebas"
)

func tiposDePrueba() []catalogo.TipoReparacion {
	return []catalogo.TipoReparacion{
		pruebas.Tipo("Mantenimiento básico", 45, true),
		pruebas.Tipo("Cambio de neumáticos y alineación", 60, true),
		pruebas.Tipo("Sin duración", 0, true),
		pruebas.Tipo("Duración negativa", -10, true),
	}
}

func TestNueva(t *testing.T) {
	casos := []struct {
		nombre string
		quiero error
	}{
		{"Mantenimiento básico", nil},
		{"Cambio de neumáticos y alineación", nil},
		{"Lavado de coche", catalogo.ErrTipoDesconocido},
		{"Sin duración", catalogo.ErrDuracionInvalida},
		{"Duración negativa", catalogo.ErrDuracionInvalida},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenida, err := cita.Nueva(tiposDePrueba(), c.nombre, pruebas.Fecha(23, 8, 0))
			if !errors.Is(err, c.quiero) {
				t.Errorf("error obtenido %v, quiero %v", err, c.quiero)
			}
			if err != nil && !reflect.DeepEqual(obtenida, cita.Cita{}) {
				t.Errorf("la cita debería estar vacía, obtenida %v", obtenida)
			}
		})
	}
}

func TestNuevaGuardaElInicioYElTipo(t *testing.T) {
	tipos := tiposDePrueba()
	inicio := pruebas.Fecha(23, 8, 0)
	obtenida, err := cita.Nueva(tipos, "Mantenimiento básico", inicio)
	if err != nil {
		t.Fatal(err)
	}
	if !obtenida.Inicio().Equal(inicio) {
		t.Errorf("inicio obtenido %v, quiero %v", obtenida.Inicio(), inicio)
	}
	if !reflect.DeepEqual(obtenida.Tipo(), tipos[0]) {
		t.Errorf("tipo obtenido %v, quiero %v", obtenida.Tipo(), tipos[0])
	}
}

func TestFin(t *testing.T) {
	casos := []struct {
		nombre string
		cita   cita.Cita
		quiero time.Time
	}{
		{"cabe en el día", pruebas.CitaDe(pruebas.Fecha(23, 8, 0), 45), pruebas.Fecha(23, 8, 45)},
		{"continúa al día siguiente", pruebas.CitaDe(pruebas.Fecha(23, 20, 30), 45), pruebas.Fecha(24, 8, 15)},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenido := c.cita.Fin()
			if !obtenido.Equal(c.quiero) {
				t.Errorf("obtenido %v, quiero %v", obtenido, c.quiero)
			}
		})
	}
}

func TestOcupaElevadorEn(t *testing.T) {
	inicio := pruebas.Fecha(23, 9, 0)
	conElevador := pruebas.CitaDe(inicio, 60)
	sinElevador := pruebas.CitaSinElevador(inicio, 60)
	casos := []struct {
		nombre  string
		cita    cita.Cita
		momento time.Time
		quiero  bool
	}{
		{"antes de empezar", conElevador, pruebas.Fecha(23, 8, 59), false},
		{"justo al empezar", conElevador, pruebas.Fecha(23, 9, 0), true},
		{"a mitad", conElevador, pruebas.Fecha(23, 9, 30), true},
		{"justo al acabar", conElevador, pruebas.Fecha(23, 10, 0), false},
		{"sin elevador a mitad", sinElevador, pruebas.Fecha(23, 9, 30), false},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			if obtenido := c.cita.OcupaElevadorEn(c.momento); obtenido != c.quiero {
				t.Errorf("obtenido %v, quiero %v", obtenido, c.quiero)
			}
		})
	}
}
