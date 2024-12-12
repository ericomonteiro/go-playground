package main

import "fmt"

type Customer struct {
	Name string
}

func (c Customer) UpdateNameByValue(name string) {
	c.Name = name
}

func (c *Customer) UpdateNameByPointer(name string) {
	c.Name = name
}

func main() {
	// Receive by value create a copy of the object
	fmt.Println("Case 1 - Receive by value create a copy of the object")
	case1()
	fmt.Println("")

	// Receive by pointer receive the reference of the object
	fmt.Println("Case 2 - Receive by pointer receive the reference of the object")
	case2()
	fmt.Println("")

	// Update the object by value doesn't change the original object
	fmt.Println("Case 3 - Update the object by value doesn't change the original object")
	case3()
	fmt.Println("")

	// Update the object by pointer change the original object
	fmt.Println("Case 4 - Update the object by pointer change the original object")
	case4()
	fmt.Println("")
}

func case1() {
	c1 := Customer{Name: "John"}
	c2 := c1

	c2.Name = "Doe"

	fmt.Println(c1.Name)
	fmt.Println(c2.Name)

	// c1 and c2 are different objects
	fmt.Println(c1 == c2)
}

func case2() {
	c1 := &Customer{Name: "John"}
	c2 := c1

	c2.Name = "Doe"

	fmt.Println(c1.Name)
	fmt.Println(c2.Name)

	// c1 and c2 are the same object (the same value of the pointer)
	fmt.Println(c1 == c2)
}

func case3() {
	c := Customer{Name: "John"}
	c.UpdateNameByValue("Doe")

	fmt.Println(c.Name) // Doe
}

func case4() {
	c := &Customer{Name: "John"}
	c.UpdateNameByPointer("Doe")

	fmt.Println(c.Name) // John
}
