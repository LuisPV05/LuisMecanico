# Milestones

Cada milestone es un entregable comprobable, no una lista de tareas ni qué HU se "cierra" con cada uno.

## Milestone 0: Modelo de reparaciones y recursos del taller

Es el conocimiento del taller convertido en modelo: los tipos de reparación que existen, los recursos que exige cada una —herramientas, máquinas, elevadores—, su duración habitual y cómo se conecta todo con una cita. Hace falta antes que nada porque no se puede detectar un conflicto de recursos si primero no existe una representación fiable de qué recurso ocupa cada reparación y durante cuánto tiempo; todo lo que viene después se apoya en esta base.


## Milestone 1: Lógica de negocio y comprobación automática

Es la comprobación automática de disponibilidad: dado un recurso y una franja horaria, el sistema determina si está libre y, si no lo está, señala con qué cita choca, acompañado de los tests que garantizan que esa comprobación es correcta. Es el siguiente paso lógico porque ataca directamente el problema que motiva el proyecto —pasar de descubrir el conflicto cuando el mecánico ya está parado, a detectarlo en el momento de organizar la cita— y no tendría sentido sin el modelo construido en el Milestone 0.

| Pregunta | Milestone 0 | Milestone 1 |
| **Se considera conseguido cuando...** | dada una reparación cualquiera del taller, el modelo permite responder sin ambigüedad qué recursos necesita y cuánto tiempo los va a ocupar | al intentar registrar una cita que compite por un recurso ya ocupado, el sistema lo detecta y lo comunica antes de guardar la cita |