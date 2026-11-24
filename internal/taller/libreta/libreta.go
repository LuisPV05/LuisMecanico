package libreta

import (
	"errors"
	"slices"
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/catalogo"
	"github.com/LuisPV05/LuisMecanico/internal/taller/cita"
	"github.com/LuisPV05/LuisMecanico/internal/taller/elevadores"
	"github.com/LuisPV05/LuisMecanico/internal/taller/horario"
)

var ErrCitaDuplicada = errors.New("ya hay una cita igual apuntada")

type Libreta struct {
	citas []cita.Cita
}

func Desde(citas []cita.Cita) Libreta {
	return Libreta{citas: slices.Clone(citas)}
}

func (l Libreta) Citas() []cita.Cita {
	return slices.Clone(l.citas)
}

func (l Libreta) contieneLaMisma(nueva cita.Cita) bool {
	return slices.ContainsFunc(l.citas, func(c cita.Cita) bool {
		return c.Inicio().Equal(nueva.Inicio()) && c.Tipo() == nueva.Tipo()
	})
}

func (l Libreta) validar(nueva cita.Cita) error {
	if err := horario.ValidarInicio(nueva.Inicio()); err != nil {
		return err
	}
	if l.contieneLaMisma(nueva) {
		return ErrCitaDuplicada
	}
	return elevadores.Comprobar(l.citas, nueva)
}

func (l Libreta) Apuntar(tipos []catalogo.TipoReparacion, nombreTipo string, inicio time.Time) (Libreta, error) {
	nueva, err := cita.Nueva(tipos, nombreTipo, inicio)
	if err != nil {
		return l, err
	}
	if err := l.validar(nueva); err != nil {
		return l, err
	}
	return Libreta{citas: append(slices.Clone(l.citas), nueva)}, nil
}
