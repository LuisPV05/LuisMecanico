# Objetivo 3: gestor de dependencias, gestor de tareas y tarea `check`

Este documento recoge cómo se eligieron las herramientas que automatizan las tareas del proyecto. Cada decisión parte de un problema planteado en un issue del Milestone 1 y se toma con criterios que se pueden comprobar.

## Gestor de dependencias

Go tiene un estándar propio para gestionar dependencias: los módulos (`go.mod` y `go.sum`), integrados en el comando `go` y oficiales desde Go 1.11 ([notas de la versión](https://go.dev/doc/go1.11#modules)). La guía del objetivo dice que, si hay un estándar, primero hay que conocerlo y después seguirlo, adaptando el resto de criterios.

## Gestor de tareas

**Criterios:**
1. Mantenimiento: fecha de la última versión. Según el enunciado, una herramienta que no se mantiene con asiduidad aumenta la deuda técnica. Se considera mantenida la que ha publicado una versión en el último año.
2. Coste de instalación: tiempo medido, porque se usará en los objetivos 5 y 6.
3. Coste de ejecución: sobrecoste medido frente a `go test ./...` directo.

**Opciones:**
- make: fichero de tareas `Makefile`. La última versión que he encontrado es la 4.4.1, de 26-02-2023 ([NEWS](https://fossies.org/linux/make/NEWS)).
- just: fichero de tareas `justfile`. Versión 1.59.0, de 15-10-2026 ([crates.io](https://crates.io/crates/just)).
- Task: fichero de tareas `Taskfile.yml`. Versión v3.54.0, de 01-10-2026 ([GitHub](https://github.com/go-task/task/releases/tag/v3.54.0)).
- Mage: fichero de tareas `magefile.go`, con las tareas escritas en Go. Versión v1.17.2, de 23-04-2026 ([GitHub](https://github.com/magefile/mage/releases)).
- xc: las tareas se escriben como bloques de código dentro de un fichero Markdown. Versión v0.9.0, de 05-01-2026 ([GitHub](https://github.com/joerdav/xc/releases)).
- goyek: biblioteca de Go; las tareas son código Go dentro del proyecto. Versión v3.0.1, de 09-12-2025 ([GitHub](https://github.com/goyek/goyek/releases)).
- mise: fichero de tareas `mise.toml`. Versión v2026.11.10, de 18-11-2026 ([GitHub](https://github.com/jdx/mise/releases)).

### Mediciones de los gestores de tareas

Se mide en Linux, porque es donde se probará y desplegará el proyecto: un contenedor Docker `golang:latest` (Debian 13, Go 1.27.4), nuevo para cada herramienta. Cada una se instala con el método que recomienda su documentación y los tiempos incluyen la descarga. Es una medición por herramienta.

- make: ya viene en la imagen (GNU Make 4.4.1), así que no hay instalación que medir. Captura: [captura](../media/tareas-linux-instalacion.png).
- just: con el script de su documentación (`curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to ~/bin`). Tardó 1,193 s. Captura: [captura](../media/tareas-linux-instalacion.png).
- goyek: `go get github.com/goyek/goyek/v3` (v3.0.1), porque es una biblioteca. Tardó 2,029 s. Captura: [captura](../media/tareas-linux-instalacion.png).
- Task: con el script de su documentación (`sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin`). Tardó 5,302 s. Captura: [captura](../media/tareas-linux-instalacion.png).
- xc: `go install github.com/joerdav/xc/cmd/xc@latest` (v0.9.0). Tardó 9,565 s. Captura: [captura](../media/tareas-linux-instalacion.png).
- mise: `curl -fsSL https://mise.run | sh`. Tardó 11,013 s. Captura: [captura](../media/tareas-linux-instalacion.png).
- Mage: `go install github.com/magefile/mage@latest` (v1.17.2). Tardó 19,419 s. Captura: [captura](../media/tareas-linux-instalacion.png).

#### Ejecución de la tarea `test`

Se mide en el mismo contenedor Linux, con una tarea `test` que ejecuta `go test ./...` en cada herramienta, usando `hyperfine` (10 ejecuciones de calentamiento y 50 medidas por comando, todas en la misma tanda). El `go test ./...` directo se mide al principio y al final para comprobar que la máquina no cambia de velocidad: 109,3 ± 7,0 ms y 103,9 ± 2,8 ms. Se toma su media, 106,6 ms, y el sobrecoste de cada herramienta es su media menos esa. Una diferencia menor que la desviación no se considera medible.

- just: 106,5 ± 3,1 ms, sin diferencia medible (−0,1 ms).
- make: 108,7 ± 4,8 ms, +2,1 ms.
- xc: 108,7 ± 4,3 ms, +2,1 ms.
- mise: 123,9 ± 5,2 ms, +17,3 ms.
- Task: 163,2 ± 5,1 ms, +56,6 ms.
- goyek: 172,1 ± 3,1 ms, +65,5 ms. Se ejecuta con `go run .`, que compila el programa en cada ejecución.
- Mage: 179,8 ± 5,4 ms, +73,2 ms. Compila el `magefile.go`.

Captura: [captura](../media/tareas-linux-ejecucion.png).

**Decisión: just.** Cumple el criterio de mantenimiento y, entre las herramientas que lo cumplen, es la mejor en los dos criterios medidos: es la más rápida de instalar (1,193 s) y no añade un tiempo medible al ejecutar (−0,1 ms, dentro del ruido). Su última versión es la 1.59.0, de 15-10-2026.

- **Frente a make:** en lo medido empatan, porque ejecuta igual (+2,1 ms, dentro del ruido) y ya viene en la imagen, sin coste de instalación. Lo que decide es el criterio de mantenimiento del enunciado: la última versión de make que he encontrado es la 4.4.1, de 26-02-2023, hace más de tres años, y el enunciado dice que una herramienta que no se mantiene con asiduidad aumenta la deuda técnica. Es el único criterio en el que se separan.
- **Frente a Task, Mage, goyek, mise y xc:** todas están mantenidas, pero tardan más en instalarse (5,302 s, 19,419 s, 2,029 s, 11,013 s y 9,565 s). En ejecución, Task suma +56,6 ms, Mage +73,2 ms, goyek +65,5 ms y mise +17,3 ms. xc ejecuta igual que just (+2,1 ms, dentro del ruido), pero tarda más en instalarse.

Costes asumidos: hay que instalar just (1,193 s), mientras que make ya está en la imagen. Si en los objetivos 5 y 6 el tiempo de instalación resulta un problema, se revisará la decisión.

Consecuencia: `justfile` en la raíz con la tarea `check`, y `iv.yaml` con `fichero: justfile` y `orden: just` en la clave `automatizar`.

## Tarea `check`

El enunciado pide que `check` compruebe la sintaxis de los fuentes existentes y solo eso.

**Decisión: `gofmt -e . > /dev/null`.** `gofmt -e` informa de todos los errores de sintaxis sin compilar el código. `go build`, `go vet` y `go test` también compilan, así que van más allá de la sintaxis. Se redirige la salida estándar porque `gofmt` imprime el código formateado; los errores salen por la salida de error y siguen apareciendo.

Costes asumidos: no detecta errores de tipos ni que el código no compile.