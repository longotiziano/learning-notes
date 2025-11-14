# Generative AI.
Resumen del curso de Inteligencia Artifical Generativa aplicada a la ingeniería de datos
## Data Engineering and Generative AI.

### IA Generativa para Ingeniería de datos.
* En la Ingeniería de datos, se puede implementar la IA generativa en aspectos como:
    - `Generación de datos sintéticos`: "Imita" los datos del mundo real, crea diversos y
    representativos datasets, cumple con los requerimientos de privacidad y utiliza 
    generalmente los modelos GANs. Un ejemplo de herramienta podría ser `Tonic.ia`, o también `Hazy Tool`,
    que imita un dataset a partir del original pero cambiando los datos sensibles.
    - `Anonimato de datos`: Modifica los datos originales para proteger la privacidad individual,
    pero protegiendo las características originales -> _General Data Protection Regulation (GDPR)_.
    Utilizado para entrenamiento de modelos, entre otros objetivos sin comprometer la privacidad del
    usuario. Un ejemplo de herramienta podría ser `Tonic.ia`.
    - `Aumento en calidad de datos`: Identificación y corrección en anomalías de datos e inconstistencias. 
    Aumento significativo en la calidad de la información. Un ejemplo de herramienta podría ser `DataRobot`,
    `Databricks AutoML` o `SyntheticGuru`.
    - `Exploración de datos y visualización`: Con simulación de escenarios, aumento de información, entre otros.
    Un ejemplo podría ser `GPT-5`, de OpenIA.

### Cómo los ingenieros de datos aprovechan el poder de la IA Generativa.
- `IA Generativa contra modelos ML discriminativos`: La diferencia principal que tienen entre ellos es,
principalmente, que mientras los modelos generativos sirven para crear, simular o generar, los discriminativos
se enfocan en la predicción, clasificación y distinción. Pero cabe recalcar que ambos son importantísimos a
la hora de trabajar como ingeniero de datos.
* `Modelos de IA Generativa`:
    - `Generative Adversarial Networks (GANs)`: Se basan en dos redes neuronales que compiten entre ellos: el
    `Generador`, que crea los datos, y el `Discriminador`, que evalua si esos datos son falsos o acordes al 
    dataset original.
    - `Variational autoencoders (VAEs)`: Reduce los datos en una representación comprimida y luego los 
    reconstruye.
    - `Transformer`: Se basa en el mecanismo del _self-attention_, que permite establecer jerarquías y
    relaciones entre las palabras o texto que procesa con la ayuda del _feed-foward layer_, que es la red 
    neuronal.
- `Responsabilidades de los ingenieros de datos`: Mantenimiento de la infraestructura de datos, desarrollo
de pipelines y calidad de los datos.
* Herramientas y tecnologías comúnmente utilizadas por los ingenieros:
    - `TensorFlow y PyTorch`: Ambas open source, que tienen soporte para GANs, VAEs y autoencoders.
    Son también flexibles, tienen una extensa comunidad y muy utilizada en ambientes de producción.
    Permiten desarrollar modelos de IA generativa.
    - `Hugging Face transformers`: Integración con las bibliotecas mencionadas, modelos pre-entrenados,
    interfaz amigable...
    - `Entornos de desarrollo`: Jupyter Notebooks y Google Colab.

### Rol de la IA Generativa en la arquitectura de datos.
* `Ventajas de usar frameworks para IA generativa en arquitecturas de datos`:
    - Procesos de diseño más rápidos y simplificados.
    - Mayor desempeño de los sistemas y optimización de recursos.
    - Reducción de costos y esfuerzo humano.
    - Arquitecturas predictivas y adaptables, preparadas para futuros cambios en volumen o tipo de datos. 

### Automatizando el diseño de esquemas con IA generativa.
* `Herramientas recomendadas para la creación de esquemas`:
    - Staq: Basado en datos de ejemplo, crea relaciones y campos.
    - StichData: Analiza tipos de datos, relaciones y patrones.
    - SchemaWriter.ai: Generador de esquemas optimizado para _Search Engine Optimization_ (SEO) de páginas web.
    - En cuanto a la normalización, podemos encontrar Google's AutoML y OpenAI's Codex.

### Data Augmentation.
* Para los diferentes tipos de datos, hay diferentes IAs que se acoplan al aumento de datos:
    - `Estructurados`: CTGAN, SDV. Útiles para balancear datos, llenar nulos y en tema privacidad.
    - `Semi-estructurados`: GPT o Copilot.
    - `No-estructurados`: StyleGAN2, BigGAN.

### IA para setups de infraestructura.
* Las tareas en el setup suelen ser de: 
    - Configuración de servidor
    - Preparado de network
    - Seguridad de los 

### Consideraciones a la hora de usar IA generativa en industrias.
* Consideraciones de:
    - `Datos`: Su efectividad depende en la calidad de los datos. Muy importante eliminar bias al
    momento del entrenamiento.
    - `Modelo`: Basado en la _explainability_ (proceso de pensamiento), que se mejora a través de la atribución
    de características y los plots dependientes; y la _interpretability_ (proceso del output).
    - `Éticas`: No deben ser utilizadas para desarrollar actividades maliciosas, y con uso responsable.
* En las respectivas empresas también podemos encontrar consideraciones particulares:
    - Finanzas: Manejo de datos sensibles, por lo que sus modelos deben ser robustos ante posibles
    ataques maliciosos. Predominan los modelos.
    - Salud: También manejan datos sensibles, pero es importante que los modelos sean precisos. Predomina
    la ética.
    - Retail: La precisión también es muy importante. Predomina la calidad de los datos.
- `Bias`: Sesgo que el modelo aprende de los datos.

### Desafios a la hora de usar IA generativa.
* Los desafíos son:
    - `Técnicos`: Calidad de los datos, complejidad e interpretación del modelo, recursos computacionales,
    falta de herramientas... La disponibilidad de los datos y su calidad impactan en el rendimiento
    gravemente. En consecuencia de no tener todas estas características bien vistas y atendidas, se
    podría dar información imprecisa, errores en la arquitectura y en la evualuación del modelo.
    - `Organizacionales`: Problemas de derechos de autor (las organizaciones deben crear sus propios modelos de 
    IA para crear), requerimiento de gente experta, integración con sistemas ya existentes, determinar el 
    rendimiento real de la inversión...   
    - `Culturales`: Riesgo, compartimiento de datos, confianza y transparencia, y el continuo aprendizaje y 
    adaptación.d
