package main

import (
	"fmt"
)

type combination struct {
	indexA int
	indexB int
	result float32
}

func markAllCombinationsUsed(weight []float32, allCombinations []combination, usedCombination combination) {
	weight[usedCombination.indexA] = 0
	weight[usedCombination.indexB] = 0

	for i, cmb := range allCombinations {
		if cmb.indexA == usedCombination.indexA || cmb.indexA == usedCombination.indexB || cmb.indexB == usedCombination.indexA || cmb.indexB == usedCombination.indexB {
			allCombinations[i].indexA = -1
			allCombinations[i].indexB = -1
			allCombinations[i].result = -1
		}
	}
}

func efficientJanitor(weight []float32) int32 {
	// generate all possibles combinations
	combinations := make([]combination, 0)
	for iA, fA := range weight {
		if fA == 0 {
			continue
		}
		for iB, fB := range weight {
			if iA == iB || fB == 0 {
				continue
			}

			if fA+fB <= 3.0 {
				combinations = append(combinations, combination{
					indexA: iA,
					indexB: iB,
					result: fA + fB,
				})
			}
		}
	}

	travels := int32(0)

	var bigger float32
	var index int
	for {
		bigger = float32(0)
		index = -1

		for i, cmb := range combinations {
			if cmb.indexA == -1 {
				continue
			}

			if cmb.result > bigger {
				bigger = cmb.result
				index = i
			}
			if cmb.result == 3.0 {
				break
			}
		}

		if index == -1 {
			break
		}

		travels++
		markAllCombinationsUsed(weight, combinations, combinations[index])
	}

	for _, f := range weight {
		if f != 0 {
			travels++
		}
	}

	return travels
}

func main() {
	weight := []float32{1.01, 1.99, 2.5, 1.5, 1.01}

	fmt.Println(efficientJanitor(weight))
}
