package tp2

import (
	"strings"
)

/*
- hay varios servidores, con un archivo de log de dicho servidor
- cuando se llama a agregar_archivo, el programa guarda en memoria los datos del mismo. Se pierde la información cuando se corta el programa.

Formato log: 208.115.111.72	2015-05-17T11:05:15+00:00	GET	/corrector.html
*/

func DetectarOrden()

// recibe una string de log, y devuelve IP, fecha, método y sitio
func procesarRegistro(log string) (string, string, string, string) {
	campos := strings.Split(log, "\t")
	return campos[0], campos[1], campos[2], campos[3]
}

// procesa un log
// Si una misma IP realiza >= 5 peticiones en <= 2s
// O(n)
// Podríamos usar un diccionario para obtener las distintas IPs únicas y almacenar sus
func AgregarArchivo()

// Mostradas en orden asc numéricamente
// Formato:
// DoS: <IP>
func mostrarDoS()

// Mostradas en orden asc numéricamente
// Formato:
// Visitantes:
// \t<IPs...>
func mostrarVisitantes()

// Mostradas en orden asc numéricamente por cantidad de solicitudes
// Formato:
// Sitios más visitados:
// \t<sitios...> - <solicitudes>
func mostrarMasVisitados()
