# Objetivo 3: gestor de dependencias, gestor de tareas y tarea `check`

Este documento recoge cómo se eligieron las herramientas que automatizan las tareas del proyecto. Cada decisión parte de un problema planteado en un issue del Milestone 1 y se toma con criterios que se pueden comprobar.

## Punto de partida

- El proyecto está escrito en Go y solo importa `time`, de la biblioteca estándar: no tiene dependencias externas.
- No existe `go.mod`, así que `go build ./...` falla porque no hay módulo principal.
- `iv.yaml` necesita apuntar a un fichero de tareas que declare `install`, `test` y `check`.

## Gestor de dependencias

Go tiene un estándar propio para gestionar dependencias: los módulos (`go.mod` y `go.sum`), integrados en el comando `go` y oficiales desde Go 1.11. La guía del objetivo dice que, si hay un estándar, primero hay que conocerlo y después seguirlo, adaptando el resto de criterios. Por eso no se comparan otras herramientas.

**Criterios:**

1. Si es el estándar del lenguaje.
2. Si hay que instalar algo aparte del toolchain de Go.
3. Cuánto cuesta dejar el proyecto listo, porque se usará en los objetivos 5 y 6, donde la velocidad importa.
4. Número de dependencias del proyecto: hoy, ninguna externa.

### Mediciones del gestor de dependencias

Se mide cuánto cuesta dejar el proyecto listo en una máquina limpia. Máquina: Windows, Git Bash, Go 1.27.3 (`go version go1.27.3 windows/amd64`).

| Herramienta | Instalación aparte de Go | `go mod download` o equivalente | Compilar el proyecto | Captura |
|---|---|---|---|---|
| Módulos de Go | Ninguna | 0,152 s | 0,491 s | [captura](../media/dep-go-modules.png) |

**Decisión: módulos de Go.** Son el estándar del lenguaje, vienen en el propio comando `go` y no hay nada que instalar. Preparar el proyecto cuesta 0,152 s y compilarlo 0,491 s, y el proyecto no tiene dependencias externas, así que no hay ninguna razón para añadir otra herramienta. Tampoco se usan los módulos como gestor de tareas: gestionan dependencias, no tareas, y por eso hace falta una herramienta aparte, que se elige a continuación.

Consecuencia: `go.mod` en la raíz, con el módulo `github.com/LuisPV05/LuisMecanico`, y la tarea `install` ejecuta `go mod download`, que hoy no descarga nada porque no hay dependencias externas.

## Gestor de tareas

El comando `go` ejecuta subcomandos sueltos (`go build`, `go test`, `go mod download`), pero ninguno lee un fichero con tareas con nombre. `go generate` ejecuta comandos escritos en comentarios de los fuentes con la intención de crear o actualizar código Go. Los módulos de Go gestionan dependencias, no tareas, así que hace falta una herramienta aparte.

**Opciones:**

