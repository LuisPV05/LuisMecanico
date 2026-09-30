# Historias de Usuario

## Índice

| ID | Título |
|----|--------|
| HU001 | No sé si el recurso que necesita una reparación estará libre cuando apunto la cita |
| HU002 | No sé qué citas se van a ver afectadas cuando una reparación se retrasa |
| HU003 | No sé si mi cita se va a cumplir hasta que llego al taller |
| HU004 | No puedo consultar las citas y los recursos desde donde estoy trabajando |
| HU005 | No sé si podré usar el recurso que necesito al terminar mi tarea actual |

## HU001
| Campo | Descripción |
|----|--------|
| Persona | Recepcionista |
| Situación | Un cliente llama para pedir cita y hay que apuntarla en el calendario. |
| Información que interviene | xperiencia directa como recepcionista del taller, recogida en el [README](../README.md). Datos que intervienen: tipo de reparación indicado a partir del síntoma del cliente, recurso o recursos que exige esa reparación, franja horaria solicitada y ocupación actual de esos recursos en ese tramo. |
| Obstáculo | No hay forma de saber si la herramienta, máquina o elevador que necesitará la reparación estará libre a esa hora. |
| Consecuencia | El conflicto no aparece hasta que la reparación ya ha empezado, cuando el mecánico llega y se encuentra el recurso ocupado por otra cosa. |

## HU002
| Campo | Descripción |
|----|--------|
| Persona | Recepcionista |
| Situación | Una reparación en curso se alarga y sigue ocupando un recurso compartido. |
| Información que interviene | duración estimada frente a duración real de la reparación en curso, recurso concreto que sigue ocupado, y qué citas posteriores tienen asignado ese mismo recurso. Se considera que una cita está "afectada" cuando su hora de inicio prevista es posterior al momento en que el recurso queda libre según la duración real de la reparación en curso, pero anterior o igual a la hora en que el recurso vuelve a quedar libre. |
| Obstáculo | No hay forma de saber qué citas posteriores dependen de ese mismo recurso. |
| Consecuencia | Si nadie se acuerda de avisar, el cliente afectado no se entera del retraso hasta que ya está en camino o incluso hasta que llega al taller. |

## HU003
| Campo | Descripción |
|----|--------|
| Persona | Cliente |
| Situación | Se pide cita por teléfono para una reparación. |
| Información que interviene | hora acordada para la cita, estado de disponibilidad de los recursos asociados a esa cita en el momento de confirmarla, y forma de contacto del cliente para un posible aviso. |
| Obstáculo | No hay ninguna garantía de que la reparación vaya a empezar a la hora acordada, porque el propio taller tampoco puede saberlo con certeza de antemano. |
| Consecuencia | Muchas veces el cliente se presenta con el coche y ahí se entera de que hay que esperar, sin que nadie le haya avisado antes. |

## HU004
| Campo | Descripción |
|----|--------|
| Persona | Recepcionista y mecánico |
| Situación | La información de las citas está en un calendario físico que permanece en recepción. |
| Información que interviene | el conjunto de citas y recursos del taller como información compartida, y desde qué puesto (recepción o zona de trabajo) se consulta esa información en cada momento. |
| Obstáculo | Solo se puede consultar desde donde está el calendario, así que el mecánico no puede ver las citas ni el estado de los recursos sin preguntar antes en recepción. |
| Consecuencia | Cada persona trabaja con una información distinta y desactualizada según dónde se encuentre en ese momento. |

## HU005
| Campo | Descripción |
|----|--------|
| Persona | Mecánico |
| Situación | Está a punto de acabar con un vehículo y tiene que pasar al siguiente trabajo. |
| Información que interviene | recurso necesario para el siguiente trabajo asignado, estado de ese recurso en el momento en que el mecánico lo va a necesitar, y orden de los trabajos pendientes de ese mecánico.s |
| Obstáculo | No sabe de antemano si el recurso que necesitará para el próximo coche estará disponible. |
| Consecuencia | Se ve obligado a  esperar y no puede organizar en qué orden abordará los trabajos que tiene pendientes. |