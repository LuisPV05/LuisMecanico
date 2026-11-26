package libreta_test

import (
	"errors"
	"io/fs"
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

const (
	mantenimiento = "Mantenimiento básico"
	neumaticos    = "Cambio de neumáticos y alineación"
	pastillas     = "Cambio de pastillas y discos de freno"
	alternador    = "Sustitución del alternador o motor de arranque"
	bateria       = "Reemplazo de la batería"
)

func catalogoDePrueba() []catalogo.TipoReparacion {
	return []catalogo.TipoReparacion{
		pruebas.Tipo(mantenimiento, 45, true),
		pruebas.Tipo(neumaticos, 60, true),
		pruebas.Tipo(pastillas, 90, true),
		pruebas.Tipo(alternador, 120, true),
		pruebas.Tipo(bateria, 30, false),
	}
}

func apuntadas(t *testing.T, tipos []catalogo.TipoReparacion, inicio time.Time, nombres ...string) libreta.Libreta {
	t.Helper()
	agenda := libreta.Nueva(tipos)
	for _, nombre := range nombres {
		var err error
		agenda, err = agenda.Apuntar(nombre, inicio)
		if err != nil {
			t.Fatal(err)
		}
	}
	return agenda
}

func TestApuntar(t *testing.T) {
	tipos := catalogoDePrueba()
	nueve := pruebas.Fecha(23, 9, 0)
	vacia := libreta.Nueva(tipos)
	tres := apuntadas(t, tipos, nueve, neumaticos, pastillas, mantenimiento)
	casos := []struct {
		nombre      string
		previa      libreta.Libreta
		reparacion  string
		inicio      time.Time
		quiero      error
		quieroCitas int
	}{
		{"libreta vacía", vacia, mantenimiento, nueve, nil, 1},
		{"dos citas a la vez", apuntadas(t, tipos, nueve, neumaticos, pastillas), mantenimiento, nueve, nil, 3},
		{"domingo", vacia, mantenimiento, pruebas.Fecha(22, 10, 0), horario.ErrDomingo, 0},
		{"fuera de horario", vacia, mantenimiento, pruebas.Fecha(23, 7, 59), horario.ErrFueraDeHorario, 0},
		{"tipo desconocido", vacia, "Lavado de coche", nueve, catalogo.ErrTipoDesconocido, 0},
		{"tipo desconocido en domingo", vacia, "Lavado de coche", pruebas.Fecha(22, 10, 0), catalogo.ErrTipoDesconocido, 0},
		{"cita duplicada", apuntadas(t, tipos, nueve, mantenimiento), mantenimiento, nueve, libreta.ErrCitaDuplicada, 1},
		{"duplicada con los tres elevadores ocupados", tres, mantenimiento, nueve, libreta.ErrCitaDuplicada, 3},
		{"tres citas a la vez", tres, alternador, nueve, elevadores.ErrSinElevadorLibre, 3},
		{"tres citas a la vez y la nueva no usa elevador", tres, bateria, nueve, nil, 4},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			antes := slices.Clone(c.previa.Citas())
			obtenida, err := c.previa.Apuntar(c.reparacion, c.inicio)
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
	obtenida, err := libreta.Nueva(catalogoDePrueba()).Apuntar(mantenimiento, pruebas.Fecha(23, 9, 0))
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
	agenda := libreta.Nueva(tipos)
	rechazos := 0
	for _, minuto := range []int{0, 30} {
		for hora := 8; hora < 21; hora++ {
			for _, tipo := range tipos {
				var err error
				agenda, err = agenda.Apuntar(tipo.Nombre(), pruebas.Fecha(23, hora, minuto))
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

func TestAbrirCargaElCatalogoDelTaller(t *testing.T) {
	agenda, err := libreta.Abrir("../catalogo/tipos_reparacion.csv")
	if err != nil {
		t.Fatal(err)
	}
	obtenida, err := agenda.Apuntar(mantenimiento, pruebas.Fecha(23, 9, 0))
	if err != nil {
		t.Fatal(err)
	}
	if len(obtenida.Citas()) != 1 {
		t.Errorf("citas obtenidas %d, quiero 1", len(obtenida.Citas()))
	}
}

func TestAbrirFicheroInexistente(t *testing.T) {
	_, err := libreta.Abrir("no_existe.csv")
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("obtenido %v, quiero %v", err, fs.ErrNotExist)
	}
}
