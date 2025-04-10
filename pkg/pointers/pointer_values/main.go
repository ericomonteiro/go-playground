package main

import "fmt"

func main() {
	println("Getting the pointer of a variable")
	name := "John"
	println(name)
	println(&name)
	println(*&name)
	println("")

	println("Creating a pointer of a variable")
	strPointer := new(string)
	println(strPointer)
	println(*strPointer)
	println("")

	println("Creating a pointer of a string")
	*strPointer = "Doe"
	println(strPointer)
	println(*strPointer)
	println("")

	println("Using the pointer inside a map")
	m := make(map[string]*string)
	m["name"] = strPointer
	printMapValues(m)
	println("")

	println("Using the pointer inside a slice")
	s := make([]*string, 0)
	s = append(s, strPointer)
	printSliceValues(s)
	println("")

	println("")
	println("Updating the pointer value")
	*strPointer = "Smith"
	printMapValues(m)
	printSliceValues(s)

}

func printMapValues(m map[string]*string) {
	println("Printing map values")
	for k, v := range m {
		println(fmt.Sprintf("Key: %s, Value: %s", k, *v))
	}
}

func printSliceValues(s []*string) {
	println("Printing slice values")
	for i, v := range s {
		println(fmt.Sprintf("Index: %d, Value: %s", i, *v))
	}
}
