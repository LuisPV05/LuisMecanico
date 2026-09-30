# Milestones
Cada milestone es un entregable comprobable, no una lista de tareas ni qué HU se "cierra" con cada uno.

## Milestone 0: Modelo de reparaciones y recursos del taller

Se entrega el modelo del dominio del taller, junto con un breve documento que recoja las decisiones tomadas al modelarlo, sin ninguna lógica de negocio todavía. Ese modelo se obtiene aplicando una metodología de diseño de dominio a partir de las [historias de usuario](historias-usuarios.md), las [jornadas de usuario](user-journeys.md) y lo descrito en el [objetivo 0](objetivo-0.md), y solo se puede dar por válido si se ha seguido ese proceso. Hace falta antes que nada porque todo lo que viene después se apoya en él.

Se considera que se ha seguido el proceso cuando existe, junto al modelo, un documento de decisiones de diseño que cualquier otra persona pueda revisar y contrastar contra las historias de usuario de las que parte, sin necesidad de juzgar subjetivamente si el código "está bien".

HU Relacionada: HU001.

## Milestone 1: Lógica de negocio y comprobación automática

Se entrega la primera lógica de negocio sobre el modelo del Milestone 0, acompañada de pruebas automáticas que comprueban su comportamiento. Se desarrolla con una metodología dirigida por pruebas, con la infraestructura de comprobación automática funcionando desde el principio. Qué regla concreta se resuelve y cómo se implementa se decide durante el propio desarrollo. Es el siguiente paso lógico porque empieza a atacar el problema que motiva el proyecto apoyándose en la estructura ya construida.

Se considera que se ha seguido el proceso cuando existen las pruebas automáticas correspondientes y estas se ejecutan y pasan, en lugar de valorar si la lógica implementada "parece" correcta a simple vista.

HU Relaconada: HU002.

| Pregunta | Milestone 0 | Milestone 1 |
|---|---|---|
| **Se considera conseguido cuando...** | existe el modelo con su documentación de decisiones y se puede comprobar que se ha obtenido siguiendo la metodología de diseño de dominio, y no solo que el resultado parezca correcto | existe la lógica sobre el modelo, con sus pruebas automáticas funcionando, y se puede comprobar que se ha desarrollado siguiendo la metodología dirigida por pruebas |