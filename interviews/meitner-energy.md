# Preparación para la entrevista de Meitner Energy

5. Si te piden obtener el monto total de ventas acumulado día por día, ¿qué función de ventana usarías y por qué?

> Usaría una función de ventana con SUM() OVER (PARTITION BY cliente_id ORDER BY fecha) para calcular el acumulado, ya que las funciones de ventana permiten calcular agregados acumulativos sin perder el detalle de cada fila. 

1. Para optimizar el uso de memoria en el cargado de un CSV con pandas de un millón de filas sería identificar las columnas que quiero al momento de leer el archivo con `.usecols`. Pero también está la alternativa de usar Polars, la cual es una librería más moderna, con características como el paralelismo que optimizan el cargado de archivos grandes.
s
2. Para unir ambos DataFrames se utilizaría `ventas.merge(clientes, how='tipo_de_join')`

3. Para detectar filas con valores nulos se utiliza `df.isna()`. La decisión del qué hacer con ellas variaría dependiendo si es reemplazable ese valor o no y su importancia dentro del dataset. Por ejemplo, si falta el país de un cliente, pero tenemos su ciudad, son valores nulos que pueden ser rellenados.

4. Podría hacer `pd.read_excel('filename.xlsx', sheet_name=None) para obtener el diccionario con los dataframes y posteriormente pd.concat(dfs.values(), ignore_index=True)` el ignore index es importante para que no se repitan índices entre los registros

5. Eso sería bastante simple, utilizaría pandas para el proceso completo. Primero haría la lectura del archivo utilizando try en caso de no encontrar el nombre, luego haría el wrangling necesario o pedido, finalmente utilizaría to_excel donde en el nombre del archivo utilizaría una f-string con la fecha actual, utilizando la librería datetime.