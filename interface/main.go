
package main

import (
	"dependenciasProyecto/interface/animal"
)

func main() {
	
	miPerro := animal.Perro{Nombre: "Max"}
	miGato := animal.Gato{Nombre: "Tom"}

	animal.HacerSonido(miPerro) // Imprime "Max hace guau guau"
	animal.HacerSonido(miGato)  // Imprime "Tom hace miau"
	
	
}
