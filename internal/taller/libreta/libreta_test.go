package libreta_test

import (
	"errors"
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/catalogo"
	"github.com/LuisPV05/LuisMecanico/internal/taller/cita"
	"github.com/LuisPV05/LuisMecanico/internal/taller/elevadores"
	"github.com/LuisPV05/LuisMecanico/internal/taller/horario"
	"github.com/LuisPV05/LuisMecanico/internal/taller/libreta"
	"github.com/LuisPV05/LuisMecanico/internal/taller/pruebas"
)

func catalogoDePrueba() []catalogo.TipoReparacion {
	return []catalogo.TipoReparacion{
		pruebas.Tipo("Mantenimiento básico", 45, true),
		pruebas.Tipo("Cambio de neumáticos y alineación", 60, true),
		pruebas.Tipo("Reemplazo de la batería", 30, false),
	}
}

func TestApuntar(t *testing.T) {
	tipos := catalogoDePrueba()
	mantenimiento := pruebas.CitaDe(pruebas.Fecha(23, 9, 0), 45)
	tresALaVez := pruebas.CitasALaVez(3, pruebas.Fecha(23, 9, 0), 60)
	casos := []struct {
		nombre      string
		previa      libreta.Libreta
		reparacion  string
		inicio      time.Time
		quiero      error
		quieroCitas int
	}{
		{"libreta vacía", libreta.Libreta{}, "Mantenimiento básico", pruebas.Fecha(23, 9, 0), nil, 1},
		{"dos citas a la vez", libreta.Desde(pruebas.CitasALaVez(2, pruebas.Fecha(23, 9, 0), 60)), "Cambio de neumáticos y alineación", pruebas.Fecha(23, 9, 0), nil, 3},
		{"domingo", libreta.Libreta{}, "Mantenimiento básico", pruebas.Fecha(22, 10, 0), horario.ErrDomingo, 0},
		{"fuera de horario", libreta.Libreta{}, "Mantenimiento básico", pruebas.Fecha(23, 7, 59), horario.ErrFueraDeHorario, 0},
		{"tipo desconocido", libreta.Libreta{}, "Lavado de coche", pruebas.Fecha(23, 9, 0), catalogo.ErrTipoDesconocido, 0},
		{"tipo desconocido en domingo", libreta.Libreta{}, "Lavado de coche", pruebas.Fecha(22, 10, 0), catalogo.ErrTipoDesconocido, 0},
		{"cita duplicada", libreta.Desde([]cita.Cita{mantenimiento}), "Mantenimiento básico", pruebas.Fecha(23, 9, 0), libreta.ErrCitaDuplicada, 1},
		{"duplicada con los tres elevadores ocupados", libreta.Desde(append(pruebas.CitasALaVez(2, pruebas.Fecha(23, 9, 0), 60), mantenimiento)), "Mantenimiento básico", pruebas.Fecha(23, 9, 0), libreta.ErrCitaDuplicada, 3},
		{"tres citas a la vez", libreta.Desde(tresALaVez), "Cambio de neumáticos y alineación", pruebas.Fecha(23, 9, 0), elevadores.ErrSinElevadorLibre, 3},
		{"tres citas a la vez y la nueva no usa elevador", libreta.Desde(tresALaVez), "Reemplazo de la batería", pruebas.Fecha(23, 9, 0), nil, 4},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			antes := slices.Clone(c.previa.Citas())
			obtenida, err := c.previa.Apuntar(tipos, c.reparacion, c.inicio)
			if !errors.Is(err, c.quiero) {
				t.Errorf("error obtenido %v, quiero %v", err, c.quiero)
			}
			if len(obtenida.Citas()) != c.quieroCitas {
				t.Errorf("citas obtenidas %d, quiero %d", len(obtenida.Citas()), c.quieroCitas)
			}
			if !reflect.DeepEqual(c.previa.Citas(), antes) {
				t.Error("la libreta original ha cambiado")
			}
		})
	}
}

func TestApuntarAnadeLaCita(t *testing.T) {
	obtenida, err := libreta.Libreta{}.Apuntar(catalogoDePrueba(), "Mantenimiento básico", pruebas.Fecha(23, 9, 0))
	if err != nil {
		t.Fatal(err)
	}
	quiero := []cita.Cita{pruebas.CitaDe(pruebas.Fecha(23, 9, 0), 45)}
	if !reflect.DeepEqual(obtenida.Citas(), quiero) {
		t.Errorf("obtenidas %v, quiero %v", obtenida.Citas(), quiero)
	}
}

func TestApuntarNuncaSuperaLosElevadores(t *testing.T) {
	tipos := []catalogo.TipoReparacion{
		pruebas.Tipo("Corta", 45, true),
		pruebas.Tipo("Media", 60, true),
		pruebas.Tipo("Larga", 90, true),
	}
	agenda := libreta.Libreta{}
	rechazos := 0
	for _, minuto := range []int{0, 30} {
		for hora := 8; hora < 21; hora++ {
			for _, tipo := range tipos {
				var err error
				agenda, err = agenda.Apuntar(tipos, tipo.Nombre(), pruebas.Fecha(23, hora, minuto))
				if errors.Is(err, elevadores.ErrSinElevadorLibre) {
					rechazos++
				}
				ocupados := elevadores.Ocupados(agenda.Citas(), pruebas.Fecha(23, 8, 0), pruebas.Fecha(25, 21, 0))
				if ocupados > elevadores.Cantidad {
					t.Fatalf("hay %d elevadores ocupados a la vez, el máximo es %d", ocupados, elevadores.Cantidad)
				}
			}
		}
	}
	if rechazos == 0 {
		t.Error("la secuencia no ha rechazado ninguna cita por falta de elevador")
	}
}
