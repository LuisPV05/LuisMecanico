# LuisMecánico

## Problema

En el taller mecánico he observado problemas para organizar las citas de los clientes durante el día. Cuando se recibe la llamada de un cliente para pedir una cita, se apunta en un calendario, pero no se tiene en cuenta si las herramientas necesarias para realizar la reparación están disponibles en ese momento.

Provoca que cuando dos o más reparaciones que coinciden en el tiempo necesitan utilizar la misma herramienta o máquina. En ese caso, aunque las citas estén correctamente registradas en el calendario, los trabajos no pueden realizarse de forma independiente porque existe un recurso que debe ser compartido.

En los casos en los que hay retrasos en la reparación o revisión del vehículo por tema de que la herramienta necesaria este en uso, se llama al cliente y se le avisa del retraso inmediatamente. Las herramientas estan disponibles para uso si no están siendo usadas por ningún mecánico, si están en uso se tiene que esperar la finalización de su uso para usarlas.

Esto genera perdida de tiempo en la eficiencia de la reparación y que el cliente decida si puede dejar el coche para ser revisado aún llevando retraso o cancelar la cita, coger su vehículo e irse, lo que hace perder dinero al taller.

Por tanto, el problema no es simplemente organizar las citas, sino detectar y evitar los conflictos entre citas que coinciden en el tiempo y necesitan los mismos recursos.

## De dónde viene mi conocimiento sobre el problema

Mi experiencia al problema surge de una situación propia. En los ultimos tres años, durante vacaciones y fines de semana he trabajado como recepcionista del taller mecanico de mi tío y vi de primera mano cómo se organizaban las citas y los problemas que aparecian durante la jornada.

Durante este tiempo, las citas se suelen gestionar mediante llamadas telefónicas y se van apuntando en un calendario. En este se apunta el tipo de vehículo que es, el problema que tiene con el vehículo, sea cambio de ruedas, fallos en la transmisión o lo que el cliente note que no funciona bien en si vehículo, y su respectiva sensaciones y efectos que sienten con este fallo que permita reconocer el error que sucede para saber cuales seran las herramientas/máquinas necesarias para revisar y reparar el vehículo. 

Además, mi conocimiento sobre el tema de reparaciones viene de mi tío y por parte mía, del interés que he tenido desde chico en los coches, habiendo aprendido poco a poco el tema de que herramientas necesitan ciertas reparaciones, entonces cuando llamaba un cliente sabía las herramientas necesarias para la reparación.

Por tanto, el problema esta planteado por una situacion vivida de primera mano. Parte de esta de experiencia personal trabajando en la recepcion del taller y de la información obtenida por parte de mi tío.


## Cómo se organizan hoy las citas, y por qué no basta

Actualmente , el cliente llama para solicitar una cita, se recogen los datos del vehículo y el problema que tiene y se acuerda una fecha y una hora. Esa cita se registra en un calendario para saber qué trabajos están previstos durante la jornada.

El problema es que el calendario permite saber cuándo hay cita, pero no los datos necesarios para saber que se tiene que reparar.

Por ejemplo, si dos clientes tienen cita a las 10:00 y las dos reparaciones necesitan utilizar lo mismo, ambas citas se pueden dar a la misma hora. A la hora de la reparación, una de estas se tendra que esperar a que la otra termine.

Para entender y medir este problema es necesario relacionar, para cada cita, al menos la fecha y hora, la duración estimada de la reparación, el tipo de reparación y los recursos necesarios. También es necesario conocer para cada recurso cuántas unidades hay disponibles y durante qué periodos están ocupadas.

## Afectados por este problema

- **Mecánicos**: pueden encontrarse con que la herramienta necesaria este siendo usada en otra reparación, lo que provoca retrasos. 
- **Clientes**: tienen la posiblidad de sufrir retrasos en la reparación de su vehículo por la falta de recursos para la reparación.
- **El taller**: pierde tiempo con las esperas, provocando cambios en la planificación o retrasos.
- **Encargado del taller**: tiene que organizar las citas teniendo en cuenta no solamente los horarios, sino también las máquinas y herramientas disponibles, coordinando las citas y recursos disponibles.

## Datos que intervienen en el problema

En el taller existen datos relacionados con las citas y con los recursos utilizados en las reparaciones. Para poder estudiar el problema de forma cuantitativa, estos datos deben identificarse y obtenerse del funcionamiento real del taller.

