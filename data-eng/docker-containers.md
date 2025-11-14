# Docker y Containers

## 1. Características de un Container
- Virtualiza el sistema operativo.  
- Rápido, liviano, aislado, seguro y portable.  
- Requiere menos memoria.  
- Las librerías y binarios dentro del container permiten su funcionamiento.  
- Una máquina puede correr múltiples containers simultáneamente.

## 2. Beneficios
- Desarrollo rápido de aplicaciones con menor costo.  
- Mejor utilización de recursos del sistema.  
- Soporte para aplicaciones de nueva generación (next-gen).

## 3. Desafíos
- Seguridad afectada si hay un ataque al sistema operativo.  
- Dificultad de administrar muchos containers en escenarios complejos.

## 4. Docker
- Escrito en **Go**.  
- Utiliza el kernel de Linux para funcionalidades de containers.  
- Entorno para definir, construir y correr containers.  
- Adecuado para entornos ligeros; menos recomendado si:
  - Se requiere **alto rendimiento o seguridad crítica**.  
  - Se necesita **arquitectura Monolith** o **GUIs complejas**.  
  - Se corren tareas limitadas.

### 4.1 Dockerfile
- **FROM:** define la imagen base.  
- **CMD:** define los comandos predeterminados del container.  
- **RUN:** ejecuta comandos arbitrarios durante la construcción.  

### 5. Imágenes y Containers
* `Imágenes`:
    - Plantillas que contienen instrucciones para crear containers.
    - Se crean mediante Dockerfiles.
    - Cada línea del Dockerfile genera una capa.
    - Formato de creación: hostname/repository:tag.
* `Containers`:
* `Instancias de imágenes que pueden`:
    - Crearse
    - Correr
    - Eliminarse
* `Comunicación`: se conectan mediante networks.
* `Almacenamiento`:
    - `Volúmenes`: guardan datos y procesos cuando el container está apagado.
    - `Bind mounts`: conectan almacenamiento del host con el container.
    - `Plugins`: permiten conexión a almacenamiento externo.

6. Arquitectura de Docker
* `CLI / REST API`: envía instrucciones al Docker Host.
* `Daemon`: escucha, procesa y corre comandos; administra imágenes y registros.
* `Registros`: donde se guardan imágenes.
    - push: subir imágenes al registro.
    - pull: descargar imágenes a otra máquina.

7. Orquestación de Containers
- Necesaria en proyectos con muchos containers.
- Automatiza el ciclo de vida: despliegue, administración, escalamiento, red y disponibilidad.
- Simplifica la complejidad y mejora la eficiencia.

8. Configuración
* `Archivos en YAML o JSON`:
    - Configuran containers, recursos y redes.
    - Guardan logs.
    - Permiten cronogramas de despliegue.
    - Asignan hosts predefinidos.