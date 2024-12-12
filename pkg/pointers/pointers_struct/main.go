package main

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
	println("Case 1 - Receive by value create a copy of the object")
	case1()
	println("")

	// Receive by pointer receive the reference of the object
	println("Case 2 - Receive by pointer receive the reference of the object")
	case2()
	println("")

	// Update the object by value doesn't change the original object
	println("Case 3 - Update the object by value doesn't change the original object")
	case3()
	println("")

	// Update the object by pointer change the original object
	println("Case 4 - Update the object by pointer change the original object")
	case4()
	println("")
}

func case1() {
	c1 := Customer{Name: "John"}
	c2 := c1

	c2.Name = "Doe"

	println(c1.Name)
	println(c2.Name)

	// c1 and c2 are different objects
	println(c1 == c2)
}

func case2() {
	c1 := &Customer{Name: "John"}
	c2 := c1

	c2.Name = "Doe"

	println(c1.Name)
	println(c2.Name)

	// c1 and c2 are the same object (the same value of the pointer)
	println(c1 == c2)
}

func case3() {
	c := Customer{Name: "John"}
	c.UpdateNameByValue("Doe")

	println(c.Name) // Doe
}

func case4() {
	c := &Customer{Name: "John"}
	c.UpdateNameByPointer("Doe")

	println(c.Name) // John
}
