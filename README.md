# LuisMecánico

## Problema

Un taller mecánico tiene problemas para organizar las citas de los clientes durante el día. Cuando un cliente llama para pedir una cita, se apunta en un calendario, pero no siempre se comprueba si las máquinas y herramientas necesarias para realizar la reparación están disponibles en ese momento.

Esto puede provocar que dos mecánicos necesiten utilizar la misma máquina o el mismo elevador a la misma hora, generando esperas, cambios de horario y retrasos en las reparaciones.

## De dónde viene esto

Este problema surge de la organización habitual de las citas en el taller. Actualmente, las citas se realizan mediante llamadas telefónicas y se van apuntando en un calendario.

Durante la llamada, el cliente explica cuál es el problema que tiene su vehículo y se intenta asignar una fecha y una hora para llevarlo al taller. Sin embargo, al organizar la cita no siempre se tiene en cuenta qué máquinas, herramientas o elevadores serán necesarios para realizar el trabajo.

Por tanto, aunque exista una hora libre en el calendario, puede que los recursos necesarios para esa reparación estén ocupados por otro mecánico.

## Cómo se organizan hoy las citas, y por qué no basta

Actualmente las citas se gestionan mediante llamadas telefónicas y se apuntan en un calendario.

El problema es que el calendario únicamente permite saber cuándo hay una cita, pero no proporciona una visión completa de los recursos que serán necesarios para realizar cada trabajo.

Por ejemplo, si dos clientes tienen cita a las 10:00 y las dos reparaciones necesitan utilizar el mismo elevador, ambas citas pueden quedar registradas correctamente en el calendario, pero cuando llegue el momento de realizar las reparaciones uno de los mecánicos tendrá que esperar.

Además, durante la llamada el cliente puede explicar el problema que tiene el vehículo, pero es necesario relacionar ese problema con las herramientas y máquinas que pueden ser necesarias para solucionarlo.

## Quienes son los afectados por este problema

- **Mecánicos**: pueden encontrarse con que una herramienta, máquina o elevador que necesitan está siendo utilizado por otro compañero. Esto provoca tiempos de espera y dificulta la organización del trabajo.

- **Clientes**: pueden sufrir retrasos en la reparación de su vehículo porque los recursos necesarios no estaban disponibles cuando se había planificado la cita.

- **El taller**: pierde tiempo debido a una mala coordinación de los recursos y puede acumular retrasos durante el día.

- **Encargado del taller**: tiene que organizar las citas teniendo en cuenta no solamente los horarios, sino también las máquinas y herramientas disponibles.

## Qué datos tenemos

El taller ya dispone de información que puede utilizarse para solucionar el problema.

Por un lado, existe un listado de las máquinas y herramientas disponibles en el taller. Por otro, se dispone de las citas de los clientes, que actualmente se registran en un calendario.

Además, durante la llamada el cliente proporciona información sobre el problema que presenta su vehículo. Esta información puede utilizarse para determinar qué recursos pueden ser necesarios para realizar la reparación.

El problema no es tanto la falta de datos, sino que actualmente estos datos no están relacionados entre sí de una forma que permita comprobar automáticamente la disponibilidad de los recursos.

## Quedan tareas pendientes de procesar.

Actualmente, cuando se recibe una llamada, habría que comprobar manualmente qué problema tiene el vehículo, qué herramientas o máquinas pueden ser necesarias y si estarán disponibles en el horario solicitado.

La aplicación debería centralizar esta información y relacionar las citas con los recursos necesarios para cada reparación.

Por ejemplo, si un cliente solicita una cita para una reparación que necesita un elevador concreto, el sistema debería comprobar si ese elevador estará disponible durante el horario elegido.

Si otro mecánico ya tiene asignada una reparación que utiliza ese mismo recurso durante ese periodo, el sistema debería detectar el conflicto y evitar que ambas citas se programen de forma incompatible.

De esta manera, el sistema podría ayudar al taller a organizar las citas teniendo en cuenta tanto el horario como las máquinas y herramientas disponibles, reduciendo los retrasos y los conflictos entre mecánicos.

También sería posible tener en cuenta que un cliente no acuda finalmente a su cita, de forma que el recurso reservado pueda quedar disponible para otra reparación.

## Por qué es necesaria la app en la nube

La aplicación necesita estar disponible desde el taller para que la información sobre las citas, máquinas, herramientas y recursos utilizados esté centralizada.

Al estar en la nube, los datos pueden mantenerse en un único lugar y ser consultados y modificados por las personas encargadas de organizar el trabajo del taller.

Esto permite que cuando se registre una nueva cita, el sistema pueda consultar inmediatamente las reservas existentes y comprobar si los recursos necesarios están disponibles.

Además, cualquier cambio realizado en una cita puede quedar reflejado para el resto de usuarios que estén utilizando la aplicación.

## Qué debería tener la aplicación para resolver este problema

La aplicación debería permitir al taller:

- Registrar las citas de los clientes.
- Guardar información sobre el problema del vehículo.
- Asociar una reparación con las máquinas y herramientas que necesita.
- Consultar la disponibilidad de los recursos del taller.
- Detectar conflictos entre citas que necesitan los mismos recursos.
- Evitar que dos mecánicos tengan asignado el mismo recurso al mismo tiempo.
- Organizar las citas teniendo en cuenta tanto el horario como los recursos disponibles.
- Gestionar los cambios o cancelaciones de las citas.

El objetivo final es que el taller pueda organizar su jornada de trabajo de una forma más eficiente, evitando que los mecánicos tengan que esperar porque una máquina, herramienta o elevador esté siendo utilizado por otro compañero.


## Documentación adicional y configuración
La documentación del objetivo 0 está en doc/objetivos/objetivo-0.md.