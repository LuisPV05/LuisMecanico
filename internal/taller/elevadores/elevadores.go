package elevadores

import (
	"errors"
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/cita"
)

const Cantidad = 3

var ErrSinElevadorLibre = errors.New("no hay ningún elevador libre")

func Ocupados(citas []cita.Cita, inicio, fin time.Time) int {
	instantes := []time.Time{inicio}
	for _, c := range citas {
		if c.Inicio().After(inicio) && c.Inicio().Before(fin) {
			instantes = append(instantes, c.Inicio())
		}
	}
	maximo := 0
	for _, instante := range instantes {
		ocupados := 0
		for _, c := range citas {
			if c.OcupaElevadorEn(instante) {
				ocupados++
			}
		}
		maximo = max(maximo, ocupados)
	}
	return maximo
}

func Comprobar(citas []cita.Cita, nueva cita.Cita) error {
	if !nueva.Tipo().UsaElevador() {
		return nil
	}
	if Ocupados(citas, nueva.Inicio(), nueva.Fin()) >= Cantidad {
		return ErrSinElevadorLibre
	}
	return nil
}
