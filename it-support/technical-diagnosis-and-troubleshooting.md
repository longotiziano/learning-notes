# Technical Diagnostics and Troubleshooting Techniques
Arreglo de sistemas operativos (Windows, Linux y MacOS), network, documentación de procesos...

## Introduction to troubleshooting methodologies
En esta parte del curso se nos introduce a las metodologías y los errores comunes de hardware en las computadoras.

### Troubleshooting - Root cause analysis
Se divide en cuatro etapas:
1. **Identificación**: Ver que es lo que está pasando exactamente y entender los síntomas. Se necesita precisión y documentación.
2. **Análisis**: Complejidad de sistema, análisis profundo del problema, desde hardware y configuración de software hasta integraciones externas del sistema y interacciones del usuario. Cómo el problema ocurrió en un primer lugar y como hacer que no vuelva a pasar.\
**Se recomiendan:** 
    - los 5 por qués, para ir yendo más profundo en el análisis
    - el fishbone diagram: para identificar y separar causas 
    - el pareto chart 
    - kepner-tregoe: que es la metodología más formal, qué es y qué no es
3. **Acción**: Momento de elaborar y ejecutar una solución sostenible y robusta, no un parche.
4. **Verificación**: En esta parte ademas de realizar la verificación definitiva de la solución, también se aprovecha a documentar en detalle lo sucedido.
Otra forma de _troubleshootear_ es **prevenir** los errores.

### Happiness Engineering
Porfavor, hay que ser empáticos, pacientes, claros, follows-up personalizados y dar soluciones eficientes al cliente.

Algunas buenas ideas para una mejor experiencia:
- **Hacer las preguntas correctas**
- **Conseguir información**: logs, actualizaciones en la base de datos, indicadores de rendimiento donde halla una coincidencia con el problema ocurrido
- **Hacer un método lógico**: aislar el problema para identificarlo, testear las hipótesis que se te vayan ocurriendo y conocer las causas del problema
- **Uso de la tecnología**: software de diagnóstico, monitorear soluciones, realizar automatizaciones
- **Documentación**: Desde el inicio hasta el final, utilizala también para aprender y compartir.
- **Priorizar problemas, seguir guías paso a paso, escalar el problema cuando sea necesario...**

## Best practices for effective documentation
- Clearly identify and define the problem
- Document the investigation process
- Capture solutions and results: Detailing what worked (and what didn’t)
- Maintain documentation consistency: Ensure that all documentation follows a consistent format, style, and tone
- Use clear and concise language
- Incorporate visuals: Where applicable
- Regularly update documentation

# Cosas super interesantes que aprendí
- Si abris "Device Manager" en la computadora, podes ver el estado de todos los periféricos y funcionalidades de la computadora, permitiendo así una solución de erorres mucho más cercana.
- **IDEAL framework**: Identifying, Diagnosing, Establishing, Acting, Learning
- **SWOT analysis**: Strengths, weaknesses, opportunities, threats y a su vez el uso de matrices de decisión
- **Para defragmentación y optimización del disco**: Buscá "defr..." en el buscador de Windows y ya debería de aparecer. También en este caso para ofrecer una mejor solución le podés decir al usuario que justo en el mismo panel seleccione la opción de cronogramar la optimización y defragmentación del disco.
