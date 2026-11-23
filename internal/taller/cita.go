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

func (c Cita) fin() time.Time {
	return finDeCita(c.inicio, c.tipo.duracion)
}

func (c Cita) solapaConTramo(inicio, fin time.Time) bool {
	return c.inicio.Before(fin) && inicio.Before(c.fin())
}

func solapan(a, b Cita) bool {
	return a.solapaConTramo(b.inicio, b.fin())
}

func (c Cita) ocupaElevadorEn(momento time.Time) bool {
	return c.tipo.usaElevador && !momento.Before(c.inicio) && momento.Before(c.fin())
}