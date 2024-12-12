package main

import "fmt"

/*
When we create a slice, we are creating a reference to an array. The slice is a structure that contains:
- A pointer to the array
- The length of the slice
- The capacity of the slice
*/

func main() {
	fmt.Println("Trying to append a value to a slice")
	stringSlice := []string{"John", "Doe", "Smith"}
	fmt.Println("Before: ", stringSlice)
	tryToAppend(stringSlice)
	fmt.Println("After: ", stringSlice)
	fmt.Println("")

	fmt.Println("Updating a value inside a slice")
	stringSlice = []string{"John", "Doe", "Smith"}
	fmt.Println("Before: ", stringSlice)
	updateValueInsideSlice(stringSlice)
	fmt.Println("After: ", stringSlice)
}

func tryToAppend(s []string) {
	s = append(s, "Jane")
	s[0] = "Updated?"
	fmt.Println("Inside the function: ", s)
}

func updateValueInsideSlice(s []string) {
	s[0] = "Jane"
	s[1] = "Jane"
	s[2] = "Jane"
	fmt.Println("Inside the function: ", s)
}
