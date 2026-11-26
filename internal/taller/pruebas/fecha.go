package pruebas

import "time"

func Fecha(dia, hora, minuto int) time.Time {
	return time.Date(2026, time.November, dia, hora, minuto, 0, 0, time.UTC)
}