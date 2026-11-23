package taller

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func citaDe(inicio time.Time, minutos int) Cita {
	return Cita{inicio: inicio, tipo: tipo("Mantenimiento básico", minutos, true)}
}

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

func TestFinDeLaCita(t *testing.T) {
	casos := []struct {
		nombre string
		cita   Cita
		quiero time.Time
	}{
		{"cabe en el día", citaDe(fecha(23, 8, 0), 45), fecha(23, 8, 45)},
		{"continúa al día siguiente", citaDe(fecha(23, 20, 30), 45), fecha(24, 8, 15)},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenido := c.cita.fin()
			if !obtenido.Equal(c.quiero) {
				t.Errorf("obtenido %v, quiero %v", obtenido, c.quiero)
			}
		})
	}
}

func TestSolapan(t *testing.T) {
	base := citaDe(fecha(23, 9, 0), 60)
	nocturna := citaDe(fecha(23, 20, 30), 45)
	casos := []struct {
		nombre string
		a, b   Cita
		quiero bool
	}{
		{"termina antes de que empiece", base, citaDe(fecha(23, 8, 0), 30), false},
		{"contigua por delante", base, citaDe(fecha(23, 8, 0), 60), false},
		{"solapa por el inicio", base, citaDe(fecha(23, 8, 30), 60), true},
		{"está dentro", base, citaDe(fecha(23, 9, 15), 30), true},
		{"es la misma", base, citaDe(fecha(23, 9, 0), 60), true},
		{"solapa por el final", base, citaDe(fecha(23, 9, 30), 60), true},
		{"contigua por detrás", base, citaDe(fecha(23, 10, 0), 30), false},
		{"empieza después", base, citaDe(fecha(23, 11, 0), 30), false},
		{"coincide en la noche de una cita larga", nocturna, citaDe(fecha(24, 8, 0), 30), true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			if obtenido := solapan(c.a, c.b); obtenido != c.quiero {
				t.Errorf("solapan(a, b): obtenido %v, quiero %v", obtenido, c.quiero)
			}
			if obtenido := solapan(c.b, c.a); obtenido != c.quiero {
				t.Errorf("solapan(b, a): obtenido %v, quiero %v", obtenido, c.quiero)
			}
		})
	}
}
