# Historias de Usuario
Cada historia de usuario lleva su identificador [HUxxx] y la etiqueta `user-stories` en el issue correspondiente de GitHub.

## Índice

| ID | Título |
|----|--------|
| HU001 | No sé si el recurso que necesita una reparación estará libre cuando apunto la cita |
| HU002 | No sé qué citas se van a ver afectadas cuando una reparación se retrasa |
| HU003 | No sé si mi cita se va a cumplir hasta que llego al taller |
| HU004 | No sé si podré usar el recurso que necesito al terminar mi tarea actual |
| HU005 | No puedo consultar las citas y los recursos desde donde estoy trabajando |

## HU001
- **Quién:** Recepcionista
- **Situación:** Un cliente llama para pedir cita y hay que apuntarla en el calendario.
- **Obstáculo:** No hay forma de saber si la herramienta, máquina o elevador que necesitará la reparación estará libre a esa hora.
- **Consecuencia:** El conflicto no aparece hasta que la reparación ya ha empezado, cuando el mecánico llega y se encuentra el recurso ocupado por otra cosa.
- **Datos:** experiencia directa como recepcionista del taller, recogida en el [README](../../../README.md). Datos que intervienen: tipo de reparación indicado a partir del síntoma del cliente, recurso o recursos que exige esa reparación, franja horaria solicitada y ocupación actual de esos recursos en ese tramo.

## HU002
- **Quién:** Recepcionista
- **Situación:** Una reparación en curso se alarga y sigue ocupando un recurso compartido.
- **Obstáculo:** No hay forma de saber qué citas posteriores dependen de ese mismo recurso.
- **Consecuencia:** Si nadie se acuerda de avisar, el cliente afectado no se entera del retraso hasta que ya está en camino o incluso hasta que llega al taller.
- **Datos:** duración estimada frente a duración real de la reparación en curso, recurso concreto que sigue ocupado, y qué citas posteriores tienen asignado ese mismo recurso. Se considera que una cita está "afectada" cuando su hora de inicio prevista es posterior al momento en que el recurso queda libre según la duración real de la reparación en curso, pero anterior o igual a la hora en que el recurso vuelve a quedar libre.

## HU003
- **Quién:** Cliente
- **Situación:** Se pide cita por teléfono para una reparación.
- **Obstáculo:** No hay ninguna garantía de que la reparación vaya a empezar a la hora acordada, porque el propio taller tampoco puede saberlo con certeza de antemano.
- **Consecuencia:** Muchas veces el cliente se presenta con el coche y ahí se entera de que hay que esperar, sin que nadie le haya avisado antes.
- **Datos:** hora acordada para la cita, estado de disponibilidad de los recursos asociados a esa cita en el momento de confirmarla, y forma de contacto del cliente para un posible aviso.

## HU004
- **Quién:** Mecánico
- **Situación:** Está terminando  un vehículo y le toca pasar al siguiente.
- **Obstáculo:** No tiene forma de saber si el recurso que necesita para el siguiente coche va a estar libre.
- **Consecuencia:** Se queda esperando y no puede organizar en qué orden atacar los trabajos que tiene pendientes.
- **Datos:** recurso necesario para el siguiente trabajo asignado, estado de ese recurso en el momento en que el mecánico lo va a necesitar, y orden de los trabajos pendientes de ese mecánico.

## HU005
- **Quién:** Recepcionista y mecánico
- **Situación:** La información de las citas está en un calendario físico que permanece en recepción.
- **Obstáculo:** Solo se puede consultar desde donde está el calendario, así que el mecánico no puede ver las citas ni el estado de los recursos sin preguntar antes en recepción.
- **Consecuencia:** Cada persona trabaja trabajando con una información distinta y desactualizada según dónde se encuentre en ese momento.
- **Datos:** el conjunto de citas y recursos del taller como información compartida, y desde qué puesto (recepción o zona de trabajo) se consulta esa información en cada momento.