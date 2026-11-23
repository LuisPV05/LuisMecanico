package taller

import (
	"errors"
	"slices"
	"time"
)

var ErrCitaDuplicada = errors.New("ya existe una cita con el mismo inicio y el mismo tipo")

func (c Cita) esDuplicadaDe(otra Cita) bool {
	return c.inicio.Equal(otra.inicio) && c.tipo.nombreEnCatalogo == otra.tipo.nombreEnCatalogo
}

func (l Libreta) comprobarDuplicada(nueva Cita) error {
	esDuplicada := func(existente Cita) bool { return nueva.esDuplicadaDe(existente) }
	if slices.ContainsFunc(l.citas, esDuplicada) {
		return ErrCitaDuplicada
	}
	return nil
}

func (l Libreta) conCita(nueva Cita) Libreta {
	return Libreta{citas: append(slices.Clone(l.citas), nueva)}
}

func (l Libreta) reglasDeApuntado() []func(Cita) error {
	return []func(Cita) error{
		func(nueva Cita) error { return validarInicio(nueva.inicio) },
		l.comprobarDuplicada,
		l.comprobarElevador,
	}
}

func (l Libreta) comprobarReglas(nueva Cita) error {
	for _, regla := range l.reglasDeApuntado() {
		if err := regla(nueva); err != nil {
			return err
		}
	}
	return nil
}

func (l Libreta) Apuntar(catalogo []TipoReparacion, reparacion string, inicio time.Time) (Libreta, error) {
	nueva, err := nuevaCita(catalogo, reparacion, inicio)
	if err != nil {
		return l, err
	}
	if err := l.comprobarReglas(nueva); err != nil {
		return l, err
	}
	return l.conCita(nueva), nil
}