package main

import (
	"fmt"
	"strconv"
)

/**
Type conversion means converting a value from one data type to another data type
Go does not support automatic(implicit) type conversion.
Conversion Syntax:
	TargetType(value)
eg:
	float64(10)
	int(3.45)
	string(123456)
*/

func main() {

	// int to float
	var a int = 10
	var b float32 = float32(a)
	fmt.Printf("a: %d, b: %f \n", a, b)

	// float to int
	var x float32 = 6.456
	var y int = int(x)
	fmt.Printf("x: %f, y: %d \n", x, y)

	// number to string conversion using rune
	num := 65
	str := string(num)
	fmt.Printf("num: %d, str: %s \n", num, str)

	// string to number
	cost := "1456"
	// int can't be used to convert string to number
	// strconv
	var c, _ = strconv.Atoi(cost)
	fmt.Printf("cost: %s, converted cost: %d \n", cost, c)
}
