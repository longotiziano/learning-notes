package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADic "tdas/diccionario"
)

const COMM_AGREGAR = "agregar_archivo"
const COMM_VISITANTES = "ver_visitantes"
const COMM_VISITADOS = "ver_mas_visitados"
const ERROR_COMM_INVALIDO = "Error en comando"
const ERROR_ENTRADA_INVALIDA = "Error en entrada estándar"

// muestra el error de comando inválido por stderr.
func mostrarErrorComm(comando string) {
	fmt.Fprintf(os.Stderr, "%s %s\n", ERROR_COMM_INVALIDO, comando)
}

// convierte el parámetro del comando introducido a entero. En caso de fallar, muestra error de comando inválido.
func convertirParamInt(comando, param string) (int, bool) {
	n, err := strconv.Atoi(param)
	if err != nil {
		mostrarErrorComm(comando)
		return 0, false
	}
	return n, true
}

/*
Consulta la entrada estándar y procesa el comando según corresponda, devolviendo OK en caso de comando válido.
En caso de no encontrar el comando, se imprime "Error en comando <comando>" por stderr y se finaliza la ejecución.
*/
func EscucharComandos() {
	scanner := bufio.NewScanner(os.Stdin)
	visitantes := TDADic.CrearABB[*ip, int](cmpIps)
	sitios := TDADic.CrearABB[string, int](strings.Compare)
	for scanner.Scan() {
		lineaComm := scanner.Text()
		listaComm := strings.Split(lineaComm, " ")

		if len(listaComm) < 2 {
			mostrarErrorComm(lineaComm)
			continue
		}
		comando := listaComm[0]
		param := listaComm[1]
		switch comando {
		case COMM_AGREGAR:
			AgregarArchivo(param, visitantes, sitios)
		case COMM_VISITANTES:
			if len(listaComm) != 3 {
				mostrarErrorComm(comando)
				continue
			}
			limInf := CrearIp(param)
			limSup := CrearIp(listaComm[2])
			VerVisitantes(limInf, limSup, visitantes)
		case COMM_VISITADOS:
			if len(listaComm) == 2 {
				cant, ok := convertirParamInt(comando, param)
				if !ok {
					mostrarErrorComm(comando)
					continue
				}
				SitiosMasVisitados(cant, sitios)
			} else {
				mostrarErrorComm(comando)
			}
		default:
			mostrarErrorComm(comando)
		}
	}
	if err := scanner.Err(); err != nil {
		mostrarErrorComm(err.Error())
	}
}

/*
Procesa un archivo de log completo. En caso de encontrar amenazas DoS, las devuelve por salida estándar.
Complejidad resultante: O(n log v), con n la cantidad de registros del archivo y v la cantidad de visitantes distintos.
*/
func AgregarArchivo(archivo string, visitantes TDADic.DiccionarioOrdenado[*ip, int], sitios TDADic.Diccionario[string, int]) {
	arch, err := os.Open(archivo)
	ips := TDADic.CrearHash[string, *ip]()
	if err != nil {
		mostrarErrorComm(COMM_AGREGAR)
		return
	}
	scanner := bufio.NewScanner(arch)
	for scanner.Scan() {
		log := scanner.Text()
		procesarRegistro(log, ips, visitantes, sitios)
	}
	if err := scanner.Err(); err != nil {
		mostrarErrorComm(err.Error())
		return
	}
	mostrarDoS(ips)
}

/*
Dado un rango de ips, imprime por salida estándar los visitantes que se encuentran en ese rango.
Complejidad: O(v) en el peor caso, y O(log v) en promedio para pocos visitantes, con v la cantidad de visitantes distintos.
*/
func VerVisitantes(limInf, limSup *ip, visitantes TDADic.DiccionarioOrdenado[*ip, int]) {
	fmt.Println("Visitantes:")
	for iter := visitantes.IteradorRango(&limInf, &limSup); iter.HayAlgoMas(); iter.Avanzar() {
		ip, _ := iter.VerActual()
		fmt.Println("\t" + ip.VerIp())
	}
	fmt.Println("OK")
}

/*
Dada una cantidad y un diccionario de sitios, imprime por salida estándar los sitios más visitados.
Complejidad: O(s + k log s), con s la cantidad de sitios distintos y k la cantidad de sitios a mostrar.
*/
func SitiosMasVisitados(cant int, sitios TDADic.Diccionario[string, int]) {
	heap := crearHeapSitios(sitios) // O(s)
	i := 0
	fmt.Println("Sitios más visitados:")
	for !heap.EstaVacia() && i < cant { // O(k log s) (heap de tamaño s k veces)
		sitio := heap.Desencolar()
		fmt.Printf("\t%s - %d\n", sitio.nombre, sitio.visitas)
		i++
	}
	fmt.Println("OK")
}
