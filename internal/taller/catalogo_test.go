package taller

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"time"
)

func tipo(nombre string, minutos int, usaElevador bool) TipoReparacion {
	return TipoReparacion{
		nombreEnCatalogo: nombre,
		duracion:         time.Duration(minutos) * time.Minute,
		usaElevador:      usaElevador,
	}
}

func TestConvertirFilasValidas(t *testing.T) {
	filas := [][]string{
		{"Mantenimiento básico", "45", "true"},
		{"Reemplazo de la batería principal y comprobación del alternador", "30", "false"},
	}
	quiero := []TipoReparacion{
		tipo("Mantenimiento básico", 45, true),
		tipo("Reemplazo de la batería principal y comprobación del alternador", 30, false),
	}
	obtenido, err := convertirFilas(filas)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(obtenido, quiero) {
		t.Errorf("obtenido %v, quiero %v", obtenido, quiero)
	}
}

func TestConvertirFilasInvalidas(t *testing.T) {
	casos := []struct {
		nombre string
		filas  [][]string
		quiero error
	}{
		{"minutos no numéricos", [][]string{{"Mantenimiento básico", "abc", "true"}}, ErrMinutosNoNumericos},
		{"minutos cero", [][]string{{"Mantenimiento básico", "0", "true"}}, ErrDuracionInvalida},
		{"minutos negativos", [][]string{{"Mantenimiento básico", "-5", "true"}}, ErrDuracionInvalida},
		{"elevador inválido", [][]string{{"Mantenimiento básico", "45", "quizá"}}, ErrUsaElevadorInvalido},
		{"nombre vacío", [][]string{{"", "45", "true"}}, ErrNombreVacio},
		{"nombre repetido", [][]string{{"Mantenimiento básico", "45", "true"}, {"Mantenimiento básico", "60", "false"}}, ErrNombreRepetido},
		{"fila incompleta", [][]string{{"Mantenimiento básico", "45"}}, ErrFilaIncompleta},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			_, err := convertirFilas(c.filas)
			if !errors.Is(err, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", err, c.quiero)
			}
		})
	}
}

func TestBuscarTipo(t *testing.T) {
	tipos := []TipoReparacion{
		tipo("Mantenimiento básico", 45, true),
		tipo("Sustitución de amortiguadores", 150, true),
	}
	casos := []struct {
		nombre    string
		quiero    TipoReparacion
		quieroErr error
	}{
		{"Sustitución de amortiguadores", tipo("Sustitución de amortiguadores", 150, true), nil},
		{"Lavado de coche", TipoReparacion{}, ErrTipoDesconocido},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			obtenido, err := buscarTipo(tipos, c.nombre)
			if !errors.Is(err, c.quieroErr) {
				t.Errorf("error obtenido %v, quiero %v", err, c.quieroErr)
			}
			if !reflect.DeepEqual(obtenido, c.quiero) {
				t.Errorf("obtenido %v, quiero %v", obtenido, c.quiero)
			}
		})
	}
}

func TestCargarCatalogo(t *testing.T) {
	quiero := []TipoReparacion{
		tipo("Mantenimiento básico", 45, true),
		tipo("Cambio de pastillas y discos de freno", 90, true),
		tipo("Cambio de neumáticos y alineación", 60, true),
		tipo("Cambio de kit de distribución (Bomba de agua y correas)", 240, true),
		tipo("Reparación del sistema de embrague", 300, true),
		tipo("Sustitución del alternador o motor de arranque", 120, true),
		tipo("Diagnóstico y reparación de falla eléctrica (Cortocircuito o sensor defectuoso)", 180, true),
		tipo("Reemplazo de la batería principal y comprobación del alternador", 30, false),
		tipo("Sustitución de amortiguadores (Suspensión delantera o trasera)", 150, true),
		tipo("Sustitución completa del motor", 2340, true),
	}
	obtenido, err := cargarCatalogo("tipos_reparacion.csv")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(obtenido, quiero) {
		t.Errorf("obtenido %v, quiero %v", obtenido, quiero)
	}
}

func TestCargarCatalogoFicheroInexistente(t *testing.T) {
	_, err := cargarCatalogo("no_existe.csv")
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("obtenido %v, quiero %v", err, fs.ErrNotExist)
	}
}
