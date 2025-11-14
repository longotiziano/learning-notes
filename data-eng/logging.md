# logging

## Loggeo en Python
- **Logger**: Es el objeto que usas en tu código (`logger = logging.getLogger("mi_logger")`). Define el nivel mínimo de logs 
y qué handlers usar.
- **Handler**: Determina dónde se envían los logs: consola, archivo, CSV, DB, etc.
- **Formatter**: Define el formato del mensaje: qué info mostrar (fecha, nivel, mensaje, módulo, etc.)

## Estructura recomendada para logging
- Al momento de hacer un loggeo, es recomendado seguir la siguiente estructura:
    -> Cuándo pasó -> Dónde pasó -> Qué pasó -> Con qué datos