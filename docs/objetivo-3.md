# Objetivo 3: gestor de dependencias, gestor de tareas y tarea `check`

Este documento recoge cómo se eligieron las herramientas que automatizan las tareas del proyecto. Cada decisión parte de un problema planteado en un issue del Milestone 1 y se toma con criterios que se pueden comprobar.

## Punto de partida

- El proyecto está escrito en Go y solo importa `time`, de la biblioteca estándar: no tiene dependencias externas.
- No existe `go.mod`, así que `go build ./...` falla porque no hay módulo principal.
- `iv.yaml` necesita apuntar a un fichero de tareas que declare `install`, `test` y `check`.

## Gestor de dependencias

**Opciones:**

| Opción | Qué es | Instalación aparte del toolchain de Go | Estado |
|---|---|---|---|
| Módulos de Go (`go.mod`, `go.sum`) | Sistema integrado en el comando `go` | Ninguna | Soporte oficial desde Go 1.11 ([README de golang/dep](https://github.com/golang/dep)) |
| dep | Gestor anterior a los módulos | Sí | Obsoleto y archivado en 2020 a favor de los módulos ([README de golang/dep](https://github.com/golang/dep)) |
| glide | Gestor anterior a los módulos | Sí | Retirado, junto con dep y otros, del buildpack de Go de Heroku el 10-03-2026 ([Heroku](https://devcenter.heroku.com/changelog-items/3619)) |
| govendor | Gestor anterior a los módulos | Sí | Retirado igual que glide ([Heroku](https://devcenter.heroku.com/changelog-items/3619)) |
| Bazel y Nix | Sistemas de construcción generales, no específicos de Go | Sí | No se evalúan en detalle |

**Criterios:**

1. Si es el estándar del lenguaje (si lo hay, se sigue).
2. Estado de mantenimiento: que no esté archivado ni retirado.
3. Si requiere instalar algo aparte del toolchain de Go.
4. Número de dependencias del proyecto: hoy, ninguna externa.

**Decisión: módulos de Go.** Son el sistema oficial desde Go 1.11 y vienen en el propio comando `go` (`go help modules`: "Modules are how Go manages dependencies"), así que no hay nada que instalar. dep está archivado, glide y govendor están retirados de Heroku, que solo admite módulos. Bazel y Nix no resuelven ningún problema de un proyecto sin dependencias externas y añadirían una herramienta más.

Consecuencia: `go.mod` en la raíz, con el módulo `github.com/LuisPV05/LuisMecanico`, y la tarea `install` ejecuta `go mod download`, que hoy no descarga nada porque no hay dependencias externas.