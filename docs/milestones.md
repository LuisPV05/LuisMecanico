# Milestones

## Milestone 0: Modelo de reparaciones y recursos del taller

A partir de las [historias de usuario](historias-usuarios.md) y las [jornadas de usuario](user-journeys.md) se plantean issues que enuncian los problemas de su dominio, identificando los conceptos clave que aparecen en ellas (por ejemplo, qué es una cita, qué es un recurso, qué tipos de recurso existen). Esos issues guían la aplicación de domain driven design: en cada uno se discute qué conceptos son objetos valor, cuáles entidades y qué relaciones hay entre ellos, y se codifican en el lenguaje elegido siguiendo sus buenas prácticas. Cada commit resuelve un issue y lo referencia. Todavía no se incluye ninguna lógica de negocio: el objetivo es tener una representación correcta del dominio sobre la que trabajar en el siguiente milestone.

## Milestone 1: Lógica de negocio y comprobación automática

El objetivo de este hito es incorporar una primera pieza de comportamiento al sistema tomando como punto de partida la estructura del Milestone 0. El desarrollo se apoyará desde el comienzo en una batería de tests automatizados, que será la referencia para validar que el trabajo está realmente terminado.No se establece previamente cuál será la regla exacta ni cómo deberá resolverse técnicamente, ya que ambas cuestiones se concretarán mientras se implementa la funcionalidad. La aceptación del trabajo dependerá de que la implementación resultante disponga de sus pruebas correspondientes y de que estas finalicen correctamente. Por tanto, el entregable consiste en esa lógica inicial y en el conjunto de tests que demuestra su funcionamiento.

| Pregunta | Milestone 0 | Milestone 1 |
|---|---|---|
| **Se considera conseguido cuando...** | existe el modelo con su documentación de decisiones y se puede comprobar que se ha obtenido siguiendo la metodología de diseño de dominio, y no solo que el resultado parezca correcto | existe la lógica sobre el modelo, con sus pruebas automáticas funcionando, y se puede comprobar que se ha desarrollado siguiendo la metodología dirigida por pruebas |