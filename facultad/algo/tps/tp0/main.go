package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func leerArchivo(ruta string) []int {
	file, _ := os.Open(ruta)
	defer file.Close()

	var numeros []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		num, _ := strconv.Atoi(linea)
		numeros = append(numeros, num)
	}

	return numeros
}

func main() {
	vector1 := leerArchivo("archivo1.in")
	vector2 := leerArchivo("archivo2.in")
	var vectorOrdenado []int
	if ejercicios.Comparar(vector1, vector2) == -1 {
		ejercicios.Seleccion(vector2)
		vectorOrdenado = vector2
	} else {
		ejercicios.Seleccion(vector1)
		vectorOrdenado = vector1
	}

	for _, value := range vectorOrdenado {
		fmt.Println(value)
	}
}
