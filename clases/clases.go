package main

import (
	"fmt"
)

type Persona struct {
	nombre string
	edad   int
	correo string
}

//metodo de clase
func (p *Persona) Imprimir() {
	fmt.Println(p.nombre)
	fmt.Println(p.edad)
	fmt.Println(p.correo)
}


func main() {

	var p Persona
	p.nombre = "Juan"
	p.edad = 21
	p.correo = "kT0wA@example.com"

	p2 := Persona{
		nombre: "Juan2",
		edad:   22,
		correo: "2kT0wA@example.com",
	}

	fmt.Println(p)
	fmt.Println(p2)
	p2.Imprimir()
}