package main

import (
	
	"testing"
)

//el nombre del archivo test siempre debe ser nombre_test 
// todos los test empiezan con Test
func TestSuma(t *testing.T) {
	
	if suma(2, 3) != 5 {
		t.Error("La suma no es 5")
	}
}