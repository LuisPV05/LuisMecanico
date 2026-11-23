package taller

import "time"

type Cita struct {
	inicio time.Time
	tipo   TipoReparacion
}

type TipoReparacion struct {
	nombreEnCatalogo string
	duracion         time.Duration
	usaElevador      bool
}

func nuevaCita(catalogo []TipoReparacion, nombreTipo string, inicio time.Time) (Cita, error) {
	encontrado, err := buscarTipo(catalogo, nombreTipo)
	if err != nil {
		return Cita{}, err
	}
	if encontrado.duracion <= 0 {
		return Cita{}, ErrDuracionInvalida
	}
	return Cita{inicio: inicio, tipo: encontrado}, nil
}