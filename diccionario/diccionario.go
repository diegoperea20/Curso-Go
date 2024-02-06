package main

import (
	"fmt"
)

func main() {
	//los diccionarios se ordenann por orden alfabetico
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors["green"])
	//agregar
	colors["white"] = "#ffffff"
	fmt.Println(colors)
	//borrar
	delete(colors, "white")
	fmt.Println(colors)

	//verificar valor
	valor,ok := colors["red"]
	if ok {
		fmt.Println("El valor es:", valor)
	}else{
		fmt.Println("No existe")
	}

	//recorrer un diccionario
	for clave , valor := range colors{
		fmt.Println(clave,valor)
	}
	
}