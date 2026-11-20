# Objetivo 3: gestor de dependencias, gestor de tareas y tarea `check`

Este documento recoge cómo se eligieron las herramientas que automatizan las tareas del proyecto. Cada decisión parte de un problema planteado en un issue del Milestone 1 y se toma con criterios que se pueden comprobar.

## Gestor de dependencias

Go tiene un estándar propio para gestionar dependencias: los módulos (`go.mod` y `go.sum`), integrados en el comando `go` y oficiales desde Go 1.11 ([notas de la versión](https://go.dev/doc/go1.11#modules)). La guía del objetivo dice que, si hay un estándar, primero hay que conocerlo y después seguirlo, adaptando el resto de criterios.

### Medición de los módulos de Go

Esta medición no sirve para comparar, solo para comprobar cuánto cuesta dejar el proyecto listo, porque se usará en los objetivos 5 y 6, donde la velocidad importa. Máquina: la de desarrollo, con Windows, Git Bash y Go 1.27.3 (`go version go1.27.3 windows/amd64`).

Con los módulos de Go no hay que instalar nada aparte de Go. `go mod download` tarda 0,152 s (no descarga nada, porque el proyecto no tiene dependencias externas) y compilar el proyecto con `go build ./...` tarda 0,491 s. Captura: [captura](../media/dep-go-modules.png).

**Decisión: módulos de Go.** Son el estándar del lenguaje y no hay nada que instalar aparte de Go. 

## Gestor de tareas

El comando `go` ejecuta subcomandos sueltos (`go build`, `go test`, `go mod download`), pero ninguno lee un fichero con tareas con nombre. `go generate` ejecuta comandos escritos en comentarios de los fuentes con la intención de crear o actualizar código Go. Los módulos de Go gestionan dependencias, no tareas, así que hace falta una herramienta aparte.

**Criterios:**
1. Mantenimiento: fecha de la última versión. Según el enunciado, una herramienta que no se mantiene con asiduidad aumenta la deuda técnica.
2. Coste de instalación: tiempo medido, porque se usará en los objetivos 5 y 6.
3. Coste de ejecución: sobrecoste medido frente a `go test ./...` directo.
4. Cuánto hay que mantener: tamaño del fichero de tareas y si las tareas son configuración o código del proyecto.
5. Si se instala con el toolchain de Go que el proyecto ya exige.

**Opciones:**

- make: fichero de tareas `Makefile`. No viene con Go, pero está en la mayoría de las imágenes y runners de CI. La última versión que he encontrado es la 4.4.1, de 26-02-2023 ([NEWS](https://fossies.org/linux/make/NEWS)).
- just: fichero de tareas `justfile`. Es un binario ya compilado, escrito en Rust; en Windows se instala con `winget`. Versión 1.59.0, de 15-10-2026 ([crates.io](https://crates.io/crates/just)).
- Task: fichero de tareas `Taskfile.yml`. Se instala con `go install`, con el toolchain de Go que el proyecto ya necesita. Versión v3.54.0, de 01-10-2026 ([GitHub](https://github.com/go-task/task/releases/tag/v3.54.0)).
- Mage: fichero de tareas `magefile.go`, con las tareas escritas en Go. Se instala con `go install`. Versión v1.17.2, de 23-04-2026 ([GitHub](https://github.com/magefile/mage/releases)).
- xc: las tareas se escriben como bloques de código dentro de un fichero Markdown, normalmente el `README.md`. Se instala con `go install`. Versión v0.9.0, de 05-01-2026 ([GitHub](https://github.com/joerdav/xc/releases)).
- goyek: no es un programa aparte, sino una biblioteca de Go; las tareas son código Go dentro del propio proyecto. Se añade con `go get`. Versión v3.0.1, de 09-12-2025 ([GitHub](https://github.com/goyek/goyek/releases)).
- mise: fichero de tareas `mise.toml`. Es un binario ya compilado, escrito en Rust; en Windows se instala con `winget`. Última versión v2026.11.9, de 09-11-2026 ([GitHub](https://github.com/jdx/mise/releases)); winget instala la 2026.11.2, de 05-11-2026.

### Mediciones de los gestores de tareas

Se mide lo que cuesta tener cada herramienta disponible en la máquina de desarrollo, porque se usarán en entornos donde la velocidad de instalación importa (objetivos 5 y 6). Máquina: Windows, Git Bash, Go 1.27.3.

- Task: la instalé con `go install`. Tardó 4,629 s y el binario ocupa 75.064.320 B (unos 75 MB). Captura: [captura](../media/tareas-instalacion.png).
- Mage: la instalé con `go install` (v1.17.2). Tardó 2,497 s y el binario ocupa 7.531.520 B (unos 7,5 MB). Captura: [captura](../media/tareas-instalacion.png).
- just: lo instalé con `winget install --id Casey.Just -e`. Tardó 12,209 s (incluye la descarga) y el binario ocupa 5.922.304 B (unos 5,9 MB). Captura: [captura](../media/tareas-just-instalacion.png).
- make: ya estaba presente en el entorno (GNU Make 4.4.1, compilado para Cygwin). No se instala con Go, así que no medí su instalación. Captura: [captura](../media/tareas-instalacion.png).
- xc: la instalé con `go install github.com/joerdav/xc/cmd/xc@latest` (v0.9.0). Tardó 52,488 s (incluye la descarga de sus dependencias) y el binario ocupa 9.081.856 B (unos 9,1 MB). Captura: [captura](../media/tareas-xc-instalacion.png).
- goyek: la añadí a un módulo nuevo con `go get github.com/goyek/goyek/v3` (v3.0.1). Tardó 22,277 s (incluye la descarga) y no genera un binario aparte. Captura: [captura](../media/tareas-goyek-instalacion.png).
- mise: lo instalé con `winget install --id jdx.mise -e` (2026.11.2), tras desinstalar una instalación previa. Tardó 17,775 s (incluye la descarga de 81,6 MB) y el binario ocupa 252.385.912 B (unos 252 MB). Captura: [captura](../media/tareas-mise-instalacion.png).

#### Ejecución de la tarea `test` y tamaño del fichero de tareas

Se mide ejecutar `test` con cada herramienta frente a `go test ./...` directo, en una carpeta de prueba con el mismo código (la tarea `test` solo existe en esta prueba). El sobrecoste es la media de las dos diferencias con `go test` directo.

- `go test ./...` directo: 0,245 s y 0,248 s. Captura: [captura](../media/tareas-ejecucion.png).
- make: 0,295 s y 0,291 s, un sobrecoste medio de +0,047 s. El `Makefile` tiene 8 líneas. Captura: [captura](../media/tareas-ejecucion.png).
- Task: 0,337 s y 0,320 s, un sobrecoste medio de +0,082 s. El `Taskfile.yml` tiene 16 líneas. Captura: [captura](../media/tareas-ejecucion.png).
- Mage: 0,722 s la primera vez, porque compila el `magefile.go`, y 0,522 s la segunda, un sobrecoste medio de +0,376 s. El `magefile.go` tiene 21 líneas. Captura: [captura](../media/tareas-ejecucion.png).
- just: 0,344 s y 0,314 s, un sobrecoste medio de +0,077 s. El `justfile` tiene 8 líneas. Captura: [captura](../media/tareas-just-ejecucion.png).
- xc: 0,582 s y 0,348 s, un sobrecoste medio de +0,067 s. El `README.md` con las tareas tiene 21 líneas. Captura: [captura](../media/tareas-xc-mise-ejecucion.png).
- mise: 1,073 s y 0,692 s, un sobrecoste medio de +0,484 s. El `mise.toml` tiene 8 líneas. Captura: [captura](../media/tareas-xc-mise-ejecucion.png).
- goyek: 0,438 s y 0,431 s, con su propio `go test ./...` directo de 0,248 s y 0,242 s, un sobrecoste medio de +0,190 s. Se ejecuta con `go run .`, que compila el programa en cada ejecución. El fichero `build/main.go` tiene 23 líneas y solo define la tarea `test`, no las tres. Captura: [captura](../media/tareas-goyek-ejecucion.png).

**Decisión: Task.** Las medidas no dan un ganador en todos los criterios. make es el más rápido de ejecutar (+0,047 s frente a `go test` directo) y just, mise y make tienen el fichero más corto (8 líneas). Mage es el más barato de instalar (2,497 s), pero el más lento de ejecutar entre los que se instalan con Go (+0,376 s), y mise es todavía más lento (+0,484 s). Task queda entre los más rápidos de ejecutar (+0,082 s, a menos de 40 ms de make) y en líneas (16), pero tarda 4,629 s en instalarse.

Se elige Task por estas razones:

- **Se instala con el toolchain de Go que el proyecto ya exige.** Con `go install` (4,629 s) no hace falta ningún gestor de paquetes del sistema ni descargar un binario aparte. make no se instala con Go, y just y mise los conseguí con `winget` (12,209 s y 17,775 s, con descarga). Que make esté ya en muchos entornos de CI lo hace fácil de conseguir allí, así que este argumento pesa menos. Mage, xc y goyek también se instalan con Go, y en ellos decide el resto de criterios.
- **Frente a Mage y goyek:** Task tiene menos sobrecoste de ejecución (+0,082 s frente a +0,376 s y +0,190 s) y un fichero más corto (16 líneas frente a 21 y a 23, y el de goyek solo define una tarea). Además, ambos escriben las tareas como código Go que se compila en cada ejecución.
- **Frente a xc:** xc se instala con `go install`, pero tardó 52,488 s frente a 4,629 s de Task. Esa comparación no es exacta, porque el tiempo de xc incluye la descarga de sus dependencias. En ejecución, la diferencia (+0,067 s frente a +0,082 s) está dentro del ruido de la medición. xc escribe las tareas dentro de un Markdown.
- **Frente a mise:** mise tiene el mayor sobrecoste de ejecución (+0,484 s) y tardó 17,775 s en instalarse con `winget`.
- **Mantenimiento:** Task tiene la versión v3.54.0, de 01-10-2026. Este criterio no lo separa de just (1.59.0, 15-10-2026), de Mage (v1.17.2, 23-04-2026), de goyek ni de mise, que también están activos. Frente a make, la última versión que he encontrado es la 4.4.1, de 26-02-2023: es un argumento débil, porque que no haya versiones nuevas también puede indicar madurez y no abandono.

Costes asumidos: 4,629 s de instalación y unos 35 ms más por ejecución que make. Si en los objetivos 5 y 6 el tiempo de instalación resulta un problema, se revisará la decisión.

Consecuencia: `Taskfile.yml` en la raíz con las tareas `install` y `check`, y `iv.yaml` con `fichero: Taskfile.yml` y `orden: task` en la clave `automatizar`.

## Tarea `check`

El enunciado pide que `check` compruebe la sintaxis de los fuentes existentes y solo eso.

**Opciones:** `gofmt -e .`, `go build ./...`, `go vet ./...` y `go test ./...`. Se prueban sobre tres proyectos mínimos (uno correcto, uno con una llave sin cerrar y otro con una variable que no existe) en Windows, Git Bash, Go 1.27.3. Captura: [captura](../media/check-comandos.png).

- `gofmt -e .`: falla con el error de sintaxis (salida 2) y **pasa** con el error de tipos (salida 0). Tarda 0,294 s en el proyecto correcto.
- `go build ./...`, `go vet ./...` y `go test ./...`: fallan con los dos errores (salida 1). Tardan entre 0,343 s y 0,435 s en el proyecto correcto.

**Criterio:** que falle ante errores de sintaxis y no ante nada más, porque la tarea comprueba solo eso.

**Decisión: `gofmt -e . > /dev/null`.** Es el único que falla ante el error de sintaxis y pasa ante el de tipos, y además es el más rápido. Se redirige la salida estándar porque `gofmt` imprime el código formateado; los errores salen por la salida de error y siguen apareciendo.

Costes asumidos: `gofmt -e` no detecta los errores de tipos ni que el código no compile. Eso lo detectan `go build ./...` y `go test ./...`, que fallan ante los dos errores, como muestra la medición, y que se podrán añadir como tareas cuando se pidan.