package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}


	// Crear un canal , para ver los resultados en paralelo
	ch:= make(chan string)

	// Recorreer los apis
	for _, api := range apis {
		go checkAPI(api , ch)  //escribir go en la funcion hace que se ejecute en paralelo
	}
	//leer del canal
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	} 

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())
}

// Función que verifica los APIS
func checkAPI(api string , ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s está caído!\n", api)  //para enviarlo al canal
		return
	}

	ch <- fmt.Sprintf("SUCCESS: ¡%s está en funcionamiento!\n", api)//para enviarlo al canal
}