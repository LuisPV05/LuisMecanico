package catalogo

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

const columnasPorFila = 3

var (
	ErrTipoDesconocido     = errors.New("tipo de reparación desconocido")
	ErrMinutosNoNumericos  = errors.New("los minutos no son un número")
	ErrDuracionInvalida    = errors.New("la duración debe ser mayor que cero")
	ErrUsaElevadorInvalido = errors.New("usa_elevador debe ser true o false")
	ErrNombreVacio         = errors.New("el nombre del tipo está vacío")
	ErrNombreRepetido      = errors.New("el nombre del tipo está repetido")
	ErrFilaIncompleta      = errors.New("la fila no tiene tres columnas")
)

func leerFilas(ruta string) ([][]string, error) {
	fichero, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer fichero.Close()
	lector := csv.NewReader(fichero)
	lector.FieldsPerRecord = -1
	if _, err := lector.Read(); err != nil {
		return nil, err
	}
	return lector.ReadAll()
}

func convertirFila(fila []string) (TipoReparacion, error) {
	if len(fila) != columnasPorFila {
		return TipoReparacion{}, ErrFilaIncompleta
	}
	nombre := strings.TrimSpace(fila[0])
	if nombre == "" {
		return TipoReparacion{}, ErrNombreVacio
	}
	minutos, err := strconv.Atoi(strings.TrimSpace(fila[1]))
	if err != nil {
		return TipoReparacion{}, ErrMinutosNoNumericos
	}
	if minutos <= 0 {
		return TipoReparacion{}, ErrDuracionInvalida
	}
	var usaElevador bool
	switch strings.TrimSpace(fila[2]) {
	case "true":
		usaElevador = true
	case "false":
		usaElevador = false
	default:
		return TipoReparacion{}, ErrUsaElevadorInvalido
	}
	return NuevoTipo(nombre, time.Duration(minutos)*time.Minute, usaElevador), nil
}

func DesdeFilas(filas [][]string) ([]TipoReparacion, error) {
	tipos := make([]TipoReparacion, 0, len(filas))
	vistos := make(map[string]struct{}, len(filas))
	for _, fila := range filas {
		convertido, err := convertirFila(fila)
		if err != nil {
			return nil, err
		}
		if _, repetido := vistos[convertido.Nombre()]; repetido {
			return nil, ErrNombreRepetido
		}
		vistos[convertido.Nombre()] = struct{}{}
		tipos = append(tipos, convertido)
	}
	return tipos, nil
}

func Buscar(tipos []TipoReparacion, nombre string) (TipoReparacion, error) {
	for _, candidato := range tipos {
		if candidato.Nombre() == nombre {
			return candidato, nil
		}
	}
	return TipoReparacion{}, ErrTipoDesconocido
}

func Cargar(ruta string) ([]TipoReparacion, error) {
	filas, err := leerFilas(ruta)
	if err != nil {
		return nil, err
	}
	return DesdeFilas(filas)
}
