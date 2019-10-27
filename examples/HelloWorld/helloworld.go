/*	Every go program starts with a package	*/
package main

/*	fmt is equivalent to STDIO in C */
import "fmt"

/*
	Every function func keyword and then func_name 
	Rule in GO that every function exposed to outer world (public access specifier)
	would start with a capital letter. Every function that is private would start with a small letter.
	Println starts with caps as it's a public function.
*/
func main() {
	fmt.Println("Hello World!")
}