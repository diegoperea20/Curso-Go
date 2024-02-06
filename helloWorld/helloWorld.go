package main

//importar el paquete 
import (
	"fmt"
	"rsc.io/quote"

)



func main(){
	fmt.Println("Hello, World! Go is great!")
	println("Hello, World! Go is great form normal!")
	//Se usa el paquete rsc.io/quote
	fmt.Println(quote.Go())
}

// para guardar dependencias: go mod init NombreDelArchivo

//para instalar paquetes de go (Se añade el paquete go.mod ): go get NombreDelPaquete

// ttodo se debe usar o da error