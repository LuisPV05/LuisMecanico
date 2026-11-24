package cita

import (
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/catalogo"
	"github.com/LuisPV05/LuisMecanico/internal/taller/horario"
)

type Cita struct {
	inicio time.Time
	tipo   catalogo.TipoReparacion
}

func Nueva(tipos []catalogo.TipoReparacion, nombreTipo string, inicio time.Time) (Cita, error) {
	encontrado, err := catalogo.Buscar(tipos, nombreTipo)
	if err != nil {
		return Cita{}, err
	}
	if encontrado.Duracion() <= 0 {
		return Cita{}, catalogo.ErrDuracionInvalida
	}
	return Cita{inicio: inicio, tipo: encontrado}, nil
}

func (c Cita) Inicio() time.Time {
	return c.inicio
}

func (c Cita) Tipo() catalogo.TipoReparacion {
	return c.tipo
}

func (c Cita) Fin() time.Time {
	return horario.FinDeCita(c.inicio, c.tipo.Duracion())
}

func Solapan(a, b Cita) bool {
	return a.inicio.Before(b.Fin()) && b.inicio.Before(a.Fin())
}

func (c Cita) OcupaElevadorEn(momento time.Time) bool {
	return c.tipo.UsaElevador() && !momento.Before(c.inicio) && momento.Before(c.Fin())
}
