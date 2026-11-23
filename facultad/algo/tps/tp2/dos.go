package main

import (
	"fmt"
	"strings"
	TDADic "tdas/diccionario"
	"time"
)

const CANTIDAD_REQUESTS_DOS = 5
const PERIODO_REQUESTS_DOS = 2

// Agrega un sitio al diccionario de sitios, incrementando su contador de solicitudes si ya existía, o agregándolo con contador 1 si no existía.
func agregarSitio(sitio string, sitios TDADic.Diccionario[string, int]) {
	if sitios.Pertenece(sitio) {
		cant := sitios.Obtener(sitio)
		sitios.Guardar(sitio, cant+1)
	} else {
		sitios.Guardar(sitio, 1)
	}
}

/*
procesa una línea de log y actualiza la información de la ip correspondiente en el diccionario de ips, visitantes y sitios.
- en el diccionario de visitantes guarda en tiempo logaritmico O(log v), con v la cantidad de visitantes distintos.
- guarda los sitios en O(1) promedio.
*/
func procesarRegistro(log string, ips TDADic.Diccionario[string, *ip], visitantes TDADic.DiccionarioOrdenado[*ip, int], sitios TDADic.Diccionario[string, int]) {
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
		visitantes.Guardar(ipActual, 0)
	}
	ipActual.AgregarRequest(fecha, metodo, sitio)
	agregarSitio(sitio, sitios)
}

// Devuelve true si la diferencia entre las fechas de los requests es menor o igual al periodo establecido para DoS
func rompePeriodoDoS(request1, request2 ipRequest) bool {
	return request2.fecha.Sub(request1.fecha) < PERIODO_REQUESTS_DOS*time.Second
}

/*
Recorre las requests de una ip y determina si es un atacante DoS, es decir, si tiene 5 requests en un periodo de 2 segundos.
Complejidad: O(r) con r la cantidad de requests de la ip.
*/
func esAtacante(ip *ip) bool {
	requests := ip.requests
	if requests.Largo() < CANTIDAD_REQUESTS_DOS {
		return false
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
			return true
		}
	}
	return false
}

/*
Recorro todo el diccionario O(n) (n cantidad de registros), y dado que el archivo de log está ordenado cronológicamente,
intento acumular 5 requests en un plazo de 2 segundos a medida que recorro los registros.

Complejidad: O(n) con n la cantidad de registros en el archivo de log.
*/
func obtenerAtacantes(ips TDADic.Diccionario[string, *ip]) []*ip {
	atacantes := make([]*ip, 0)
	for iterIps := ips.Iterador(); iterIps.HayAlgoMas(); iterIps.Avanzar() {
		_, ip := iterIps.VerActual()
		if esAtacante(ip) {
			atacantes = append(atacantes, ip)
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
	atacantes = OrdenarIps(atacantes)
	for _, ip := range atacantes {
		fmt.Println("DoS: " + ip.VerIp())
	}
	fmt.Println("OK")
}
