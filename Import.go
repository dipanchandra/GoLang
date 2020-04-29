// Imports
//--------

// The import keyword is used to declare the packages that are used by the code in the file.

// import example good:
// Let use xecute this program, which intends to receive just 2 input prameters:

package main

import (
	"fmt"
	"os"
	)
	
	
func main () {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Love you",os.Args[1])
}

//Output 1: With proper number of parameters
[nabodip@GyanHouse codes]$ go run import_example.go 3000
Love you 3000

//Output 2: With improper numnber of parameters
[nabodip@GyanHouse codes]$ go run import_example.go 3000 Iron Man
exit status 1

// What if the Println line was:
mt.println("Love you",os.Args[1]) 

//Output:
[nabodip@GyanHouse codes]$ go run import_example.go
# command-line-arguments
./import_example.go:15: cannot refer to unexported name fmt.println
./import_example.go:15: undefined: fmt.println


