# Jornadas de usuario

Las siguientes jornadas describen el uso completo del sistema una vez desarrollado. Alguna funcionalidad (como el aviso automático de retrasos en Jornada 3) corresponde a historias de usuario que se implementarán en milestones futuros.

## Jornada 1: Luis registra una cita sin conflictos

Kips llama porque su coche pierde presión en los frenos. Luis apunta el vehículo y el síntoma, y el sistema identifica que se trata probablemente de un cambio de pastillas y discos de freno, indicando que se necesitará un elevador durante 1 hora y 30 minutos. Luis consulta la disponibilidad de los elevadores para las 10:00, que es cuando el cliente puede acercarse. El sistema confirma que hay uno de los tres elevadores libre en ese tramo, así que Luis registra la cita y le da al cliente una hora aproximada de finalización.

## Jornada 2: Luis detecta un conflicto antes de confirmar la cita

Poco después, otro cliente llama pidiendo cita también a las 10:00 para un cambio de amortiguadores (2 horas y 30 minutos, también con elevador). Al intentar registrar la cita, el sistema avisa a Luis de que a esa hora los tres elevadores ya estarán ocupados por reparaciones en curso. Luis consulta las franjas libres que le propone el sistema y ofrece al cliente el primer horario disponible, evitando así que el cliente llegue con el coche y tenga que esperar sin saberlo de antemano.

## Jornada 3: Snoopy sufre un retraso que afecta a la siguiente cita

Snoopy, el mecánico está haciendo un cambio de kit de distribución, que ocupa un elevador durante 4 horas, y la reparación se alarga más de lo previsto. El sistema detecta que ese elevador seguirá bloqueado cuando empiece la siguiente cita programada sobre el mismo recurso, y avisa a Luis del posible retraso para que pueda llamar al cliente afectado antes de que llegue al taller, tal y como se hace actualmente pero sin depender de que Luis se dé cuenta a tiempo por sí mismo.