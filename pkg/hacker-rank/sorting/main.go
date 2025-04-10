package main

import (
	"fmt"
	"slices"
)

func calculateModule3(n int) int {
	return n % 3
}

func RemainderSorting(strArr []string) []string {
	sorting := make([]string, len(strArr))
	for i, str := range strArr {
		sorting[i] = fmt.Sprintf("%v%s", calculateModule3(len(str)), str)
	}

	slices.Sort(sorting)

	for i, str := range sorting {
		sorting[i] = str[1:]
	}

	return sorting
}

func main() {
	fmt.Println(RemainderSorting([]string{"Colorado", "Utah", "Wisconsin", "Oregon"}))
}
