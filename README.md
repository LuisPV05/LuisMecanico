# LuisMecánico

## Administrador

Soy el recepcionista de un taller mecánico.

## Problema del administrador

En el taller mecánico he observado problemas para organizar las citas de los clientes durante el día. Actualmente, cuando un cliente llama para pedir una cita, se apunta en un calendario, pero no se tiene en cuenta si las herramientas o máquinas necesarias para realizar la reparación están disponibles en ese momento.

Este conflicto provoca que cuando dos o más reparaciones que coinciden en el tiempo necesitan utilizar la misma herramienta o máquina se pierda tiempo en el taller. En ese caso, aunque las citas estén correctamente registradas en el calendario, los trabajos no pueden realizarse de forma independiente porque existe un recurso que debe ser compartido.

En los casos en los que hay retrasos en la reparación o revisión del vehículo por tema de que la herramienta necesaria este en uso, se llama al cliente y se le avisa del retraso inmediatamente. Las herramientas estan disponibles para uso si no están siendo usadas por ningún mecánico, si están en uso se tiene que esperar la finalización de su uso para usarlas.

Esto genera perdida de tiempo en la eficiencia de la reparación y que el cliente decida si puede dejar el coche para ser revisado aún llevando retraso o cancelar la cita, coger su vehículo e irse, lo que hace perder dinero al taller.

Por tanto, el problema no es simplemente organizar las citas, sino detectar y evitar los conflictos entre citas que coinciden en el tiempo y necesitan los mismos recursos.

## De dónde viene mi conocimiento y relación sobre el problema

Mi experiencia al problema surge de una situación propia. En los ultimos tres años, durante vacaciones y fines de semana he trabajado como recepcionista del taller mecanico de mi tío y vi de primera mano cómo se organizaban las citas y los problemas que aparecian durante la jornada.

Durante este tiempo, las citas se suelen gestionar mediante llamadas telefónicas y se van apuntando en un calendario. En este se apunta el tipo de vehículo que es, el problema que tiene con el vehículo, sea cambio de ruedas, fallos en la transmisión o lo que el cliente note que no funciona bien en si vehículo, y su respectiva sensaciones y efectos que sienten con este fallo que permita reconocer el error que sucede para saber cuales seran las herramientas/máquinas necesarias para revisar y reparar el vehículo. 

Además, mi conocimiento sobre el tema de reparaciones viene de mi tío y por parte mía, del interés que he tenido desde chico en los coches, habiendo aprendido poco a poco el tema de que herramientas necesitan ciertas reparaciones, entonces cuando llamaba un cliente sabía las herramientas necesarias para la reparación.

Por tanto, el problema esta planteado por una situacion vivida de primera mano. Parte de esta de experiencia personal trabajando en la recepcion del taller y de la información obtenida por parte de mi tío.


## Casos dados desde la experiencia

Tengo dos clientes con cita a las 10:00. Una de las reparaciones necesita utilizar un elevador y la otra también. Si ambos trabajos comienzan a la misma hora, uno de los vehículos tendrá que esperar a que el elevador quede libre.

En una jornada con carga alta, se pueden producir aproximadamente varios conflictos de este tipo. La espera de una cita afectada puede estar habitualmente entre 20 y 60 minutos, dependiendo de la reparación y de los recursos que necesite


## Por qué el enfoque actual resulta insuficiente

- El calendario permite saber cuándo hay una cita y de que trata, pero no relaciona directamente la cita con los recursos necesarios para realizar la reparación.
- No se comprueba automáticamente si las herramientas o máquinas necesarias están disponibles durante toda la reparación.
- Dos citas pueden coincidir en el tiempo aunque necesiten utilizar el mismo recurso.
- Los conflictos se detectan cuando llega el momento de realizar la reparación, en lugar de detectarse al organizar la cita.
- Cuando aparece un conflicto, el encargado/recepcionista tiene que reorganizar manualmente la planificación y avisar al cliente si se produce un retraso.

## Afectados por este problema

- **Mecánicos**: pueden encontrarse con que la herramienta necesaria este siendo usada en otra reparación, lo que provoca retrasos. 
- **Clientes**: tienen la posiblidad de sufrir retrasos en la reparación de su vehículo por la falta de recursos para la reparación.
- **El taller**: pierde tiempo con las esperas, provocando cambios en la planificación o retrasos.
- **Encargado/recepcionista del taller**: tiene que organizar las citas teniendo en cuenta no solamente los horarios, sino también las máquinas y herramientas disponibles, coordinando las citas y recursos disponibles.

## Datos que intervienen en el problema y datos que no se tienen en cuenta al apuntar la cita

Actualmente, las citas se gestionan mediante llamadas telefónicas y se registran en un calendario.

Según la experiencia en el taller, se registra la información de la cita de:

- Fecha y hora prevista de la cita.
- Duración estimada de la reparación.
- Tipo de reparación.
- Problema o síntomas indicados por el cliente.
- Herramientas o máquinas necesarias.
- Mecánico asignado.
- Retrasos producidos y su duración, cuando los haya.

A parte de esta información del taller, a la hora de registrar los datos que una cita no se apuntan:
- Disponibilidad de cada herramienta o máquina.
- Número de unidades disponibles de cada recurso ( Si de esa herramienta/máquina hay más de una unidad )
- Periodos en los que cada recurso está ocupado.

En una jornada habitual se gestionan alrededor de 12 citas. Las reparaciones pueden durar desde unos 45 minutos hasta varias horas. Por ejemplo, una diagnosis electrónica puede tardar aproximadamente 60 minutos, mientras que un cambio de amortiguadores puede llegar a durar unas 2 horas.


## Cosas a tener en cuenta y mejorar para la resolución del problema

Las citas, los mecánicos y los recursos del taller deben poder consultarse desde el lugar donde se organiza la jornada. La información sobre las citas y la disponibilidad de las máquinas debe estar centralizada para que se pueda comprobar en tiempo real si una nueva cita entra en conflicto con las que ya están programadas.

Por este motivo, el método actual no permite mantener la información accesible y actualizada para las personas que necesiten gestionar las citas y los recursos del taller.

Ante una nueva cita, se debería comprobar:

- Identificar el tipo de reparación a partir del problema indicado por el cliente.
- Determinar qué herramientas y máquinas son necesarias para realizar la reparación.
- Comprobar la disponibilidad de esos recursos en el horario solicitado.
- Comparar la nueva cita con las citas ya programadas.
- Detectar si existe algún conflicto por utilizar el mismo recurso durante un periodo coincidente.
- Calcular la duración prevista de la reparación y durante cuánto tiempo estarán ocupados los recursos.
- Avisar si la cita genera un conflicto y permitir buscar otro horario disponible.

El valor de la solución está en detectar el conflicto cuando se organiza la cita, en vez de descubrirlo cuando empieza la reparación.

## Objetivo que se aborda en el problema

El objetivo es ver como se puede mejorar la organización de citas teniendo en cuenta los horarios, las citas previstas y la disponibilidad de herramientas.

Con esto se pretende no tener conflictos y evitar tiempos de espera innecesarios y disminuir los retrasos que afectan a las reparaciones y los clientes.

Para la solucíon, esta en comprender el foco del problema y definir correctamente los datos que tiene el problema para poder solucionarlo de manera correcta.

## Documentación adicional y configuración

La documentación relativa al objetivo 0 se encuentra [aquí](doc/objetivos/objetivo-0.md).
