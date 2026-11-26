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
	tipos []catalogo.TipoReparacion
	citas []cita.Cita
}

func Nueva(tipos []catalogo.TipoReparacion) *Libreta {
	return &Libreta{tipos: slices.Clone(tipos)}
}

func Abrir(rutaCatalogo string) (*Libreta, error) {
	tipos, err := catalogo.Cargar(rutaCatalogo)
	if err != nil {
		return nil, err
	}
	return Nueva(tipos), nil
}

func (l *Libreta) Citas() []cita.Cita {
	return slices.Clone(l.citas)
}

func (l *Libreta) contieneLaMisma(nueva cita.Cita) bool {
	return slices.ContainsFunc(l.citas, func(c cita.Cita) bool {
		return c.Inicio().Equal(nueva.Inicio()) && c.Tipo() == nueva.Tipo()
	})
}

func (l *Libreta) validar(nueva cita.Cita) error {
	if err := horario.ValidarInicio(nueva.Inicio()); err != nil {
		return err
	}
	if l.contieneLaMisma(nueva) {
		return ErrCitaDuplicada
	}
	return elevadores.Comprobar(l.citas, nueva)
}

func (l *Libreta) Apuntar(nombreTipo string, inicio time.Time) error {
	nueva, err := cita.Nueva(l.tipos, nombreTipo, inicio)
	if err != nil {
		return err
	}
	if err := l.validar(nueva); err != nil {
		return err
	}
	l.citas = append(l.citas, nueva)
	return nil
}
