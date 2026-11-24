package elevadores

import (
	"errors"
	"time"

	"github.com/LuisPV05/LuisMecanico/internal/taller/cita"
)

const Cantidad = 3

var ErrSinElevadorLibre = errors.New("no hay ningún elevador libre")

func ocupadosEn(citas []cita.Cita, momento time.Time) int {
	ocupados := 0
	for _, c := range citas {
		if c.OcupaElevadorEn(momento) {
			ocupados++
		}
	}
	return ocupados
}

func iniciosDentroDe(citas []cita.Cita, inicio, fin time.Time) []time.Time {
	instantes := []time.Time{inicio}
	for _, c := range citas {
		if c.Inicio().After(inicio) && c.Inicio().Before(fin) {
			instantes = append(instantes, c.Inicio())
		}
	}
	return instantes
}

func Ocupados(citas []cita.Cita, inicio, fin time.Time) int {
	maximo := 0
	for _, instante := range iniciosDentroDe(citas, inicio, fin) {
		maximo = max(maximo, ocupadosEn(citas, instante))
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
