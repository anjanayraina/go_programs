package main

import "fmt"

func main() {
	// using Print
	var i, j = 10, 20.45
	fmt.Print(i, j, "\n") // simply prints the statements

	// using Println
	var x, y = "Anjanay Raina", 20
	fmt.Println(x, y) // prints the statements and adds a \n in the end

	// using Printf
	var l, m = 100, 10
	fmt.Printf("%v and %v are the numbers \n", l, m) // formats the arguments and then prints the statements

	// formatting verbs
	fmt.Printf("The varianble is : %v and the type is %T", x, x)
}