Por un lado, de las citas interesa conocer datos como la fecha, la hora prevista de inicio, la duración estimada, el tipo de reparación y los recursos necesarios. Un ejemplo de registro podría tener la forma: fecha, hora, duración, reparación, herramientas o máquinas necesarias y mecánico asignado. También es importante registrar, cuando existan, los retrasos producidos y su duración.

Por otro lado, de las herramientas y máquinas interesa conocer qué recursos existen, cuántas unidades hay de cada uno y qué reparaciones necesitan cada recurso. También hay que confirmar si un mismo recurso puede utilizarse simultáneamente o si cada unidad solo puede atender una reparación cada vez.

Además, durante la llamada el cliente dice el fallo o error de su vehículo, que determina el tipo de reparación que se hara y las herramientas que se utilizarán.

Para comprender mejor el problema que ocurreen el taller una experiencia propia que vi en primizia, el taller diariamente gestiona alrededor de 12 citas en una jornada habitual. Las reparaciones pueden durar desde unos 45 minutos hasta varias horas, dependiendo del trabajo que tenga que realizarse. Por ejemplo, si es hacer una diagnosis electrónica del vehículo se tarda unos 60 minutos, mientras que un cambio de amortiguadores llega a tardar 2 horas.

Entre los recursos compartidos durantes las reparaciones se consideran, por ejemplo, dos elevadores, una máquina de diagnosis y una desmontadora/equilibradora, de forma que una reparación puede impedir que otra utilice el mismo recurso durante parte de su duración.

Cuando dos reparaciones requieren el mismo recurso en un periodo coincidente, la situación se resuelve de la siguiente manera: el encargado mantiene una de las reparaciones en espera y el mecánico continúa con el trabajo que ya está utilizando el recurso. Cuando termina, el recurso pasa a la siguiente reparación; si la espera afecta demasiado al resto de la jornada, también se puede mover la cita a otro momento, avisando al dueño del vehículo del retraso en la reparación. En este escenario, el problema se puede medir mediante el número de conflictos entre citas, el número de citas afectadas y el tiempo de espera producido hasta que el recurso queda disponible. Como referencia, consideramos que en una jornada con carga alta, como la de ese día, se produjo aproximadamente tres conflictos de este tipo y que la espera de la cita afectada puede encontrarse habitualmente entre 20 y 60 minutos, dependiendo de la reparacion y de los recursos que necesite.

El problema viene dado de la gestion separada de esto. No hay una relación entre la cita, el trabajo que se hará y las herramientas necesarias durante la reparación

## Cosas a tener en cuenta

Al recibir una llamada, se tiene en cuanta al mismo tiempo:

1.El horario solicitado.
2.El problema presentado.
3.El tipo de reparación que puede ser realizado.
4.Las herramientas que se usarán.
5.Si las herramientas estan disponibles en ese horario.
6.Si otra reparación va a usar esos recursos durante ese periodo
7.La duración prevista de la reparación y, por tanto, durante cuánto tiempo estarán ocupados esos recursos.

La información exacta de cada uno de estos puntos debe comprobarse con el personal del taller. Esto permitirá conocer cómo se toman las decisiones actualmente y qué parte del proceso provoca los conflictos.

Cuando no se tiene en cuenta la disponibilidad de los recursos durante todo el periodo de la reparación, pueden producirse los problemas ya comentados. 

## Necesidades detectadas desde la experiencia

A partir de mi experiencia, el cambio de como se organizan las citas es necesario no solo viendo el horario, sino los recursos disponibles que hay para esa reparación al la hora de la cita.

La necesidad principal es saber a tiempo si la cita crea conflicto a otra por el uso de ambas de una herramienta necesaria al mismto tiempo.

De esta manera, el problema se detecta al momento de organizar la cita, en vez de verlo al momento de la reparación.

## Objetivo que se aborda en el problema

El objetivo es ver como se puede mejorar la organización de citas teniendo en cuenta los horarios, las citas previstas y la disponibilidad de herramientas.

Con esto se pretende no tener conflictos y evitar tiempos de espera innecesarios y disminuir los retrasos que afectan a las reparaciones y los clientes.

Para la solucíon, esta en comprender el foco del problema y definir correctamente los datos que tiene el problema para poder solucionarlo de manera correcta.

## Documentación adicional y configuración

La documentación relativa al objetivo 0 se encuentra [aquí](doc/objetivos/objetivo-0.md).
