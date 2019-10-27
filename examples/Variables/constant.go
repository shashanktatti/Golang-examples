package main

import (
	"fmt"
)

func main() {
	/*	PI is declared as const
		1. Const variable should be defined when declared else compiler throws error
		2. When multiplying with radius(int), the expression expects everything to
		   be of the same data type. You basically cannot multiply a float and an
		   int unless int is explicitly converted to float format.
		   So EXPLICIT CONVERSION should be done by the programmer.
		3. When two integers are divided, the result is an integer. So if 22/7 is
		   done, then PI is 3 even if the type of PI is explicitly mentioned as float64
		   So to produce float result, make either numerator or denominator as float.
	*/

	// const PI = 22/7; // gives 3
	
	const PI =  (22/7.0); // gives 3.1412....
	
	var radius int;
	var area float64;

	radius = 5
	area =  (PI * float64(radius) * float64(radius));

	// PI = 3.1412  // gives you error as PI is const

	fmt.Printf("Area of Circle: %f , Value Of PI : %f \n", area, PI);
}