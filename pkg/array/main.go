package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}

	b := append(a[:2], a[3:]...)

	for _, num := range b {
		fmt.Print(num, ", ")
	}
}
