# Visitas técnicas y escalamientos

## Cuándo abrir visita

La visita se agenda cuando CODI completa el flujo y ofrece esa acción. Antes de crear una nueva:

- Comprobar producto/servicio, órdenes y citas pendientes.
- Si la visita pendiente corresponde a otro problema, usar `REALIZAR NUEVO DIAGNÓSTICO`.
- Si la visita existente sigue siendo pertinente y el período es razonable, las notas sugieren reagendar/modificarla en vez de abrir otra; confirmar el flujo.
- Consultar incidencias generales antes de crear orden individual.

## Workforce, OT y citas

**Workforce** es el sistema de coordinación de citas y personal técnico. Hay que distinguir:

- **Orden de trabajo (OT/WO):** tareas que deben realizarse.
- **Cita de servicio:** fecha/hora y organización de la visita para un servicio.
- Una OT puede tener varias citas por cancelación, bloqueo, anulación, reagenda o por requerir diagnóstico e instalación separados.

Estados de OT anotados:

| Estado | Significado registrado |
|---|---|
| Registrada | No debería ser el estado normal observado. |
| Programada | Antes del día de visita; todavía sin técnico asignado. |
| Asignada | Técnico asignado. |
| Movilizando | Se envía WhatsApp al cliente con datos del técnico, según la nota. |
| En ejecución | Visita en curso. |
| Bloqueada | No puede cumplirse por un impedimento; se informa por app. |
| Cerrada / cerrada parcial | Trabajo finalizado total o parcialmente. |
| Anulada | Cancelación/anulación de la OT. |

Estados de reserva FAN anotados: registrada (espera confirmación), confirmada, cumplida y cancelada. Confirmar el nombre/estado en la pantalla actual.

## Reserva, reagenda y comunicación

Ruta anotada para reagendar en FAN:

```text
Caso -> Relacionado -> Reservar cita -> elegir horario -> confirmar -> verificar estado
```

También se menciona `Gestión de clientes -> Reserva de citas`. No crear una segunda orden solo por no ver de inmediato la cita: verificar OT/WO y el estado. En el formulario confirmar cita, tipo de teléfono, código de área, número y comentario conciso para técnico. CODI puede enviar ya el detalle técnico.

Al cliente:

- Confirmar día/franja, teléfono y número de contacto.
- Explicar requisitos de acceso anotados (persona adulta de 18+ y acceso a techo/terraza/garaje cuando aplique) usando el speech vigente.
- Explicar seguimiento por WhatsApp/WIMT (Where Is My Tech?) cuando esté disponible.
- Indicar cómo reagendar por app y cuándo debería aparecer la cita, sin comprometer un plazo no confirmado.
- Para asistencia en hogar, atender condiciones meteorológicas y de seguridad conforme a política.

## Bloqueos, cancelaciones y errores

Motivos anotados para cancelación: clima/lluvia, cliente ausente, cancelación por el cliente, interrupción del proceso por el cliente, zona peligrosa, “sin visitar” y “ausente”.

Motivos anotados para bloqueo: sin visitar, ya instalado, edificio sin armar, suspensión por contingencia, ausente en puerta, meteorología y cliente no quiere. Interpretar el código/motivo exacto en Workforce; no asumir que el motivo es correcto o que la acción siguiente es una nueva visita.

Errores de `Continuar` anotados:

- **WO ya generada:** localizar y revisar la orden existente; no duplicar.
- **Sin fechas disponibles:** en las notas se indica que puede escalarse automáticamente y ofrecerse agenda posterior; verificar notificación y plazo del sistema.
- **Error al cargar formulario:** consultar iTracker/procedimiento.
- Consultar iTracker ante problemas técnicos de la herramienta.

Verificación de “ausente en puerta”: la nota describe un SMS al cliente que pregunta si está en el domicilio; si confirma, puede volver el mismo técnico dentro de hasta dos horas; de lo contrario, se bloquea/cierra y se orienta a reagendar. Confirmar que ese mecanismo esté vigente y seguir lo que refleje el sistema.

Las notas indican que una reagenda puede contar como service/métrica, con excepción anotada para casos propios; no usar este apunte como regla de cálculo actual.

## Horarios de técnicos anotados

Las jornadas se describen de lunes a sábado, con guardias domingos/feriados; franjas generales de 7:00 a 21:30/22:00.

- Servicio técnico: 8–14, 13–19, 17:30–22 y 21–23.
- Instalaciones/altas: 8–14 y 13–18.

Son horarios de una nota de capacitación, no promesa de disponibilidad. Ofrecer únicamente las franjas que muestre el sistema.

## Escalamiento a L2

Escalar **únicamente si CODI lo indica**, con flujo completo y luego de verificar que no haya otro escalamiento en las últimas 72 h.

No se escala solo porque:

- El cliente consulta por un masivo.
- El tema es software interno.
- Quiere hablar con alguien de mayor nivel.
- No está en el domicilio o no tiene PC/equipo requerido para los chequeos.
- No se ha podido resolver antes de completar el flujo.

Usar el mecanismo destinado a cada situación. Si el cliente no puede completar una prueba, no inventar resultado ni presentar eso como diagnóstico completo.

## Preparación y llamada tripartita

1. Explicar motivo, próximos pasos y tiempos que estén confirmados.
2. Completar la planilla correcta. Si no está a mano, buscar “Planillas” en BC.
3. Incluir solo campos aplicables: URL, SN/PON ID de ONT, ID de dispositivo desde MINERVA, zona, NAT, IP, DNS, resultado de `nslookup`, acceso por IP, número de contrato, contacto y resumen de pruebas/observaciones.
4. Hacer llamada tripartita si lo indica el procedimiento.
5. Presentar al cliente, resumir el problema a PWTF/L2 y responder preguntas.
6. Esperar confirmación para cortar y dejar estado conforme al flujo.

Un speech anotado explica que el caso se deriva al soporte de segundo nivel y que participarán las tres personas; dirigirse al área como “superior” según la indicación de capacitación, con lenguaje claro y sin prometer resolución.

Al cumplir requisitos, una nota indica que el caso puede pasar de `Anulado` a `Derivado` y que al aviso del cliente se responde con la planilla. Confirmar estado y respuesta correcta en el sistema.

## CBS, PWTF y ventana del caso

- **CBS:** en las notas, gestiona escalamientos de 00:00 a 9:00 y callback luego de la derivación, hasta 72 h.
- **Power to Front (PWTF):** recibe llamadas de 9:00 a 00:00 y busca resolver en la misma llamada.
- Caso de soporte puede permanecer abierto hasta 72 h y resolverse o pasar a L3, según el apunte.

Horarios y SLA deben verificarse en la operatoria vigente antes de informarlos.

## Continuidad de servicio

Si corresponde, revisar alternativas de gigas, Wi-Fi Zone y Flow App en [Atención y herramientas](01-atencion-y-herramientas.md). No ofrecerlas automáticamente: verificar convergencia, titularidad, causa y elegibilidad.
