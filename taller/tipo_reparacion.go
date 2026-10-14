package taller

import "time"

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}

func TiposReparacion() []TipoReparacion {
	return []TipoReparacion{
		{nombre: "Mantenimiento básico", duracion: 45 * time.Minute, usaElevador: true},
		{nombre: "Cambio de pastillas y discos de freno", duracion: 90 * time.Minute, usaElevador: true},
		{nombre: "Cambio de neumáticos y alineación", duracion: 60 * time.Minute, usaElevador: true},
		{nombre: "Cambio de kit de distribución (Bomba de agua y correas)", duracion: 240 * time.Minute, usaElevador: true},
		{nombre: "Reparación del sistema de embrague", duracion: 300 * time.Minute, usaElevador: true},
		{nombre: "Sustitución del alternador o motor de arranque", duracion: 120 * time.Minute, usaElevador: true},
		{nombre: "Diagnóstico y reparación de falla eléctrica (Cortocircuito o sensor defectuoso)", duracion: 180 * time.Minute, usaElevador: true},
		{nombre: "Reemplazo de la batería principal y comprobación del alternador", duracion: 30 * time.Minute, usaElevador: false},
		{nombre: "Sustitución de amortiguadores (Suspensión delantera o trasera)", duracion: 150 * time.Minute, usaElevador: true},
		{nombre: "Sustitución completa del motor", duracion: 3 * 13 * time.Hour, usaElevador: true},
	}
}
