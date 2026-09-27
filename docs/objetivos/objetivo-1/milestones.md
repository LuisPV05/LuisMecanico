# Milestones

Cada milestone es un entregable comprobable, no una lista de tareas ni qué HU se "cierra" con cada uno.


## Milestone 0: Modelo de reparaciones y recursos del taller

**¿Qué representa este hito?**
El conocimiento del taller convertido en modelo: tipos de reparación, recursos que exige cada una (herramientas, máquinas, elevadores), su duración habitual, y cómo se conecta todo con una cita.

**¿Por qué hace falta antes que nada?**
Porque no se puede detectar un conflicto de recursos si antes no existe una representación fiable de qué recurso ocupa cada reparación y durante cuánto tiempo. Este milestone es la base sobre la que se apoya todo lo demás.

**¿Cómo se sabe que está conseguido?**
Cuando, dada una reparación cualquiera del taller, el modelo permite responder sin ambigüedad qué recursos necesita y cuánto tiempo los va a ocupar.


## Milestone 1: Lógica de negocio y comprobación automática

**¿Qué representa este hito?**
La comprobación automática de disponibilidad: dado un recurso y una franja horaria, el sistema determina si está libre, y si no lo está, señala con qué cita choca. Incluye los tests que garantizan que esa comprobación es correcta.

**¿Por qué es el siguiente paso lógico?**
Porque ataca directamente el problema que motiva el proyecto: pasar de descubrir el conflicto cuando el mecánico ya está parado, a detectarlo en el momento de organizar la cita. Depende del modelo construido en el Milestone 0.

**¿Cómo se sabe que está conseguido?**
Cuando, al intentar registrar una cita que compite por un recurso ya ocupado, el sistema lo detecta y lo comunica antes de guardar la cita.
