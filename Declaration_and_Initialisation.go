// Basic way to create a structure:

citizen := Voter{
  Name := "You",
  Age  := 23,
}

// IMP: The trailing "," is required in order to avoid a Compiler error. Goenforces this behaviour.

// It is not necessary to set all or even any of the fields. e.g.

// Valid example # 1 -

citizen := Voter{}

// or

citizen := Voter{Name := "You"}
citizen.Age = 23


// Unassigned fields also have zero value.

// For Structure with new fields, the field name can be skipped while maintaining oder of field definitions:

citizen := Voter{"You",23}

// all above examples are declaring a variable "citizen", and assigning a value to it.


  
