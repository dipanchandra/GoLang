// Pointer_1
// A pointer is a memory adress whose location stores the actual value.

package main

type Voter struct {
Name string
Age int
}

func main () {
	citizen := Voter{"It is you", 23}
	Senior_Citizen(citizen)
	println(citizen.Age)
}

func Senior_Citizen(s Voter) {
s.Age += 60
}


// Output:
[nabodip@GyanHouse codes]$ go run pointer_1.go
23

// There is no modification in the value of "Age" as the function Senior_Citizen made changes to a copy of original values of 
"citizen" and hence changes made in "Age" were not reflected in the caller.

// In order to get this modification into effect, the pointer to the value needs to be passed.
package main

type Voter struct {
Name string
Age int
}

func main () {
	citizen := &Voter{"It is you", 23}
	Senior_Citizen(citizen)
	println(citizen.Age)
}

func Senior_Citizen(s *Voter) {
s.Age += 60
}


[nabodip@GyanHouse codes]$ go run pointer_1.go
83

// The modifications done in this code were:
//  1.Use of "&" opertor to get adresss of the value(i.e. adress of operator)
//  2.Changing the type of parameters which the function "Senior_Citizen" expects, we are now supplying it with
//    a "pointer to value of type Voter".

