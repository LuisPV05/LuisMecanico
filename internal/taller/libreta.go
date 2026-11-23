package taller

import "time"

const (
	horaApertura     = 8
	horaCierre       = 21
	elevadoresTaller = 3
)

type Libreta struct {
	citas []Cita
}

func (l Libreta) ocupadosEn(momento time.Time) int {
	ocupados := 0
	for _, c := range l.citas {
		if c.ocupaElevadorEn(momento) {
			ocupados++
		}
	}
	return ocupados
}

func (l Libreta) iniciosDentroDe(inicio, fin time.Time) []time.Time {
	instantes := []time.Time{inicio}
	for _, c := range l.citas {
		if c.inicio.After(inicio) && c.inicio.Before(fin) {
			instantes = append(instantes, c.inicio)
		}
	}
	return instantes
}

func (l Libreta) elevadoresOcupados(inicio, fin time.Time) int {
	maximo := 0
	for _, instante := range l.iniciosDentroDe(inicio, fin) {
		maximo = max(maximo, l.ocupadosEn(instante))
	}
	return maximo
}