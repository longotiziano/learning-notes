package tp2

import (
	"strings"
	TDADic "tdas/diccionario"
	TDAHeap "tdas/heap"
	"time"
)

const CANTIDAD_REQUESTS_DOS = 5
const PERIODO_REQUESTS_DOS = 2

/*
procesa una línea de log y la guarda en el diccionario según corresponde:
- si ya existe la ip, agrega la request
- caso contrario, la crea desde 0 y agrega la request
*/
func procesarRegistro(log string, ips TDADic.Diccionario[string, *ip]) {
	campos := strings.Split(log, "\t")
	claveIp := campos[0]
	fecha, err := time.Parse(time.RFC3339, campos[1])
	if err != nil {
		panic(err)
	}
	metodo := campos[2]
	sitio := campos[3]

	var ipActual *ip
	if ips.Pertenece(claveIp) {
		ipActual = ips.Obtener(claveIp)
	} else {
		ipActual = CrearIp(claveIp)
		ips.Guardar(claveIp, ipActual)
	}
	ipActual.AgregarRequest(fecha, metodo, sitio)
}

// Devuelve true si la diferencia entre las fechas de los requests es menor o igual al periodo establecido para DoS
func rompePeriodoDoS(request1, request2 ipRequest) bool {
	return request2.fecha.Sub(request1.fecha).Seconds() <= PERIODO_REQUESTS_DOS
}

/*
Recorro todo el diccionario O(n) (n cantidad de registros), y dado que el archivo de log está ordenado cronológicamente,
intento acumular 5 requests en un plazo de 2 segundos a medida que recorro los registros.

Ip es atacante si tiene más de 5 requests en menos de 2 segundos.

Complejidad: O(n) con n la cantidad de registros en el archivo de log.
*/
func obtenerAtacantes(ips TDADic.Diccionario[string, *ip]) []*ip {
	atacantes := make([]*ip, 0)
	for iterIps := ips.Iterador(); iterIps.HayAlgoMas(); iterIps.Avanzar() {
		_, ip := iterIps.VerActual()
		requests := ip.requests
		if requests.Largo() < CANTIDAD_REQUESTS_DOS {
			continue
		}
		reqs := make([]ipRequest, requests.Largo())
		i := 0
		for iter := requests.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
			reqs[i] = iter.VerActual()
			i++
		}
		for i := 0; i <= len(reqs)-CANTIDAD_REQUESTS_DOS; i++ {
			limInf := reqs[i]
			limSup := reqs[i+CANTIDAD_REQUESTS_DOS-1]
			if rompePeriodoDoS(limInf, limSup) {
				atacantes = append(atacantes, ip)
				break
			}
		}
	}
	return atacantes
}

/*
Recibe un diccionario con todas las IPs y sus requests, y muestra por salida estándar las IPs que son atacantes DoS.
Complejidad: O(n) con n la cantidad de registros en el archivo de log.
*/
func mostrarDoS(ips TDADic.Diccionario[string, *ip]) {
	atacantes := obtenerAtacantes(ips)
	heap := TDAHeap.CrearHeapArr(atacantes, cmpIps)
	for !heap.EstaVacia() {
		ip := heap.Desencolar()
		println("DoS: " + ip.VerIp())
	}
	println("OK")
}

// Mostradas en orden asc numéricamente
// Formato:
// Visitantes:
// \t<IPs...>
//
// O(v) en el peor caso. O(log v) promedio (v cantidad de visitantes)
//
// Seria el total entre todos los archivos
func mostrarVisitantes()

// Mostradas en orden asc numéricamente por cantidad de solicitudes
// Formato:
// Sitios más visitados:
// \t<sitios...> - <solicitudes>
//
// O(s+k log s) con s cantidad de diferentes sitios en toda la historia y k el parámetro
//
// Seria el total entre todos los archivos
func mostrarMasVisitados(k int)
