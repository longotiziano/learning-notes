package tp2

/*
En este módulo desarrollo el tipo de dato de IP.
*/

import (
	"strconv"
	"strings"
	TDALis "tdas/lista"
	"time"
)

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
compara ips numéricamente, devolviendo un número negativo si ip1 > ip2, 0 si son iguales y un número positivo si ip1 < ip2
destinado a ser usado como función de comparación para un heap de mínimos.
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
