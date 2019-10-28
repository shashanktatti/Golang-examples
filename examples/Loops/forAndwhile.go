package main

import (
	"fmt"
)

func main() {
	fmt.Println("Printing 1 to 10 using for loop")

	/*
		No braces in the init and condition.
	*/

	for counter := 1; counter <= 10 ; counter++ {
		fmt.Printf("%d", counter)
		fmt.Printf(" ")
	}

	fmt.Printf("\n")

	/*
		OR
	*/

	fmt.Println("Printing 1 to 10 using for type 2")
	counter := 1		// Implicit type inference
	for ;counter <=10; counter++ {
		fmt.Printf("%d", counter)
		fmt.Printf(" ")
	}

	fmt.Printf("\n")

	/*
		If while loop is to written then for is go's while. There's no while keyword in GO!
	*/

	fmt.Println("Printing 1 to 10 using while with for type 1")
	counter = 1		// Notice that counter is declared before, so := is not required.
	for counter <=10 {
		fmt.Printf("%d", counter)
		fmt.Printf(" ")
		counter++;
	}
}