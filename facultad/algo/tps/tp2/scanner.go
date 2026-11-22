package tp2

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
func convertirParamInt(comando, param string) int {
	n, err := strconv.Atoi(param)
	if err != nil {
		mostrarErrorComm(comando)
	}
	return n
}

/*
Consulta la entrada estándar y procesa el comando según corresponda, devolviendo OK en caso de comando válido.
En caso de no encontrar el comando, se imprime "Error en comando <comando>" por stderr y se finaliza la ejecución.
*/
func EscucharComandos() {
	scanner := bufio.NewScanner(os.Stdin)
	ips := TDADic.CrearHash[string, *ip]()
	for scanner.Scan() {
		lineaComm := scanner.Text()
		listaComm := strings.Split(lineaComm, " ")

		// los 3 comandos posibles siempre van acompañados de un parámetro
		if len(listaComm) < 2 {
			continue
		}
		comando := listaComm[0]
		param := listaComm[1]
		switch comando {
		case COMM_AGREGAR:
			AgregarArchivo(param, ips)
		case COMM_VISITANTES:
			cant := convertirParamInt(comando, param)
			VerVisitantes(cant, ips)
		case COMM_VISITADOS:
			cant := convertirParamInt(comando, param)
			SitiosMasVisitados(cant, ips)
		default:
			mostrarErrorComm(comando)
		}
	}
	if scanner.Err() != nil {
		mostrarErrorComm(ERROR_ENTRADA_INVALIDA)
	}
}

/*
Procesa un archivo de log completo. En caso de encontrar amenazas DoS, las devuelve por salida estándar.
Complejidad lineal O(n), con n la cantidad de registros del archivo.
*/
func AgregarArchivo(archivo string, ips TDADic.Diccionario[string, *ip]) {
	arch, err := os.Open(archivo)
	if err != nil {
		mostrarErrorComm(COMM_AGREGAR)
		return
	}
	scanner := bufio.NewScanner(arch)
	for scanner.Scan() {
		log := scanner.Text()
		procesarRegistro(log, ips)
	}
	if scanner.Err() != nil {
		mostrarErrorComm(COMM_AGREGAR)
		return
	}
	mostrarDoS(ips)
}

func VerVisitantes(cant int, ips TDADic.Diccionario[string, *ip])

func SitiosMasVisitados(cant int, ips TDADic.Diccionario[string, *ip])
