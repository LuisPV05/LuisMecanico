# Milestones

## Milestone 0: Modelo de reparaciones y recursos del taller

El modelo del dominio del taller se obtiene aplicando una metodología de diseño de dominio a partir de las [historias de usuario](historias-usuarios.md), las [jornadas de usuario](user-journeys.md) y lo descrito en el [objetivo 0](objetivo-0.md). Se considera que se ha seguido ese proceso cuando, junto al modelo, existe un documento de decisiones de diseño que cualquier otra persona pueda revisar y contrastar contra las historias de usuario de las que parte, sin necesidad de juzgar subjetivamente si el código "está bien". A partir de ahí se entrega el modelo en sí, sin ninguna lógica de negocio todavía, junto con ese documento de decisiones.

## Milestone 1: Lógica de negocio y comprobación automática

El objetivo de este hito es incorporar una primera pieza de comportamiento al sistema tomando como punto de partida la estructura del Milestone 0. El desarrollo se apoyará desde el comienzo en una batería de tests automatizados, que será la referencia para validar que el trabajo está realmente terminado.No se establece previamente cuál será la regla exacta ni cómo deberá resolverse técnicamente, ya que ambas cuestiones se concretarán mientras se implementa la funcionalidad. La aceptación del trabajo dependerá de que la implementación resultante disponga de sus pruebas correspondientes y de que estas finalicen correctamente. Por tanto, el entregable consiste en esa lógica inicial y en el conjunto de tests que demuestra su funcionamiento.

| Pregunta | Milestone 0 | Milestone 1 |
|---|---|---|
| **Se considera conseguido cuando...** | existe el modelo con su documentación de decisiones y se puede comprobar que se ha obtenido siguiendo la metodología de diseño de dominio, y no solo que el resultado parezca correcto | existe la lógica sobre el modelo, con sus pruebas automáticas funcionando, y se puede comprobar que se ha desarrollado siguiendo la metodología dirigida por pruebas |