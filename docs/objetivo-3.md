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

### Mediciones del gestor de dependencias

Se mide cuánto cuesta dejar el proyecto listo en una máquina limpia, porque esta herramienta se usará en entornos donde la velocidad de instalación importa (objetivos 5 y 6). Máquina: Windows, Git Bash, Go 1.27.3 (`go version go1.27.3 windows/amd64`).

| Herramienta | Instalación aparte de Go | `go mod download` o equivalente | Compilar el proyecto | Captura |
|---|---|---|---|---|
| Módulos de Go | Ninguna | 0,152 s | 0,491 s | [captura](../media/dep-go-modules.png) |

## Gestor de tareas

El comando `go` ejecuta subcomandos sueltos (`go build`, `go test`, `go mod download`), pero ninguno lee un fichero con tareas con nombre. `go generate` ejecuta comandos escritos en comentarios de los fuentes con la intención de crear o actualizar código Go. Los módulos de Go gestionan dependencias, no tareas, así que no se usan también como gestor de tareas: hace falta una herramienta aparte.

**Opciones:**

| Opción | Fichero de tareas | Instalación | Última versión |
|---|---|---|---|
| make | `Makefile` | Del sistema operativo, ajena a Go | 4.4.1, 26-02-2023 ([NEWS](https://fossies.org/linux/make/NEWS)) |
| just | `justfile` | Binario externo escrito en Rust | 1.59.0, 15-10-2026 ([crates.io](https://crates.io/crates/just)) |
| Task | `Taskfile.yml` | `go install`, con el toolchain de Go que el proyecto ya necesita | v3.54.0, 01-10-2026 ([GitHub](https://github.com/go-task/task/releases/tag/v3.54.0)) |
| Mage | `magefile.go`, con las tareas escritas en Go | `go install` | v1.17.2, 23-04-2026 ([GitHub](https://github.com/magefile/mage/releases)) |
| Scripts de shell | Scripts propios | Ninguna | No aplica: no hay un fichero de tareas común |
| Bazel | Ficheros `BUILD` | Sí | No se evalúa en detalle: es un sistema de construcción general |

**Criterios:**

1. Instalación: si se instala con el toolchain de Go que el proyecto ya exige o hay que instalar algo ajeno.
2. Mantenimiento: fecha de la última versión.
3. Si tiene un fichero de tareas con nombre, versionado en el repositorio, al que pueda apuntar `iv.yaml`.
4. Cuánto hay que mantener: si las tareas son configuración o código del proyecto.

**Decisión: Task.** Task y Mage se instalan con el toolchain de Go; make y just obligan a instalar algo ajeno a Go (el sistema operativo o un binario de Rust). Los dos están mantenidos, con versiones de octubre y abril de 2026. Mage escribe las tareas como código Go, que habría que mantener y compilar para tres tareas de una línea. Task las declara en `Taskfile.yml`. El coste asumido es una herramienta externa que se instala una vez con `go install`.

## Tarea `check`

El enunciado pide que `check` compruebe la sintaxis de los fuentes existentes y solo eso.

**Opciones**, probadas con Go 1.24.7 sobre un paquete con un error de sintaxis y otro con un error de tipos (una variable no definida):

| Comando | Error de sintaxis | Error de tipos | Qué más hace |
|---|---|---|---|
| `gofmt -e .` | Lo detecta | No lo detecta | Imprime el código formateado |
| `go build ./...` | Lo detecta | Lo detecta | Compila |
| `go vet ./...` | Lo detecta | Lo detecta | Análisis estático de construcciones sospechosas |
| `go test ./...` | Lo detecta | Lo detecta | Ejecuta los tests |

**Criterio:** que compruebe la sintaxis y nada más.

**Decisión: `gofmt -e . > /dev/null`.** Es el único comando que detecta los errores de sintaxis sin detectar los de tipos ni hacer nada más. Los demás van más allá de la sintaxis. Se descarta la salida estándar porque, sin eso, imprime todo el código formateado, y los errores salen por la salida de error. El coste asumido es que no comprueba que el código compile; eso lo hace `task test`.

## Dónde está cada cosa

- `go.mod`: módulo de Go.
- `Taskfile.yml`: tareas `install`, `test` y `check`.
- `iv.yaml`: clave `automatizar` con `fichero: Taskfile.yml` y `orden: task`.
- `README.md`: cómo instalar Task y ejecutar las tareas, y descripción de la entidad.