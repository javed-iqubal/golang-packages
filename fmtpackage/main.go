package main

import (
	"fmt"
	"os"
)

// fmt package

/**
1. It print output on console
2. format string
3. read user input
4. build formated string
*/

func fmtPrint() {
	// without new line
	fmt.Print("Hello!!")
	fmt.Print("How are you?")
}

func fmtPrintln() {
	// with new line
	fmt.Println("Hello!!")
	fmt.Println("How are you?")
}

func fmtPrintF() {
	// if you want to format control using varbs
	name := "Javed Iqubal"
	age := 40
	fmt.Printf("%s is %d years old.", name, age)
}

func stringFormatting() string {
	name := "Javed Iqubal"
	str := fmt.Sprintf("Hello %s", name)
	return str
}

func stringConcat() string {
	// return fmt.Sprint("Hello", "Javed", 40)
	return fmt.Sprintln("Hello", "Javed", 40)
}

// getting input from the keyboard
func readInput() (string, int) {
	var name string
	var age int
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s, Age: %d \n", name, age)
	return name, age
}

// write to other outputs (files, network, sockets, buffers etc....., console)

func writeToMany() {
	name := "Javed Iqubal"
	fmt.Fprintf(os.Stdout, "Hello %s", name)
}

func main() {
	//fmtPrint()
	// fmtPrintln()
	// fmtPrintF()
	// fmt.Printf("Formated string is : %s \n", stringFormatting())
	// fmt.Println(stringConcat())
	fmt.Println(readInput())
}