| Opción | Fichero de tareas | Instalación | Última versión |
|---|---|---|---|
| make | `Makefile` | Del sistema operativo, ajena a Go | 4.4.1, 26-02-2023 ([NEWS](https://fossies.org/linux/make/NEWS)) |
| Task | `Taskfile.yml` | `go install`, con el toolchain de Go que el proyecto ya necesita | v3.54.0, 01-10-2026 ([GitHub](https://github.com/go-task/task/releases/tag/v3.54.0)) |
| Mage | `magefile.go`, con las tareas escritas en Go | `go install` | v1.17.2, 23-04-2026 ([GitHub](https://github.com/magefile/mage/releases)) |

**Criterios:**

1. Mantenimiento: fecha de la última versión. Según el enunciado, una herramienta que no se mantiene con asiduidad aumenta la deuda técnica.
2. Coste de instalación: tiempo y tamaño medidos, porque se usará en los objetivos 5 y 6.
3. Coste de ejecución: sobrecoste medido frente a `go test ./...` directo.
4. Cuánto hay que mantener: tamaño del fichero de tareas y si las tareas son configuración o código del proyecto.
5. Si se instala con el toolchain de Go que el proyecto ya exige.

### Mediciones de los gestores de tareas

Se mide lo que cuesta tener cada herramienta disponible en una máquina limpia, porque se usarán en entornos donde la velocidad de instalación importa (objetivos 5 y 6). Máquina: Windows, Git Bash, Go 1.27.3.

| Herramienta | Cómo se obtiene | Tiempo | Tamaño en disco | Captura |
|---|---|---|---|---|
| Task | `go install` | 4,629 s | 75.064.320 B | [captura](../media/tareas-instalacion.png) |
| Mage | `go install` (v1.17.2) | 2,497 s | 7.531.520 B | [captura](../media/tareas-instalacion.png) |
| make | Ya presente en el entorno (GNU Make 4.4.1, compilado para Cygwin); no se instala con Go | No medido | No medido | [captura](../media/tareas-instalacion.png) |

Limitación: Task y Mage se compilan con Go durante la instalación, y make ya estaba disponible en el entorno. Los tiempos indican el coste real de cada vía, no una comparación entre binarios idénticos. La presencia de make depende del entorno: no viene con Go.

#### Ejecución de la tarea `test` y tamaño del fichero de tareas

Se crea una carpeta de prueba fuera del repositorio con el mismo `go.mod`, el mismo código y un fichero de tareas por herramienta (`Taskfile.yml`, `Makefile`, `magefile.go`) con las tres tareas `install`, `test` y `check`. Se calienta la caché de Go una vez y se ejecuta `test` dos veces con cada herramienta. Para comparar, se mide también `go test ./...` directo. El sobrecoste es la media de las diferencias con `go test` directo en cada pasada.

| Herramienta | 1.ª ejecución | 2.ª ejecución | Sobrecoste medio | Líneas del fichero de tareas | Captura |
|---|---|---|---|---|---|
| `go test ./...` directo | 0,245 s | 0,248 s | — | — | [captura](../media/tareas-ejecucion.png) |
| make | 0,295 s | 0,291 s | +0,047 s | 8 | [captura](../media/tareas-ejecucion.png) |
| Task | 0,337 s | 0,320 s | +0,082 s | 16 | [captura](../media/tareas-ejecucion.png) |
| Mage | 0,722 s (compila el `magefile.go`) | 0,522 s | +0,376 s | 21 | [captura](../media/tareas-ejecucion.png) |

Limitaciones: el proyecto todavía no tiene tests, así que las diferencias reflejan sobre todo el tiempo de arranque de cada herramienta. Son dos ejecuciones por herramienta y las diferencias entre make y Task (unas decenas de milisegundos) pueden ser ruido. El `Taskfile.yml` tiene una línea `desc` por tarea, que sirve para `task --list`; el `Makefile` no tiene equivalente.

**Decisión: Task.** Las medidas no dan un ganador en todos los criterios. make es el más rápido de ejecutar (+0,047 s frente a `go test` directo) y tiene el fichero más corto (8 líneas). Mage es el más barato de instalar (2,497 s y 7,5 MB), pero el más lento de ejecutar (+0,376 s). Task queda en medio en ejecución (+0,082 s) y en líneas (16), pero su binario es el más grande (75 MB) y tarda 4,629 s en instalarse.

Se elige Task por tres razones:

- **Frente a make:** make 4.4.1 es de 26-02-2023 y no ha tenido versión nueva desde entonces, mientras que Task tiene v3.54.0 del 01-10-2026. Según el enunciado, una herramienta que no se mantiene con asiduidad aumenta la deuda técnica. Es el argumento más débil: que no haya versiones nuevas también puede indicar madurez y no abandono. Además, make no viene con Go.
- **Frente a Mage:** Task tiene menos sobrecoste de ejecución (+0,082 s frente a +0,376 s) y un fichero más corto (16 líneas frente a 21). Mage, además, escribe las tareas como código Go que se compila en cada ejecución.
- **Con el proyecto:** Task se instala con el toolchain de Go que el proyecto ya exige y declara las tareas como configuración en `Taskfile.yml`.

Costes asumidos: 75 MB y 4,629 s de instalación, y unos 35 ms más por ejecución que make. Si en los objetivos 5 y 6 el tiempo de instalación resulta un problema, se revisará la decisión.

## Tarea `check`

El enunciado pide que `check` compruebe la sintaxis de los fuentes existentes y solo eso.

**Opciones:** se prueban cuatro comandos sobre tres proyectos mínimos creados fuera del repositorio: uno correcto, uno con un error de sintaxis (una llave sin cerrar) y otro con un error de tipos (una variable que no existe). Se anota el código de salida de cada comando y lo que tarda. Máquina: Windows, Git Bash, Go 1.27.3. Antes de medir se calienta la caché de Go.

| Comando | Proyecto correcto | Error de sintaxis | Error de tipos |
|---|---|---|---|
| `gofmt -e .` | Pasa (0, 0,294 s) | Falla (2, 0,269 s) | Pasa (0, 0,271 s) |
| `go build ./...` | Pasa (0, 0,435 s) | Falla (1, 0,347 s) | Falla (1, 0,357 s) |
| `go vet ./...` | Pasa (0, 0,399 s) | Falla (1, 0,410 s) | Falla (1, 0,405 s) |
| `go test ./...` | Pasa (0, 0,343 s) | Falla (1, 0,352 s) | Falla (1, 0,366 s) |

Captura: [captura](../media/check-comandos.png). Cada celda indica si el comando pasa o falla, su código de salida y su tiempo.

Limitaciones: una sola ejecución por comando y proyectos de pocas líneas, así que los tiempos reflejan sobre todo el arranque de cada herramienta y las diferencias de decenas de milisegundos pueden ser ruido.

**Criterio:** que compruebe la sintaxis y nada más.

**Decisión: `gofmt -e . > /dev/null`.** Es el único comando que detecta los errores de sintaxis sin detectar los de tipos ni hacer nada más. Los demás van más allá de la sintaxis. Se descarta la salida estándar porque, sin eso, imprime todo el código formateado, y los errores salen por la salida de error. El coste asumido es que no comprueba que el código compile; eso lo hace `task test`.

## Dónde está cada cosa

- `go.mod`: módulo de Go.
- `Taskfile.yml`: tareas `install`, `test` y `check`.
- `iv.yaml`: clave `automatizar` con `fichero: Taskfile.yml` y `orden: task`.
- `README.md`: cómo instalar Task y ejecutar las tareas, y descripción de la entidad.
- `media/`: capturas de las mediciones.