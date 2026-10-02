# Historias de Usuario

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
| Persona | Recepcionista y mecánico |
| Situación | Una reparación en curso se alarga y sigue ocupando un recurso compartido que otra cita o otro mecánico necesita a continuación. |
| Información que interviene | duración estimada frente a duración real de la reparación en curso, recurso concreto que sigue ocupado, y qué citas o trabajos posteriores dependen de ese mismo recurso. Se considera que una cita o tarea está "afectada" cuando su inicio previsto cae entre el momento en que el recurso debería quedar libre según la duración estimada y el momento en que realmente queda libre. |
| Obstáculo | Ni el recepcionista sabe qué citas se verán afectadas, ni el mecánico que espera ese recurso sabe cuándo va a estar libre. |
| Consecuencia | El aviso al cliente afectado depende de que alguien se acuerde de hacerlo a tiempo, y el mecánico se queda esperando sin poder planificar el orden de sus trabajos pendientes. |


## HU003
| Campo | Descripción |
|----|--------|
| Persona | Recepcionista y mecánico |
| Situación | La información de las citas y los recursos está en un calendario físico que permanece en recepción. |
| Información que interviene | el conjunto de citas y recursos del taller como información compartida, y desde qué puesto (recepción o zona de trabajo) se consulta esa información en cada momento. |
| Obstáculo | Solo puede consultarse desde donde está el calendario, así que el mecánico no puede ver las citas ni el estado de los recursos sin preguntar antes en recepción. |
| Consecuencia | Cada persona trabaja con una información distinta y desactualizada según dónde se encuentre en ese momento. |