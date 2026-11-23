package taller

import (
	"errors"
	"reflect"
	"slices"
	"testing"
	"time"
)

func catalogoDePrueba() []TipoReparacion {
	return []TipoReparacion{
		tipo("Mantenimiento básico", 45, true),
		tipo("Cambio de neumáticos y alineación", 60, true),
		tipo("Reemplazo de la batería", 30, false),
	}
}

func TestApuntar(t *testing.T) {
	catalogo := catalogoDePrueba()
	citaMantenimiento := Cita{inicio: fecha(23, 9, 0), tipo: catalogo[0]}
	casos := []struct {
		nombre      string
		libreta     Libreta
		reparacion  string
		inicio      time.Time
		quiero      error
		quieroCitas int
	}{
		{"libreta vacía", Libreta{}, "Mantenimiento básico", fecha(23, 9, 0), nil, 1},
		{"dos citas a la vez", Libreta{citas: citasALaVez(2, fecha(23, 9, 0), 60)}, "Cambio de neumáticos y alineación", fecha(23, 9, 0), nil, 3},
		{"domingo", Libreta{}, "Mantenimiento básico", fecha(22, 10, 0), ErrDomingo, 0},
		{"fuera de horario", Libreta{}, "Mantenimiento básico", fecha(23, 7, 59), ErrFueraDeHorario, 0},
		{"tipo desconocido", Libreta{}, "Lavado de coche", fecha(23, 9, 0), ErrTipoDesconocido, 0},
		{"tipo desconocido en domingo", Libreta{}, "Lavado de coche", fecha(22, 10, 0), ErrTipoDesconocido, 0},
		{"cita duplicada", Libreta{citas: []Cita{citaMantenimiento}}, "Mantenimiento básico", fecha(23, 9, 0), ErrCitaDuplicada, 1},
		{"duplicada con los tres elevadores ocupados", Libreta{citas: append(citasALaVez(2, fecha(23, 9, 0), 60), citaMantenimiento)}, "Mantenimiento básico", fecha(23, 9, 0), ErrCitaDuplicada, 3},
		{"tres citas a la vez", Libreta{citas: citasALaVez(3, fecha(23, 9, 0), 60)}, "Cambio de neumáticos y alineación", fecha(23, 9, 0), ErrSinElevadorLibre, 3},
		{"tres citas a la vez y la nueva no usa elevador", Libreta{citas: citasALaVez(3, fecha(23, 9, 0), 60)}, "Reemplazo de la batería", fecha(23, 9, 0), nil, 4},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			antes := slices.Clone(c.libreta.citas)
			obtenida, err := c.libreta.Apuntar(catalogo, c.reparacion, c.inicio)
			if !errors.Is(err, c.quiero) {
				t.Errorf("error obtenido %v, quiero %v", err, c.quiero)
			}
			if len(obtenida.citas) != c.quieroCitas {
				t.Errorf("citas obtenidas %d, quiero %d", len(obtenida.citas), c.quieroCitas)
			}
			if !reflect.DeepEqual(c.libreta.citas, antes) {
				t.Error("la libreta original ha cambiado")
			}
		})
	}
}

func TestApuntarAnadeLaCita(t *testing.T) {
	catalogo := catalogoDePrueba()
	obtenida, err := Libreta{}.Apuntar(catalogo, "Mantenimiento básico", fecha(23, 9, 0))
	if err != nil {
		t.Fatal(err)
	}
	quiero := []Cita{{inicio: fecha(23, 9, 0), tipo: catalogo[0]}}
	if !reflect.DeepEqual(obtenida.citas, quiero) {
		t.Errorf("obtenidas %v, quiero %v", obtenida.citas, quiero)
	}
}

func TestApuntarNuncaSuperaLosElevadores(t *testing.T) {
	catalogo := []TipoReparacion{
		tipo("Corta", 45, true),
		tipo("Media", 60, true),
		tipo("Larga", 90, true),
	}
	libreta := Libreta{}
	rechazos := 0
	for _, minuto := range []int{0, 30} {
		for hora := 8; hora < 21; hora++ {
			for _, reparacion := range catalogo {
				var err error
				libreta, err = libreta.Apuntar(catalogo, reparacion.nombreEnCatalogo, fecha(23, hora, minuto))
				if errors.Is(err, ErrSinElevadorLibre) {
					rechazos++
				}
				ocupados := libreta.elevadoresOcupados(fecha(23, 8, 0), fecha(25, 21, 0))
				if ocupados > elevadoresTaller {
					t.Fatalf("hay %d elevadores ocupados a la vez, el máximo es %d", ocupados, elevadoresTaller)
				}
			}
		}
	}
	if rechazos == 0 {
		t.Error("la secuencia no ha rechazado ninguna cita por falta de elevador")
	}
}
