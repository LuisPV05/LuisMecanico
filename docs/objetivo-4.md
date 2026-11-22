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