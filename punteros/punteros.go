package main

import (
	"fmt"
)

func main() {
	var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)
}

func editar(x *int) {
	*x = 20
}