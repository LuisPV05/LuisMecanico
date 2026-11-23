package taller

import "errors"

var ErrSinElevadorLibre = errors.New("no hay ningún elevador libre")

func (l Libreta) comprobarElevador(nueva Cita) error {
	if !nueva.tipo.usaElevador {
		return nil
	}
	if l.elevadoresOcupados(nueva.inicio, nueva.fin()) >= elevadoresTaller {
		return ErrSinElevadorLibre
	}
	return nil
}
