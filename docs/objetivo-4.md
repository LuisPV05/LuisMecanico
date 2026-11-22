# Objetivo 4: Tests

Una biblioteca de aserciones compara el resultado obtenido con el esperado y da el mensaje del fallo. Un test runner descubre los tests, los ejecuta y produce el informe de pasados, fallados y cobertura.

## Biblioteca de aserciones

### Criterios

- Mantenimiento: ha publicado versión en el último año.
- Dependencias nuevas que añade a `go.mod` y `go.sum`.
- Coste: tiempo de descarga y compilación, y sobrecarga en la ejecución frente a usar `testing` solo.
- Permite comprobar errores con `errors.Is` y comparar structs y `time.Time`.
- No escribe ficheros junto a los fuentes (lo exige el objetivo 5).

### Opciones

Última versión publicada de cada candidata, obtenida con `go list -m` en Linux (imagen `golang:latest`):

![Versiones y fechas de las opciones](../media/opciones-test-versiones.png)

Cumplen el criterio de mantenimiento:

- `testing`, de la biblioteca estándar de Go. Es la referencia con la que se compara.
- testify v1.12.1 (17-08-2026).
- gomega v1.44.0 (24-09-2026).
- shoenig/test v1.13.2 (21-04-2026).

No lo cumplen y se descartan:

- quicktest v1.14.6 (01-08-2023).
- is v1.4.1 (16-05-2022).
- gotest.tools/v3 v3.5.2 (05-09-2024).
- alecthomas/assert v2.11.0 (17-09-2024).

Sin comprobar: go-testdeep, porque `go list` no encuentra versión con la ruta `/v2`.

### Medidas

Un mismo test escrito con cada opción (comparar una estructura con `time.Time` y comprobar un error con `errors.Is`), en Linux (imagen `golang:latest`), con las cachés de módulos y de compilación vacías en cada una. La descarga es `go mod tidy`; la ejecución se mide con hyperfine sobre el binario de test ya compilado (500 ejecuciones, 20 de calentamiento).

![Medidas de las opciones de aserciones](../media/aserciones-medidas.png)

- `testing`: descarga 0,585 s, compilación 6,186 s, 0 módulos, 0 líneas en `go.sum`, ejecución 2,8 ± 0,4 ms.
- testify: descarga 1,421 s, compilación 11,217 s, 3 módulos, 5 líneas en `go.sum`, ejecución 3,9 ± 0,6 ms.
- gomega: descarga 2,512 s, compilación 12,283 s, 12 módulos, 17 líneas en `go.sum`, ejecución 4,4 ± 0,6 ms.
- shoenig/test: descarga 1,104 s, compilación 7,045 s, 2 módulos, 4 líneas en `go.sum`, ejecución 4,1 ± 0,7 ms.

Resultado por criterio:

- Errores y estructuras: las cuatro compilan y pasan el test.
- Escritura junto a los fuentes: las cuatro pasan con los fuentes en solo lectura, ejecutadas como usuario sin permisos de escritura.
- Dependencias nuevas: testify 3 módulos, gomega 12, shoenig/test 2.
- Coste de descarga y compilación frente a `testing`: testify +5,9 s, gomega +8,0 s, shoenig/test +1,4 s.
- Sobrecarga de ejecución frente a `testing`: testify +1,1 ms, gomega +1,6 ms, shoenig/test +1,3 ms. Las diferencias entre las tres están dentro de la desviación.

### Decisión

Me quedo con `testing`, sin biblioteca externa.

Las cuatro opciones cumplen mantenimiento, comprobar errores con `errors.Is`, comparar estructuras con `time.Time` y no escribir junto a los fuentes. Lo que las separa es el coste: `testing` no añade ningún módulo ni tiempo de descarga o compilación, mientras que shoenig/test añade 2 módulos y 1,4 s, testify 3 módulos y 5,9 s y gomega 12 módulos y 8,0 s. La sobrecarga de ejecución (entre 1,1 y 1,6 ms) queda dentro de la desviación y no desempata.

Descarto shoenig/test, testify y gomega porque ninguna aporta algo que mis criterios premien y todas añaden dependencias y tiempo. El precio de `testing` es escribir las comparaciones a mano, que mis criterios no medían.

## Test runner

### Criterios

- Mantenimiento: ha publicado versión en el último año.
- Tiempo de instalación en Linux con el método oficial.
- Sobrecarga de ejecución frente a `go test`.
- Funciona con el directorio de fuentes en solo lectura (lo exige el objetivo 5).
- Informa de tests pasados y fallados, y de cobertura.

### Opciones

Las versiones y fechas salen de la misma captura.

Cumplen el criterio de mantenimiento:

- `go test`, incluido en Go. Es la referencia con la que se compara.
- ginkgo v2.34.0 (02-11-2026).
- gotest v0.0.7 (10-01-2026).

No lo cumplen y se descartan:

- gotestsum v1.13.0 (11-09-2025).
- richgo v0.3.12 (26-02-2023).

### Medidas

Un mismo test escrito para cada opción, en Linux (imagen `golang:latest`). La instalación se hace con `go install` y las cachés vacías. La ejecución se mide con hyperfine (30 ejecuciones, 3 de calentamiento) sobre `go test -count=1 ./...`, o el comando equivalente de cada runner, con la compilación ya en caché; `go test` se mide al principio y al final para detectar deriva de la máquina. El solo lectura se prueba ejecutando como un usuario sin permisos de escritura en el directorio de los fuentes, con la caché de compilación en otro sitio. La variante de ginkgo usa su propio estilo de tests (`Describe` e `It`).

![Medidas de los test runners](../media/runner-medidas.png)

- `go test`: incluido en Go. Ejecución 254,3 ± 13,9 ms al inicio y 254,8 ± 9,3 ms al final (referencia: 254,6 ms).
- gotest: instalación 6,311 s. Ejecución 272,0 ± 24,0 ms.
- ginkgo: instalación 13,458 s, más 4,595 s de descarga de su biblioteca en el proyecto. Ejecución 739,1 ± 22,2 ms.

Resultado por criterio:

- Instalación: `go test` ninguna; gotest 6,3 s; ginkgo 13,5 s.
- Sobrecarga de ejecución frente a `go test`: gotest +17,4 ms, dentro de su desviación y por tanto no distinguible del ruido; ginkgo +484,5 ms.
- Solo lectura: `go test` y gotest pasan (código 0). `ginkgo` falla (código 1), con y sin `--output-dir`, porque copia el binario de test compilado al directorio del paquete (`permission denied`). Compilando el test fuera con `go test -c -o` y pasándole el binario a ginkgo, pasa (código 0), pero son dos comandos.
- Informe: `go test` y gotest muestran tests pasados y fallados y el porcentaje de cobertura (100 % en esta prueba); ginkgo muestra specs pasados y fallados. La cobertura de ginkgo no se midió, porque falló antes por el solo lectura.

![Ginkgo con el binario compilado fuera de los fuentes](../media/ginkgo-solo-lectura.png)