package main

var courses [][]int

/*
0-> 0 1 0 0 0
1-> 0 0 1 0 1
2-> 0 0 0 1 1
3-> 0 0 0 0 0
4-> 0 0 0 0 0

*/

func dfs(graph [][]int, visited []bool, node int) bool {
	if visited[node] {
		return false
	}
	visited[node] = true

	for i := 0; i < len(graph); i++ {
		if graph[node][i] == 1 {
			if !dfs(graph, visited, i) {
				return false
			}
		}
	}

	return true
}

func buildGraph(courses [][]int) [][]int {
	vertices := map[int]bool{}
	for _, course := range courses {
		vertices[course[0]] = true
		vertices[course[1]] = true
	}

	graph := make([][]int, len(vertices))
	for v, _ := range graph {
		graph[v] = make([]int, len(vertices))
	}

	for _, courseDeps := range courses {
		graph[courseDeps[0]][courseDeps[1]] = 1
	}

	return graph
}

func printGraph(graph [][]int) {
	for i := range graph {
		print(i, "-> ")
		for j := range graph[i] {
			print(graph[i][j], " ")
		}
		println()
	}

	println("------")
}

func main() {
	var graph [][]int
	var result bool
	var visited []bool

	//valid
	courses = [][]int{
		{0, 1},
		{1, 2},
		{1, 4},
		{2, 3},
		{2, 4},
	}
	graph = buildGraph(courses)
	printGraph(graph)
	visited = make([]bool, len(graph))
	result = dfs(graph, visited, 0)
	println(result)
	println("\n--------------------------")

	//invalid
	courses = [][]int{
		{0, 1},
		{1, 2},
		{1, 4},
		{2, 3},
		{3, 1},
	}
	graph = buildGraph(courses)
	printGraph(graph)
	visited = make([]bool, len(graph))
	result = dfs(graph, visited, 0)
	println(result)
	println("\n--------------------------")

}
