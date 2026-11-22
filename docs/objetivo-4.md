# Objetivo 4: Tests

Una biblioteca de aserciones compara el resultado obtenido con el esperado y da el mensaje del fallo. Un test runner descubre los tests, los ejecuta y produce el informe de pasados, fallados y cobertura.

## Biblioteca de aserciones

### Criterios

- Mantenimiento: ha publicado versión en el último año.
- Dependencias nuevas que añade a `go.mod` y `go.sum`.
- Coste: tiempo de descarga y compilación, y sobrecarga en la ejecución frente a usar `testing` solo.
- Permite comprobar errores con `errors.Is` y comparar structs y `time.Time`.
- No escribe ficheros junto a los fuentes (lo exige el objetivo 5).

## Test runner

### Criterios

- Mantenimiento: ha publicado versión en el último año.
- Tiempo de instalación en Linux con el método oficial.
- Sobrecarga de ejecución frente a `go test`.
- Funciona con el directorio de fuentes en solo lectura (lo exige el objetivo 5).
- Informa de tests pasados y fallados, y de cobertura.