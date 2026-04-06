package main

import (
	"fmt"

	"example.com/my-package/calculator"

	"example.com/my-package/util"
)

func main() {

	fmt.Println(calculator.Myapp)

	result := calculator.Add(30, 20)
	fmt.Printf("Result Add : %d \n", result)

	result = calculator.Substract(30, 20)
	fmt.Printf("Result Substract : %d \n", result)

	result = calculator.Multiply(30, 20)
	fmt.Printf("Result Multiply : %d \n", result)

	result2 := calculator.Divide(30, 20)
	fmt.Printf("Result Divide : %f \n", result2)

	fmt.Println(util.GetInfo())
}
