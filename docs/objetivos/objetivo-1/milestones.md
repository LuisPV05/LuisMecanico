# Milestones

Determinan el nivel de generalización o detalle del problema se va a mostrar al cliente en cada entrega, no qué características técnicas va a tener el producto ni qué HU se "cierra" con cada uno.


## Milestone 0: Modelo de reparaciones y recursos del taller

### Qué se entrega

Un modelo mínimo pero fiel del dominio del taller: qué tipos de reparación existen, qué recursos necesita cada una, cómo son las citas y cómo se relacionan entre sí. Este milestone corresponde al objetivo de entender y modelizar el problema antes de automatizar nada.

### Cuál es el objetivo del milestone

Asegurar que cualquier decisión de diseño posterior se apoya en una representación correcta de cómo funciona realmente el taller, en lugar de en supuestos. Es la base necesaria para poder empezar a razonar sobre conflictos de recursos en el siguiente milestone.

## Milestone 1: Lógica de negocio y comprobación automática

### Qué se entrega

La infraestructura que permite comprobar, a partir del modelo del milestone anterior, si un conjunto de recursos está disponible durante una franja horaria determinada, junto con los tests automáticos que validan ese comportamiento. Cómo se le presenta esto finalmente al recepcionista (interfaz, formato, canal) es una decisión que se toma durante el propio desarrollo, no en este documento.

### Cuál es el objetivo del milestone

Empezar a resolver el problema central del taller —detectar conflictos de recursos antes de que ocurran, en lugar de cuando ya es tarde— apoyándose en el modelo construido en el Milestone 0.