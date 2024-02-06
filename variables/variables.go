package main

//importar el paquete 
import (
	"fmt"
	//"rsc.io/quote"
	"strconv"
)

//declarar una variable constante
const CONts float32 = 65.8

//declarar una variable
var otravariable string = "otra variable"

//forma iota , cuando se pone iota  es igual 0 , una variable que se incrementa
const (
	Lunes = iota + 1
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)
func main(){
	//declarar una variable
	var name string = "Juan"
	var firstName , lastName string 
	var age int

	//otra forma de declarar una variable
	var(
		name2 , firstName2 = "Jorge","Juan"
		lastName2 = "Hernandez"
		age2 = 34
	)

	//otra forma de declarar una variable
	var name3 , firstName3 , lastName3 , age3 = "Jorge","Juan","Hernandez",36

	//otra forma de declarar una variable
	name4, firstName4 := "robin" , "david"
	numeroFloat := 12.3

	firstName = "Jorge"
	lastName = "Hernandez"
	age = 25
	nuevaVariable := "otra nueva variable" //se tiene en entedido que := es una forma de decir que se esta declarando una variable corta y con el = es para modificar

	fmt.Println(name)
	fmt.Println(firstName,lastName,age)
	fmt.Println(name2,firstName2,lastName2,age2)
	fmt.Println(name3,firstName3,lastName3,age3)
	fmt.Println(name4,firstName4)
	fmt.Println(otravariable)
	fmt.Println("el valor de Viernes es: ",Viernes)
	fmt.Println(nuevaVariable)
	// Using Sprintf and printing the result
	result := fmt.Sprintf("lunes es %d y martes es %d ", Lunes, Martes)
	fmt.Println(result)
	fmt.Println("----Conversion de tipos----")
	/* strconv.Itoa() se utiliza para convertir de int a string.
    strconv.Atoi() se utiliza para convertir de string a int. Ten en cuenta que  */
	numberAsString := strconv.Itoa(age)
	fmt.Println("de int a string",numberAsString)
	StringAsInt , _ := strconv.Atoi(numberAsString)
	fmt.Println("de string a int",StringAsInt)
	fmt.Println("de float a int",int(numeroFloat))
	fmt.Println("de int a float",float32(12))

	fmt.Println("----Inputs")
	var entrada string
	fmt.Print("ingrese un valor string : ")
	fmt.Scanln(&entrada)
	fmt.Println("la entrada fue: ",entrada)
	fmt.Printf("el tipo de dato de entrada es: %T\n", entrada)
	

} 

