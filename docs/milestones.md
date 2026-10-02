# Milestones

## Milestone 0: Modelo de reparaciones y recursos del taller

El modelo del dominio del taller se obtiene aplicando una metodología de diseño de dominio a partir de las [historias de usuario](historias-usuarios.md), las [jornadas de usuario](user-journeys.md) y lo descrito en el [objetivo 0](objetivo-0.md). Se considera que se ha seguido ese proceso cuando, junto al modelo, existe un documento de decisiones de diseño que cualquier otra persona pueda revisar y contrastar contra las historias de usuario de las que parte, sin necesidad de juzgar subjetivamente si el código "está bien". A partir de ahí se entrega el modelo en sí, sin ninguna lógica de negocio todavía, junto con ese documento de decisiones.

HU Relacionada: HU001.

## Milestone 1: Lógica de negocio y comprobación automática

La primera lógica de negocio se desarrolla sobre el modelo del Milestone 0 siguiendo una metodología dirigida por pruebas, con la infraestructura de comprobación automática funcionando desde el principio. Se considera que el proceso se ha seguido cuando existen las pruebas automáticas correspondientes, estas se ejecutan y pasan correctamente, en lugar de evaluar si la lógica implementada «parece» correcta a simple vista. La regla concreta que se resuelve y la forma de implementarla se determinan durante el propio desarrollo; lo que se entrega es esa lógica mínima junto con sus respectivas pruebas.

HU Relaconada: HU002.

| Pregunta | Milestone 0 | Milestone 1 |
|---|---|---|
| **Se considera conseguido cuando...** | existe el modelo con su documentación de decisiones y se puede comprobar que se ha obtenido siguiendo la metodología de diseño de dominio, y no solo que el resultado parezca correcto | existe la lógica sobre el modelo, con sus pruebas automáticas funcionando, y se puede comprobar que se ha desarrollado siguiendo la metodología dirigida por pruebas |