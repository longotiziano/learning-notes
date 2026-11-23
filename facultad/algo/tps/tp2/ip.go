package main

/*
En este módulo desarrollo el tipo de dato de ip y demás funciones auxiliares
*/

import (
	"strconv"
	"strings"
	TDALis "tdas/lista"
	"time"
)

const RANGO_IP = 256

type ipRequest struct {
	fecha  time.Time
	metodo string
	sitio  string
}

type ip struct {
	parteA   int
	parteB   int
	parteC   int
	parteD   int
	requests TDALis.Lista[ipRequest]
}

// Agrega una request y su información para una IP
func (ip *ip) AgregarRequest(fecha time.Time, metodo string, sitio string) {
	info := ipRequest{fecha, metodo, sitio}
	ip.requests.InsertarUltimo(info)
}

// devuelve la ip en formato string
func (ip *ip) VerIp() string {
	return strconv.Itoa(ip.parteA) + "." + strconv.Itoa(ip.parteB) + "." + strconv.Itoa(ip.parteC) + "." + strconv.Itoa(ip.parteD)
}

/*
compara ips numéricamente, devolviendo un número positivo si ip1 > ip2, 0 si son iguales y un número negativo si ip1 < ip2
*/
func cmpIps(ip1, ip2 *ip) int {
	if ip1.parteA != ip2.parteA {
		return ip1.parteA - ip2.parteA
	}
	if ip1.parteB != ip2.parteB {
		return ip1.parteB - ip2.parteB
	}
	if ip1.parteC != ip2.parteC {
		return ip1.parteC - ip2.parteC
	}
	return ip1.parteD - ip2.parteD
}

// Devuelve un puntero a una IP
func CrearIp(ipStr string) *ip {
	requests := TDALis.CrearListaEnlazada[ipRequest]()
	numsIpStr := strings.Split(ipStr, ".")
	numsIp := make([]int, 4)
	for i, numIp := range numsIpStr {
		n, err := strconv.Atoi(numIp)
		if err != nil {
			panic(err)
		}
		numsIp[i] = n
	}
	return &ip{numsIp[0], numsIp[1], numsIp[2], numsIp[3], requests}
}

// uso radix sort con counting sort como auxiliar, lo que me da una complejidad de O(4(v + k)) = O(v) con v la cantidad de ips y k = 256
func OrdenarIps(ips []*ip) []*ip {
	ipsOrdenadas := countingSortIps(ips, func(ip *ip) int { return ip.parteD })
	ipsOrdenadas = countingSortIps(ipsOrdenadas, func(ip *ip) int { return ip.parteC })
	ipsOrdenadas = countingSortIps(ipsOrdenadas, func(ip *ip) int { return ip.parteB })
	ipsOrdenadas = countingSortIps(ipsOrdenadas, func(ip *ip) int { return ip.parteA })
	return ipsOrdenadas
}

// counting sort para ordenar ips según una parte de la ip (A, B, C o D) que se pasa como parámetro
func countingSortIps(ips []*ip, selecParte func(*ip) int) []*ip {
	k := RANGO_IP
	frecuencias := make([]int, k)
	for _, ip := range ips {
		parte := selecParte(ip)
		frecuencias[parte]++
	}
	posiciones := make([]int, k)
	for i := 1; i < k; i++ {
		posiciones[i] = posiciones[i-1] + frecuencias[i-1]
	}
	salida := make([]*ip, len(ips))
	for _, ip := range ips {
		parte := selecParte(ip)
		salida[posiciones[parte]] = ip
		posiciones[parte]++
	}
	return salida
}
