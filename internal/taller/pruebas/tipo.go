package pruebas

import (
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/catalogo"
)

func Tipo(nombre string, minutos int, usaElevador bool) catalogo.TipoReparacion {
	return catalogo.NuevoTipo(nombre, time.Duration(minutos)*time.Minute, usaElevador)
}
