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

func convertirNombre(texto string) (string, error) {
	nombre := strings.TrimSpace(texto)
	if nombre == "" {
		return "", ErrNombreVacio
	}
	return nombre, nil
}

func convertirDuracion(texto string) (time.Duration, error) {
	minutos, err := strconv.Atoi(strings.TrimSpace(texto))
	if err != nil {
		return 0, ErrMinutosNoNumericos
	}
	if minutos <= 0 {
		return 0, ErrDuracionInvalida
	}
	return time.Duration(minutos) * time.Minute, nil
}

func convertirUsaElevador(texto string) (bool, error) {
	switch strings.TrimSpace(texto) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, ErrUsaElevadorInvalido
}

func convertirFila(fila []string) (TipoReparacion, error) {
	if len(fila) != columnasPorFila {
		return TipoReparacion{}, ErrFilaIncompleta
	}
	nombre, err := convertirNombre(fila[0])
	if err != nil {
		return TipoReparacion{}, err
	}
	duracion, err := convertirDuracion(fila[1])
	if err != nil {
		return TipoReparacion{}, err
	}
	usaElevador, err := convertirUsaElevador(fila[2])
	if err != nil {
		return TipoReparacion{}, err
	}
	return NuevoTipo(nombre, duracion, usaElevador), nil
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
