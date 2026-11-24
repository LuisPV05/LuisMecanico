package pruebas

import (
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/catalogo"
	"github.com/LuisPV05/LuisMecanico/internal/taller/cita"
)

func CitaDe(inicio time.Time, minutos int) cita.Cita {
	return citaConTipo(Tipo("Mantenimiento básico", minutos, true), inicio)
}

func CitaSinElevador(inicio time.Time, minutos int) cita.Cita {
	return citaConTipo(Tipo("Reemplazo de la batería", minutos, false), inicio)
}

func citaConTipo(tipo catalogo.TipoReparacion, inicio time.Time) cita.Cita {
	creada, err := cita.Nueva([]catalogo.TipoReparacion{tipo}, tipo.Nombre(), inicio)
	if err != nil {
		panic(err)
	}
	return creada
}
