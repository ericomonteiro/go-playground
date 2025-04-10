package main

func findCircleNum(isConnected [][]int) int {
	// create a province for each city
	provinces := map[int][]int{}
	for i, _ := range isConnected {
		provinces[i] = []int{i}
	}

	for baseCity, connections := range isConnected {
		baseProvince := findCityProvince(provinces, baseCity)
		for connectedCity, connected := range connections {
			//ignore cases:
			// - current city (diagonal)
			// - isn't connected
			// - already processed
			if baseCity == connectedCity || connected == 0 {
				continue
			}

			provinceConnectedCity := findCityProvince(provinces, connectedCity)
			if provinceConnectedCity == baseProvince {
				continue
			}

			mergeProvinces(provinces, baseProvince, provinceConnectedCity)

			isConnected[connectedCity][baseCity] = 0
		}
	}

	return len(provinces)
}

func findCityProvince(provinces map[int][]int, city int) int {
	for province, cities := range provinces {
		for _, c := range cities {
			if c == city {
				return province
			}
		}
	}

	return -1
}

func mergeProvinces(provinces map[int][]int, province1, province2 int) {
	provinces[province1] = append(provinces[province1], provinces[province2]...)
	delete(provinces, province2)
}

func main() {
	//[[1,1,0],[1,1,0],[0,0,1]]
	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 1},
	}
	println(findCircleNum(isConnected))
}
