// Functions can return multiple values

// Sample # 1 -
func greetings(message string) {
}
// This function does not returns any value


// Sample # 2 -
func add(a int,b int) int {
}
// This function returns one value


// Sample # 3 -
func voter(name string) (int, bool) {
}
// This function returns Two values

// Sample # 4 -
func add(x, y int) int {
}
// In case the parameters share the same data type, the syntax can be shortened by non repetetive mention of the data type.

// Usage :
person_name, is_a_voter := voter("You")
if is_a_voter == true {
  // corresponding code
}

// What if we are interested in only one return variable?
// in that case,the other values can be assigned to _ :

_, is_a_voter := voter("You")
if is_a_voter == true {
  // corresponding code
}


// Q. What is _ ?
// A. 
// 1.) _ is a Blank identifier whose speciality is that its return value is not assigned,
// 2.) _ can be thus repeatedly used over and over again irrespective of the return type.
// 3.) _ is thus basically used to discard a value.




