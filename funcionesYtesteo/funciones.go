package main

import (
	"fmt"
)

func main() {
	hello("Juan")
	resultado := suma(2, 3)
	fmt.Println("El resultado es: ", resultado)

	sum, mul := calc(2, 3)
	fmt.Println("La suma es: ", sum, " y la multiplicación es: ", mul)

	fmt.Println("Funcion Variadica",SumaLargas(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("Funcion Recursiva",factorial(5))
	fmt.Println("Funcion Orden Superior",higherOrderFunction(addOne, 3))
}

func suma(a int, b int) int {  //lo que retorna la función es de tipo int
	return a + b
}

func hello(name string){
	fmt.Println("Hello, ", name)
}

func calc(a,b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

//FUNCION VARIADICA
func SumaLargas(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//FUNCION RECURSIVA
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}


//FUNCION ORDEN SUPERIOR
func higherOrderFunction(f func(int) int, a int) int {
	return f(a * 2)
}

func addOne(a int) int {
	return a + 1
}

