package taller

import "time"

type TipoReparacion struct {
	nombre      string
	duracion    time.Duration
	usaElevador bool
}

var tiposReparacion = map[string]TipoReparacion{
	"Mantenimiento básico":                                                            {nombre: "Mantenimiento básico", duracion: 45 * time.Minute, usaElevador: true},
	"Cambio de pastillas y discos de freno":                                           {nombre: "Cambio de pastillas y discos de freno", duracion: 90 * time.Minute, usaElevador: true},
	"Cambio de neumáticos y alineación":                                               {nombre: "Cambio de neumáticos y alineación", duracion: 60 * time.Minute, usaElevador: true},
	"Cambio de kit de distribución (Bomba de agua y correas)":                         {nombre: "Cambio de kit de distribución (Bomba de agua y correas)", duracion: 240 * time.Minute, usaElevador: true},
	"Reparación del sistema de embrague":                                              {nombre: "Reparación del sistema de embrague", duracion: 300 * time.Minute, usaElevador: true},
	"Sustitución del alternador o motor de arranque":                                  {nombre: "Sustitución del alternador o motor de arranque", duracion: 120 * time.Minute, usaElevador: true},
	"Diagnóstico y reparación de falla eléctrica (Cortocircuito o sensor defectuoso)": {nombre: "Diagnóstico y reparación de falla eléctrica (Cortocircuito o sensor defectuoso)", duracion: 180 * time.Minute, usaElevador: true},
	"Reemplazo de la batería principal y comprobación del alternador":                 {nombre: "Reemplazo de la batería principal y comprobación del alternador", duracion: 30 * time.Minute, usaElevador: false},
	"Sustitución de amortiguadores (Suspensión delantera o trasera)":                  {nombre: "Sustitución de amortiguadores (Suspensión delantera o trasera)", duracion: 150 * time.Minute, usaElevador: true},
	"Sustitución completa del motor":                                                  {nombre: "Sustitución completa del motor", duracion: 3 * 13 * time.Hour, usaElevador: true},
}
