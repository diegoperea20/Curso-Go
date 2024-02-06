package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	t:=time.Now()
	hora:=t.Hour()

	//condicional if else
	if hora < 12 {
		fmt.Println("Buenos dias")
	} else if hora < 17 {
		fmt.Println("Buenas tardes")
	} else {
		fmt.Println("Buenas noches")
	}
	fmt.Println(t)

	//condicional switch
	switch {
		case hora < 12: fmt.Println("Buenos dias")
		case hora < 17: fmt.Println("Buenas tardes")
		default: fmt.Println("Buenas noches")
	}

	//condicional switch otro ejemplo
	os:=runtime.GOOS
	switch os {
		case "windows": 
			fmt.Println("Windows")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Println("Desconocido")
	}


	//for  bucle 
	i:=0
	for i <= 10 {
		fmt.Println(i)
		i++

	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

}