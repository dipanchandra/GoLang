// A simple example would be:

// Example #1:
package main

import (
	"fmt"
)

func main () {
	var age int
	//no value defined for variable "age"
  fmt.Printf("My age is %d years.\n", age)
}

// Output:
[nabodip@GyanHouse codes]$ go run var_and_dec.go
My age is 0 years.

// == By default, Go assigns a zero value to variable.

// Example #2:
func main () {
	var age int
	age = 23
	fmt.Printf("My age is %d years.\n", age)
}


// Output:
[nabodip@GyanHouse codes]$ go run var_and_dec.go
My age is 23 years.

// Example #3:
// The "short variable declaration operator :=" is used to declare as well as assign values 
// So,while the following is correct :

func main () {
	//var age int
	age := 23
	fmt.Printf("My age is %d years.\n", age)
}

// Output:
[nabodip@GyanHouse codes]$ go run var_and_dec.go
My age is 23 years.


// But the following will generate an error, asa variable is being declared twice in same scope:

func main () {
	var age int
	age := 23
	fmt.Printf("My age is %d years.\n", age)
}

// Output:
[nabodip@GyanHouse codes]$ go run var_and_dec.go
# command-line-arguments
./var_and_dec.go:9: no new variables on left side of :=

