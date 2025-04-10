package main

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	numberOfComponents := 0
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true
		for i := 0; i < n; i++ {
			if isConnected[node][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			numberOfComponents++
			dfs(i)
		}
	}

	return numberOfComponents
}

func main() {
	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	println(findCircleNum(isConnected))
}
