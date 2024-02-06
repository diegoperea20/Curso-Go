package main

import (
	"fmt"
)

func Printlist(list ...any){
	for _, value := range list {
		fmt.Println(value)
	}
}

//Crear Restricciones
//https://pkg.go.dev/golang.org/x/exp/constraints
// obtener el paquete con go get golang.org/x/exp/constraints
type misStringInt interface {
	~int | ~string
}




//Restriccion se hacen con parametros de tipo
// es arreglandole el parametro de la funcion func Sum[T int | float64](nums ...T) T 
func Sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

//Antes de no tener la restriccion solos e podia ints
/* func Sum(nums ...int) int{
	var total int
	for _, num := range nums {
		total += num
	}
	return total
} */



//otro ejemplo de restriccion 
func anion[Y misStringInt ](palabra Y){
	fmt.Println(palabra)
}


//CLASE GENERICA
type Product[T ~int | ~string] struct {
	Id T
	Name string
	Price float64
}

//CLASE GENERICA antes 
/* type Product struct {
	Id int
	Name string
	Price float64
} */



func main() {
	fmt.Println(Sum(1,2.7,3,4,5))
	anion(7)
	producto1 := Product[int]{1,"Coca Cola", 3.5}
	producto2 := Product[string]{"AAA","Coca Cola", 3.5}
	fmt.Println(producto1)
	fmt.Println(producto2)
	
		
}