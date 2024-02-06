package main

import (
	"fmt"
)



func divide(dividend, divisor int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("🤕 no puedes dividir por cero")
	}
}

func main() {
	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(124, 8)
}