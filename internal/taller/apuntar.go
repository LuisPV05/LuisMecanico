package taller

import (
	"errors"
	"slices"
	"time"
)

var ErrCitaDuplicada = errors.New("ya hay una cita igual apuntada")

func (l Libreta) contieneLaMisma(nueva Cita) bool {
	return slices.ContainsFunc(l.citas, func(c Cita) bool {
		return c.inicio.Equal(nueva.inicio) && c.tipo == nueva.tipo
	})
}

func (l Libreta) validar(nueva Cita) error {
	if err := validarInicio(nueva.inicio); err != nil {
		return err
	}
	if l.contieneLaMisma(nueva) {
		return ErrCitaDuplicada
	}
	return l.comprobarElevador(nueva)
}

func (l Libreta) Apuntar(catalogo []TipoReparacion, nombreTipo string, inicio time.Time) (Libreta, error) {
	nueva, err := nuevaCita(catalogo, nombreTipo, inicio)
	if err != nil {
		return l, err
	}
	if err := l.validar(nueva); err != nil {
		return l, err
	}
	return Libreta{citas: append(slices.Clone(l.citas), nueva)}, nil
}
