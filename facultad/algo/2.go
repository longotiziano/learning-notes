package main
import "fmt"

a := 5 // infiere el tipo de dato
var b int = 6 // mas explicito

c, d := 10, 11
a, b = b, a

const e int = 11

var (
	a int = 5
	b byte = 16
	c float = 9.0
)

// SIEMPRE asigna un valor default:
// 0 para numeros
// false para bools
// "" para strings

func sumar(a int, b int) int {
	return a + b
}

func sumar_named_value(a int, b int) (suma int) {
	suma = a + b
	return
}

func sumarYRestar(a, b int) (int, int) {
	return a+b, a-b
}

if condicion || otra_condicion {
	// coso
} else {
	// igual q en C
}

nombre := "Tiziano"

for i:=0; i<10; i++ {
	fmt.Println("HOLA %s %d VECES!\n", nombre, i)
}

// para mayor placer
for i := range 10 {
	fmt.Printf("%d", i)
}

fmt.Println("Que dia es hoy?")
today := time.Now().Weekday()
switch today {
case time.Monday:
	fmt.Println("Lunes :(")
}
case time.Saturday, time.Sunday:
	fmt.Println("Finde :)")
default:
	fmt.Println("No sé (?)")

// Notese que las variables se encapsulan dentro de cada scope, definido con cada corchete
// Es por esto que en este ejemplo, podemos "pisar" y redeclarar `b` dentro de cada iteración
// y no únicamente que no va a dar error, si no también que al salir del scope del `for` se 
// borrará el valor que se le fue asignado
b := 5
for i := 1; i<4;i++ {
	b := i
	fmt.Println(b) // 1 2 3
}
fmt.Println(b) // 5

// el defer nos permite que una línea de código se ejecute justo antes de la finalización de
// la función. Esto se hace en orden de pila (heap).
func main() {
	defer fmt.Println("world")
	fmt.Println("hello ")
}

// arreglos
var arreglo [4]int = [4]int{1, 2, 3, 4}
fmt.Println(arreglo) // [1 2 3 5]
var arreglo2 [4]int = [3]int{1, 2, 3}
fmt.Println(arreglo) // [1 2 3 0] (x el valor default)

func modificar(arr [4]int) {
	arr[0] = 0 // solo vive dentro del scope de modificar, pq no es un reference type
}
func modificarRef(arr [4]*int) {
	arr[0] = 0 // este si modifica
}
func main() {
	arreglo := [4]int{1, 2, 3, 4}
	modificar(&arreglo)
	fmt.Println(arreglo) // [0 2 3 4]
}

// Slicing 
numeros := [6]int{1, 2, 3, 4, 5}
var slice []int = numeros[1:3] // fin no incluido
fmt.Println(slice) // [2 3]
// también se pueden crear con make() para que queden en el heap

// Los slices pueden expandirse de manera dinámica
s := []int{}
s = []int{2, 3, 5, 7, 11, 13}
s = s[:4]

// si accedés a la posición inválida de un slice, el programa falla
// no se rellenan los espacios con 0 porque no está reservada esa memoria

// manejo de errores
panic("mensaje de error") // casos irrecuperables!
recover() // captura al panic

func modificar(silce []int) {
	slice[0] = 0
}
func main() {
	slice:= []int{1, 2, 3, 4}
	modificar(slice)
	fmt.Println(slice) // [0 2 3 4]
}

// valor nulo
var slice []int
fmt.Println(slice==nil) // true

var a []int
var b *int
fmt.Println(a==nil) // true
fmt.Println(b==nil) // true
fmt.Println(b==a) // false por el tipo, no podés comparar distintos types

// no se puede desreferenciar una dirección nula (nil)

// metodos de slices
slice := []int{1, 2, 3}
slice2:=append(slice, 4)
slice = append(slice, slice2)
copy(dst, src)

for ind, valor := range slice