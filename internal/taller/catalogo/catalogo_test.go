package catalogo_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"

	"github.com/LuisPV05/LuisMecanico/internal/taller/catalogo"
	"github.com/LuisPV05/LuisMecanico/internal/taller/pruebas"
)

func TestDesdeFilasValidas(t *testing.T) {
	filas := [][]string{
		{"Mantenimiento básico", "45", "true"},
		{"Reemplazo de la batería principal y comprobación del alternador", "30", "false"},
	}
	quiero := []catalogo.TipoReparacion{
		pruebas.Tipo("Mantenimiento básico", 45, true),
		pruebas.Tipo("Reemplazo de la batería principal y comprobación del alternador", 30, false),
	}
	obtenido, err := catalogo.DesdeFilas(filas)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(obtenido, quiero) {
		t.Errorf("obtenido %v, quiero %v", obtenido, quiero)
	}
}

func TestDesdeFilasInvalidas(t *testing.T) {
	casos := []struct {
		nombre string
		filas  [][]string
		quiero error
	}{
		{"minutos no numéricos", [][]string{{"Mantenimiento básico", "abc", "true"}}, catalogo.ErrMinutosNoNumericos},
		{"minutos cero", [][]string{{"Mantenimiento básico", "0", "true"}}, catalogo.ErrDuracionInvalida},
		{"minutos negativos", [][]string{{"Mantenimiento básico", "-5", "true"}}, catalogo.ErrDuracionInvalida},
		{"elevador inválido", [][]string{{"Mantenimiento básico", "45", "quizá"}}, catalogo.ErrUsaElevadorInvalido},
		{"nombre vacío", [][]string{{"", "45", "true"}}, catalogo.ErrNombreVacio},
		{"nombre repetido", [][]string{{"Mantenimiento básico", "45", "true"}, {"Mantenimiento básico", "60", "false"}}, catalogo.ErrNombreRepetido},
		{"fila incompleta", [][]string{{"Mantenimiento básico", "45"}}, catalogo.ErrFilaIncompleta},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			_, err := catalogo.DesdeFilas(c.filas)
			if !errors.Is(err, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", err, c.quiero)
			}
		})
	}
}

func TestBuscar(t *testing.T) {
	tipos := []catalogo.TipoReparacion{
		pruebas.Tipo("Mantenimiento básico", 45, true),
		pruebas.Tipo("Sustitución de amortiguadores", 150, true),
	}
	casos := []struct {
		nombre    string
		quiero    catalogo.TipoReparacion
		quieroErr error
	}{
		{"Sustitución de amortiguadores", pruebas.Tipo("Sustitución de amortiguadores", 150, true), nil},
		{"Lavado de coche", catalogo.TipoReparacion{}, catalogo.ErrTipoDesconocido},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenido, err := catalogo.Buscar(tipos, c.nombre)
			if !errors.Is(err, c.quieroErr) {
				t.Errorf("error obtenido %v, quiero %v", err, c.quieroErr)
			}
			if !reflect.DeepEqual(obtenido, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", obtenido, c.quiero)
			}
		})
	}
}

func TestCargar(t *testing.T) {
	quiero := []catalogo.TipoReparacion{
		pruebas.Tipo("Mantenimiento básico", 45, true),
		pruebas.Tipo("Cambio de pastillas y discos de freno", 90, true),
		pruebas.Tipo("Cambio de neumáticos y alineación", 60, true),
		pruebas.Tipo("Cambio de kit de distribución (Bomba de agua y correas)", 240, true),
		pruebas.Tipo("Reparación del sistema de embrague", 300, true),
		pruebas.Tipo("Sustitución del alternador o motor de arranque", 120, true),
		pruebas.Tipo("Diagnóstico y reparación de falla eléctrica (Cortocircuito o sensor defectuoso)", 180, true),
		pruebas.Tipo("Reemplazo de la batería principal y comprobación del alternador", 30, false),
		pruebas.Tipo("Sustitución de amortiguadores (Suspensión delantera o trasera)", 150, true),
		pruebas.Tipo("Sustitución completa del motor", 2340, true),
	}
	obtenido, err := catalogo.Cargar("tipos_reparacion.csv")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(obtenido, quiero) {
		t.Errorf("obtenido %v, quiero %v", obtenido, quiero)
	}
}

func TestCargarFicheroInexistente(t *testing.T) {
	_, err := catalogo.Cargar("no_existe.csv")
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("obtenido %v, quiero %v", err, fs.ErrNotExist)
	}
}
