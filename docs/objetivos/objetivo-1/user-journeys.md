# Jornadas de usuario

Recorridos completos de uso una vez desarrollado el sistema.


## Jornada 1: Luis cierra una cita por teléfono sin colgar dos veces

1. Kips llama describiendo un ruido al frenar.
2. Luis introduce el síntoma en el sistema.
3. El sistema deduce que la reparación requiere un elevador durante 1h30.
4. Kips le pide que la cita sea a las 10:00, Luis registra la cita para las 10:00; el sistema comprueba los tres elevadores y ve que están ocupados.
5. El sistema devuelve el primer hueco libre (11:30) y Luis se lo ofrece a Kips en la misma llamada.


## Jornada 2: Luis encaja un imprevisto sin repasar toda la agenda

1. Un cliente se presenta sin cita para una reparación urgente.
2. Luis necesita saber si el gato hidráulico está libre en la próxima hora.
3. Consulta la disponibilidad de ese recurso concreto.
4. El sistema muestra un hueco libre de una hora.
5. Luis acepta el trabajo sin descuadrar el resto del día.


## Jornada 3: Snoopy sufre un retraso que afecta a la siguiente cita, detectanado el efecto dominó de un retraso 

1. Snoopy lleva 45 minutos de retraso sobre lo previsto en una reparación.
2. El sistema compara la duración real con la estimada para ese recurso.
3. Identifica que la siguiente cita en el mismo elevador se verá afectada.
4. Genera un aviso para que recepción llame al cliente antes de que llegue.