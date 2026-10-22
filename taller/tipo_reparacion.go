package taller

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}

var tiposReparacion = leerTiposReparacion()

func leerTiposReparacion() map[string]TipoReparacion {
	contenido, err := os.ReadFile("tipos_reparacion.csv")
	if err != nil {
		panic(err)
	}
	tipos := map[string]TipoReparacion{}
	lineas := strings.Split(strings.TrimSpace(string(contenido)), "\n")
	for _, linea := range lineas[1:] {
		campos := strings.Split(linea, ",")
		minutos, _ := strconv.Atoi(campos[1])
		tipos[campos[0]] = TipoReparacion{
			nombre:      campos[0],
			duracion:    time.Duration(minutos) * time.Minute,
			usaElevador: campos[2] == "true",
		}
	}
	return tipos
}
