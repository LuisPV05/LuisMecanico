package taller

import "time"

func cierreDelDia(momento time.Time) time.Time {
	return time.Date(momento.Year(), momento.Month(), momento.Day(), horaCierre, 0, 0, 0, momento.Location())
}

func siguienteApertura(momento time.Time) time.Time {
	dia := momento.AddDate(0, 0, 1)
	if dia.Weekday() == time.Sunday {
		dia = dia.AddDate(0, 0, 1)
	}
	return time.Date(dia.Year(), dia.Month(), dia.Day(), horaApertura, 0, 0, 0, momento.Location())
}

func finDeCita(inicio time.Time, duracion time.Duration) time.Time {
	disponible := max(cierreDelDia(inicio).Sub(inicio), 0)
	if duracion <= disponible {
		return inicio.Add(duracion)
	}
	return finDeCita(siguienteApertura(inicio), duracion-disponible)
}
