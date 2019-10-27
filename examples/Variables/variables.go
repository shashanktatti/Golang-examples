package main

import (
	"fmt"
)

func main() {
	/*	var var_name var_type */
	var name string;
	var age int;
	var weight float32;

	/*  All Local Variables are initialized with default values.
		For String "" , int and float it's 0.
	*/

	/*
		fmt.Printf uses format specifiers like %s and %d , %f. But Println does not
		use any format specifiers.
	*/
	fmt.Printf("Name = %s, Age = %d, Weight = %f \n", name, age, weight);
	fmt.Println("Name : , Age: , Weight: ", name, age, weight)

	/*
		Dynamic Type Declaration or Type Inference is a mechanism in which Go compiler automatically infers the 
		type of the variable from the defined value.
		Second type is when defined at the time of declaration, no need to specify the type of the variable.
		And multiple variables can be of any type and can be defined simulataneously. 
	*/

	noTypeDeclaration := true
	var first, second, third = 1, 2, 3

	fmt.Printf("Type Inference : %t, Definition w/o any type specification: %d : %d : %d \n", 
	noTypeDeclaration, first, second, third)

}