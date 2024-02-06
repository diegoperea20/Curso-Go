package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}

func jugar() {
	fmt.Println("Jugando")
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10
	for intentos < maxIntentos{
		intentos++
		fmt.Printf("Ingresa un numero (Intentos restantes:%d): ",  maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Println("Ganaste")
			jugarDeNuevo()
			return
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es menor")
		} else {
			fmt.Println("El numero a adivinar es mayor")
		}
	}

	fmt.Println("Perdiste, el numero era: ", numAleatorio)
	jugarDeNuevo()
}

func jugarDeNuevo() {
	var eleccion string
	fmt.Println("Deseas jugar de nuevo? (si/no)")
	fmt.Scanln(&eleccion)

	if eleccion == "si" {
		jugar()
	}else if eleccion == "no" {
		fmt.Println("Gracias por jugar")
	}else {
		fmt.Println("Opcion invalida")
		jugarDeNuevo()
	}
}