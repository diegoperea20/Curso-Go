package main	

import (
	"fmt"
	"strconv"
)


//creacion del control del error personalizado
func divide(dividiendo , divisor int)(int , error){

	if divisor == 0{
		return 0 , fmt.Errorf("no se puede dividir por 0")
	}
	return dividiendo / divisor , nil
}

func main(){
	 str:= "10"
	 num , err := strconv.Atoi(str)
	 if err != nil{
	 fmt.Println("Error" , err)
	 }else{
	 fmt.Println(num)
	 }

	 //usando el control de error personalizado
	 resultado , err := divide(10,0)
	 if err != nil{
	 fmt.Println("Error" , err)
	 }else{
	 fmt.Println(resultado)
	 }


}