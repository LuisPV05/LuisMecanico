package taller

import (
	_ "embed"
	"encoding/csv"
	"strconv"
	"strings"
	"time"
)

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}

//go:embed tipos_reparacion.csv
var contenidoCSV string

var tiposReparacion = leerTiposReparacion()

func leerTiposReparacion() map[string]TipoReparacion {
	filas, err := csv.NewReader(strings.NewReader(contenidoCSV)).ReadAll()
	if err != nil {
		panic(err)
	}
	tipos := map[string]TipoReparacion{}
	for _, campos := range filas[1:] {
		minutos, err := strconv.Atoi(campos[1])
		if err != nil {
			panic(err)
		}
		usaElevador, err := strconv.ParseBool(campos[2])
		if err != nil {
			panic(err)
		}
		tipos[campos[0]] = TipoReparacion{
			nombre:      campos[0],
			duracion:    time.Duration(minutos) * time.Minute,
			usaElevador: usaElevador,
		}
	}
	return tipos
}
