package main

import (
	"fmt"
)

func main() {
	
	diasSemanas := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}

	fmt.Println(diasSemanas[0:3])
	//agregar elementos
	diasSemanas = append(diasSemanas, "Sabado", "Domingo")
	fmt.Println(diasSemanas)
	//eliminar elementos
	diasSemanas = append(diasSemanas[0:5], diasSemanas[6:]...)
	fmt.Println(diasSemanas)

	fmt.Println("longitud",len(diasSemanas))
	fmt.Println("capacidad",cap(diasSemanas))

	nombres:=make([]string, 3, 5) //longitud, capacidad
	nombres[0]="Maria"
	fmt.Println(nombres)

}