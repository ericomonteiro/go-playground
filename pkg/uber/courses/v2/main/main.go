package main

/*

Algorithm to check if the course list is valid.

It will receive a list of course pairs [][]int, where the first element of the pair is the course
and the second element is the course that the first depends on.

Example: [0, 1] means that course 0 depends on course 1.
List: [[0, 1] [1, 2] [1, 4] [2, 3] [2, 4]]
The course list is valid if there are no cycles, that is, if there is no course that indirectly depends on itself.
*/

func dfs(graph [][]int, visited []bool, completed []bool, node int) bool {
	if visited[node] && !completed[node] {
		return false
	}

	visited[node] = true

	for i := 0; i < len(graph); i++ {
		if graph[node][i] == 1 {
			if !dfs(graph, visited, completed, i) {
				return false
			}
		}
	}

	completed[node] = true
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

func checkGradeIsValid(courses [][]int) bool {
	var graph [][]int
	var result bool
	var visited []bool
	var completed []bool

	graph = buildGraph(courses)
	printGraph(graph)
	visited = make([]bool, len(graph))
	completed = make([]bool, len(graph))
	result = dfs(graph, visited, completed, 0)
	println(result)
	println("\n--------------------------")
	return result
}
