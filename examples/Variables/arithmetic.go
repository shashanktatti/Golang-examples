package main

import (
	"fmt"
	"math"
)

func main() {
	/*	
		Integer is of two types. int32 and int64.
	*/
	var integer32 int32 = 20;
	var integer64 int64 = 40;

	/*
		Explicit type conversion needs to be made by programmer whenever necessary.
	*/

	fmt.Printf("Type of int32 and int64 : %T %T \n", integer32, integer64)
	fmt.Printf("Adding int 32 and int 64 : %d\n", int64(integer32) + integer64)
	fmt.Printf("Subtracting int 32 and int 64 : %d\n", int64(integer32) - integer64)
	fmt.Printf("Multiplying int 32 and int 64 : %d\n", int64(integer32) * integer64)
	fmt.Printf("Dividing int 32 and int 64 : %d\n", int64(integer32) / integer64)

	/*
		Math package

		Sin function takes float64 as an argument. So int64/int32 needs to be converted to float64.
		Same for Abs also.
	*/

	fmt.Printf("Sin of %d is %f\n", integer64, math.Sin(float64(integer64)))
	fmt.Printf("Absolute value of %d - %d is : %f\n", integer32, integer64, 
	math.Abs(float64(int64(integer32) - integer64)))
}