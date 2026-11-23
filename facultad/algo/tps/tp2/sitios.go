package main

import (
	TDADic "tdas/diccionario"
	TDAHeap "tdas/heap"
)

/*
En este módulo desarrollo el tipo de dato de sitio y demás funciones auxiliares
*/

type sitio struct {
	nombre  string
	visitas int
}

// si sitio1.visitas < sitio2.visitas devuelve < 0
func cmpSitios(sitio1, sitio2 sitio) int {
	return sitio1.visitas - sitio2.visitas
}

func CrearSitio(nombre string, visitas int) sitio {
	return sitio{nombre, visitas}
}

func crearHeapSitios(sitios TDADic.Diccionario[string, int]) TDAHeap.ColaPrioridad[sitio] {
	arrSitios := make([]sitio, 0)
	for iter := sitios.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		sitio, visitas := iter.VerActual()
		if visitas > 0 {
			arrSitios = append(arrSitios, CrearSitio(sitio, visitas))
		}
	}
	return TDAHeap.CrearHeapArr(arrSitios, cmpSitios)
}
