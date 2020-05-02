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
// The compiler complains out with the above error. Note that when the variable is declared first, that time:= is used, but on 
// subsequent assignments, the operator = is used. 


// Q. Why does it says : " no new variable[s] on left side of := " ?
// A. Reason being that Go lets us assign multiple variables using either "=" or ":=".


// Declaring multiple variables in same line:
// Program Name : run var_and_dec_same_line.go
package main

import (
	"fmt"
)

func main(){
	country,year := "India", 1947
	fmt.Printf("%s got Independence in the year %d.\n", country,year)
}

//Output:
[nabodip@GyanHouse codes]$ go run var_and_dec_same_line.go 
India got Independence in the year 1947.


// := works well while returning a value from a function as well:
// Program Name : var_and_dec_fun.go

package main

import (
	"fmt"
)

func main(){
	age := getAge()
	fmt.Printf("My age is %d years.\n", age)
}

func getAge() int{
	return 23
}

//Output:
[nabodip@GyanHouse codes]$ go run var_and_dec_fun.go
My age is 23 years.



// Declaring a variablebut not using it in code :
// Program Name : var_and_dec_same_line_dec_not_use.go

package main

import (
	"fmt"
)

func main(){
	country,year := "India", 1947
	fmt.Printf("%s got Independence.\n", country)
}

//Output:
[nabodip@GyanHouse codes]$ go run  var_and_dec_same_line_dec_not_use.go
# command-line-arguments
./var_and_dec_same_line_dec_not_use.go:8: year declared and not used

//Program did not compile, because the variable "year" was declared, but was not used later in the program.


//So, final takeaways:

// For declaring avariable to its zero value:
	var NAME TYPE
// For declaring and assigning a specific value:
	NAME := VALUE
// For assigning new value to a previously declared variable:
	NAME = VALUE


