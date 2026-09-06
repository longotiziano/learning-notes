package main

import (
	"fmt"
	"strconv"
	"errors"
)

type Punto struct {
	X int
	Y int // o tambien: `x,y int` (minúsculas a propósito)
}

var p Punto = Punto{X: 1, Y: 3}
p = Punto{Y: 3}
p = Punto{1, 3}
p.Y = 5

var q *Punto = &Punto{X: 1, Y: 3}
q = &Punto{Y: 3}
q = &Punto{1, 3}
q.Y = 5
q = &Punto{}

&Punto{} === new(Punto)

func modificarPunto(p *Punto) {
	p.Y = 5
}

// METODOS
func (p Punto) DistanciaAlOrigen() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
func (p *Punto) Suma(otro Punto) Punto {
	return Punto{X: p.X + otro.X, Y: p.Y + otro.Y}
}
func (p *Punto) SetX (x int) { // importante el puntero en el tipo de dato!
	p.X = x
}
p.SetX(5)

// Anidación de Structs
type Direccion struct {
	calle string
	numero int
}

type Persona struct {
	direccion Direccion
	nombre string
	padre, madre *Persona // siempre referencia a si mismo (pointers), si faltara el puntero se generaría una recursión sin fin
}

jon := Persona{nombre: "join voight"}
angie := Persona{
	Direccion{"beverly hills st", 234},
	nombre: "angelina voight",
	&jon,
	nil
}

// por defecto, al imprimir se intenta llamar al método String, no si no existe, se usa la estrategia default.
// ejemplo:
func (d *Direccion) String() {
	fmt.Println("%s al %s", d.calle, d.numero)
}

// en los structs anidados puedo acceder al atributo, del más cercano y más lejano
fmt.Println(angie.nombre) // prioriza persona
fmt.Println(angie.calle) // da calle

// ERRORES
func ConectarServidor(ip IP) (*Servidor, error) {
	...
	if todoMal {
		return nil, fmt.Errorf("no se pudo conectar a la IP %v", ip)
	}
	return servidor, nil
}

serv, err := ConectarServidor()
if err != nil {
	...
}

// STRINGS
// Una cadena es un slice de bytes
// hay hasta 128 caracteres. Unicode puede usar hasta 4 bytes (acentos, emojis, otras letras)
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

// Unicode
var c rune
for _, c = range "chinardo Algo2 emoji" {
	fmt.Printf("Símbolo %c código Unicode %d\n", c, c)
}

// Métodos
strings.Replace
strings.Index
strings.HasPrefix

// Conversión a runa
cadena := "ááááááááá"
runes := []rune(cadena)

// Inputs
s := bufio.NewScanner(os.Stdin)
for s.Scan() {
	if s.Text() == "" {
		break
	}
	fmt.Printf("Leí: %s\n", s.Text())
}
err = s.Err()
if err!=nil {
	fmt.Println(err)
}

// Archivos
file, err := os.Create(ruta)
if err != nil {
	fmt.Printf("msg error")
}
defer file.Close()
datawriter := bufio.NewWriter(file) // escribe en el buffer
for _, dato := range []string{"hola", "como", "estas"} {
	_, err = datawriter.WriteString(dato + "\n")
	if err != nil {
		fmt.Printf("err")
	}
}
datawriter.Flush() // escribe en disco los registros acumulados en el buffer, y limpia el buffer