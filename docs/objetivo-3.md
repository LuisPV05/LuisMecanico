# Objetivo 3: gestor de dependencias, gestor de tareas y tarea `check`

Este documento recoge cómo se eligieron las herramientas que automatizan las tareas del proyecto. Cada decisión parte de un problema planteado en un issue del Milestone 1 y se toma con criterios que se pueden comprobar.

## Punto de partida

- El proyecto está escrito en Go y solo importa `time`, de la biblioteca estándar: no tiene dependencias externas.
- No existe `go.mod`, así que `go build ./...` falla porque no hay módulo principal.
- `iv.yaml` necesita apuntar a un fichero de tareas que declare `install` y `check`.

## Gestor de dependencias

Go tiene un estándar propio para gestionar dependencias: los módulos (`go.mod` y `go.sum`), integrados en el comando `go` y oficiales desde Go 1.11 ([notas de la versión](https://go.dev/doc/go1.11#modules)). La guía del objetivo dice que, si hay un estándar, primero hay que conocerlo y después seguirlo, adaptando el resto de criterios. Por eso elijo los módulos directamente y no comparo otras herramientas.

**Criterio:** que sea el estándar del lenguaje. Es el único criterio que decide, y por eso no hay otros que comparar.

**Otros gestores de dependencias para Go:** existen dep, glide y govendor, pero son anteriores a los módulos y están obsoletos. dep está archivado y recomienda migrar a los módulos, y según el aviso de Heroku del 10-03-2026 su buildpack de Go ya no admite glide ni govendor ([Heroku](https://devcenter.heroku.com/changelog-items/3619)). No los he examinado ni probado, porque al haber un estándar vigente no son candidatos.

### Medición de los módulos de Go

Esta medición no sirve para comparar, solo para comprobar cuánto cuesta dejar el proyecto listo, porque se usará en los objetivos 5 y 6, donde la velocidad importa. Máquina: la de desarrollo, con Windows, Git Bash y Go 1.27.3 (`go version go1.27.3 windows/amd64`).

Con los módulos de Go no hay que instalar nada aparte de Go. `go mod download` tarda 0,152 s (no descarga nada, porque el proyecto no tiene dependencias externas) y compilar el proyecto con `go build ./...` tarda 0,491 s. Captura: [captura](../media/dep-go-modules.png).

**Decisión: módulos de Go.** Son el estándar del lenguaje, vienen en el propio comando `go` y no hay nada que instalar. Preparar el proyecto cuesta 0,152 s y compilarlo 0,491 s, y el proyecto no tiene dependencias externas, así que no hay ninguna razón para añadir otra herramienta. Tampoco se usan los módulos como gestor de tareas: gestionan dependencias, no tareas, y por eso hace falta una herramienta aparte, que se elige a continuación.

Consecuencia: `go.mod` en la raíz, con el módulo `github.com/LuisPV05/LuisMecanico`, y la tarea `install` ejecuta `go mod download`, que hoy no descarga nada porque no hay dependencias externas.

## Gestor de tareas

El comando `go` ejecuta subcomandos sueltos (`go build`, `go test`, `go mod download`), pero ninguno lee un fichero con tareas con nombre. `go generate` ejecuta comandos escritos en comentarios de los fuentes con la intención de crear o actualizar código Go. Los módulos de Go gestionan dependencias, no tareas, así que hace falta una herramienta aparte.

**Opciones:**

- make: fichero de tareas `Makefile`. Viene del sistema operativo, no con Go. La última versión que he encontrado es la 4.4.1, de 26-02-2023 ([NEWS](https://fossies.org/linux/make/NEWS)).
- just: fichero de tareas `justfile`. Es un binario ya compilado, escrito en Rust; en Windows se instala con `winget`. Versión 1.59.0, de 15-10-2026 ([crates.io](https://crates.io/crates/just)).
- Task: fichero de tareas `Taskfile.yml`. Se instala con `go install`, con el toolchain de Go que el proyecto ya necesita. Versión v3.54.0, de 01-10-2026 ([GitHub](https://github.com/go-task/task/releases/tag/v3.54.0)).
- Mage: fichero de tareas `magefile.go`, con las tareas escritas en Go. Se instala con `go install`. Versión v1.17.2, de 23-04-2026 ([GitHub](https://github.com/magefile/mage/releases)).

**Criterios:**

1. Mantenimiento: fecha de la última versión. Según el enunciado, una herramienta que no se mantiene con asiduidad aumenta la deuda técnica.
2. Coste de instalación: tiempo medido, porque se usará en los objetivos 5 y 6.
3. Coste de ejecución: sobrecoste medido frente a `go test ./...` directo.
4. Cuánto hay que mantener: tamaño del fichero de tareas y si las tareas son configuración o código del proyecto.
5. Si se instala con el toolchain de Go que el proyecto ya exige.

### Mediciones de los gestores de tareas

Se mide lo que cuesta tener cada herramienta disponible en la máquina de desarrollo, porque se usarán en entornos donde la velocidad de instalación importa (objetivos 5 y 6). Máquina: Windows, Git Bash, Go 1.27.3.

- Task: la instalé con `go install`. Tardó 4,629 s y el binario ocupa 75.064.320 B (unos 75 MB). Captura: [captura](../media/tareas-instalacion.png).
- Mage: la instalé con `go install` (v1.17.2). Tardó 2,497 s y el binario ocupa 7.531.520 B (unos 7,5 MB). Captura: [captura](../media/tareas-instalacion.png).
- just: lo instalé con `winget install --id Casey.Just -e`. Tardó 12,209 s (incluye la descarga) y el binario ocupa 5.922.304 B (unos 5,9 MB). Captura: [captura](../media/tareas-just-instalacion.png).
- make: ya estaba presente en el entorno (GNU Make 4.4.1, compilado para Cygwin). No se instala con Go, así que no medí su instalación. Captura: [captura](../media/tareas-instalacion.png).

#### Ejecución de la tarea `test` y tamaño del fichero de tareas

Se mide ejecutar `test` con cada herramienta frente a `go test ./...` directo, en una carpeta de prueba con el mismo código (la tarea `test` solo existe en esta prueba). El sobrecoste es la media de las dos diferencias con `go test` directo.

- `go test ./...` directo: 0,245 s y 0,248 s. Captura: [captura](../media/tareas-ejecucion.png).
- make: 0,295 s y 0,291 s, un sobrecoste medio de +0,047 s. El `Makefile` tiene 8 líneas. Captura: [captura](../media/tareas-ejecucion.png).
- Task: 0,337 s y 0,320 s, un sobrecoste medio de +0,082 s. El `Taskfile.yml` tiene 16 líneas. Captura: [captura](../media/tareas-ejecucion.png).
- Mage: 0,722 s la primera vez, porque compila el `magefile.go`, y 0,522 s la segunda, un sobrecoste medio de +0,376 s. El `magefile.go` tiene 21 líneas. Captura: [captura](../media/tareas-ejecucion.png).
- just: 0,344 s y 0,314 s, un sobrecoste medio de +0,077 s. El `justfile` tiene 8 líneas. Captura: [captura](../media/tareas-just-ejecucion.png).

**Decisión: Task.** Las medidas no dan un ganador en todos los criterios. make es el más rápido de ejecutar (+0,047 s frente a `go test` directo) y tiene el fichero más corto (8 líneas). Mage es el más barato de instalar (2,497 s), pero el más lento de ejecutar (+0,376 s). Task queda en medio en ejecución (+0,082 s) y en líneas (16), pero tarda 4,629 s en instalarse.

Se elige Task por estas razones:

- **Se instala con el toolchain de Go que el proyecto ya exige.** Con `go install` (4,629 s) no hace falta ningún gestor de paquetes del sistema ni descargar un binario aparte. make no viene con Go, y just lo conseguí con `winget` (12,209 s, con descarga). Es el único criterio medido en el que Task mejora a make y a just.
- **Frente a Mage:** Task tiene menos sobrecoste de ejecución (+0,082 s frente a +0,376 s) y un fichero más corto (16 líneas frente a 21). Además, Mage escribe las tareas como código Go que se compila en cada ejecución.
- **Mantenimiento:** Task tiene la versión v3.54.0, de 01-10-2026. Este criterio no lo separa de just (1.59.0, 15-10-2026) ni de Mage (v1.17.2, 23-04-2026), que también están activos. Frente a make, la última versión que he encontrado es la 4.4.1, de 26-02-2023: es un argumento débil, porque que no haya versiones nuevas también puede indicar madurez y no abandono.

Costes asumidos: 4,629 s de instalación y unos 35 ms más por ejecución que make. Si en los objetivos 5 y 6 el tiempo de instalación resulta un problema, se revisará la decisión.

Consecuencia: `Taskfile.yml` en la raíz con las tareas `install` y `check`, y `iv.yaml`

## Tarea `check`

El enunciado pide que `check` compruebe la sintaxis de los fuentes existentes y solo eso.

**Opciones:** se prueban cuatro comandos sobre tres proyectos mínimos creados fuera del repositorio: uno correcto, uno con un error de sintaxis (una llave sin cerrar) y otro con un error de tipos (una variable que no existe). Se anota el código de salida de cada comando y lo que tarda. Máquina: Windows, Git Bash, Go 1.27.3. Antes de medir se calienta la caché de Go.

- `gofmt -e .`: pasa en el proyecto correcto (salida 0, 0,294 s), falla con el error de sintaxis (salida 2, 0,269 s) y **pasa** con el error de tipos (salida 0, 0,271 s).
- `go build ./...`: pasa en el proyecto correcto (salida 0, 0,435 s), falla con el error de sintaxis (salida 1, 0,347 s) y falla con el error de tipos (salida 1, 0,357 s).
- `go vet ./...`: pasa en el proyecto correcto (salida 0, 0,399 s), falla con el error de sintaxis (salida 1, 0,410 s) y falla con el error de tipos (salida 1, 0,405 s).
- `go test ./...`: pasa en el proyecto correcto (salida 0, 0,343 s), falla con el error de sintaxis (salida 1, 0,352 s) y falla con el error de tipos (salida 1, 0,366 s).

Captura: [captura](../media/check-comandos.png).

**Criterios:**

1. Que detecte los errores de sintaxis.
2. Que no falle por nada que no sea sintaxis, porque la tarea comprueba solo eso.
3. Que no ejecute nada más, como los tests.
4. Tiempo de ejecución.

**Decisión: `gofmt -e . > /dev/null`.** Según la medición, es el único de los cuatro comandos que falla ante el error de sintaxis (salida 2) y pasa ante el error de tipos (salida 0): comprueba la sintaxis y nada más. `go build ./...`, `go vet ./...` y `go test ./...` fallan también ante el error de tipos, que no es un error de sintaxis, así que hacen más de lo que pide la tarea: según la documentación de Go, compilan, analizan el código o ejecutan los tests. Además, `gofmt -e .` es el más rápido en el proyecto correcto (0,294 s frente a 0,343 s - 0,435 s). Se redirige la salida estándar a `/dev/null` porque, según la documentación de `gofmt`, imprime el código formateado; los errores salen por la salida de error y siguen apareciendo.

Costes asumidos: `gofmt -e` no detecta los errores de tipos ni que el código no compile. Eso lo detectan `go build ./...` y `go test ./...`, que fallan ante los dos errores, como muestra la medición, y que se podrán añadir como tareas cuando se pidan.

## Dónde está cada cosa

- `go.mod`: módulo de Go.
- `Taskfile.yml`: tareas `install` y `check`.
- `iv.yaml`: clave `automatizar` con `fichero: Taskfile.yml` y `orden: task`.
- `README.md`: cómo instalar Task y ejecutar las tareas, y descripción de la entidad.
- `media/`: capturas de las mediciones.