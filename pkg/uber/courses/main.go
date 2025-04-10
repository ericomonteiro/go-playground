package main

import (
	"fmt"
)

type graph struct {
	Vertices  int
	Inbounds  map[int][]int
	Outbounds map[int][]int
}

func (g *graph) addEdge(from, to int) {
	if g.Inbounds[from] == nil {
		g.Inbounds[from] = []int{}
	}
	if g.Outbounds[to] == nil {
		g.Outbounds[to] = []int{}
	}
	g.Inbounds[to] = append(g.Inbounds[to], from)
	g.Outbounds[from] = append(g.Outbounds[from], to)
}

func (g *graph) print() {
	for i := range g.Vertices {
		fmt.Println(i, "->", g.Inbounds[i])
		fmt.Println(i, "<-", g.Outbounds[i])
	}
}

func main() {
	var courses [][]int
	var numCourses int

	courses = [][]int{
		{0, 1},
		{1, 2},
		{1, 4},
		{2, 3},
		{3, 4},
	}
	numCourses = 5
	fmt.Println(isPossibleToFinish(courses, numCourses))

	//courses = [][]int{
	//	{0, 1},
	//	{1, 0},
	//}
	//numCourses = 2
	//fmt.Println(isPossibleToFinish(courses, numCourses))

	//courses = [][]int{
	//	{0, 1},
	//	{1, 0},
	//}
	//numCourses = 2
	//fmt.Println(isPossibleToFinish(courses, numCourses))

}

func isPossibleToFinish(courses [][]int, numCourses int) bool {
	g := &graph{
		Vertices:  numCourses,
		Inbounds:  make(map[int][]int),
		Outbounds: make(map[int][]int),
	}

	for _, course := range courses {
		g.addEdge(course[0], course[1])
	}

	currentPathOfVertices := make(map[int]bool)
	visited := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(v int) bool {
		if currentPathOfVertices[v] {
			return true
		}
		if visited[v] {
			return false
		}

		currentPathOfVertices[v] = true
		for _, neighbor := range g.Outbounds[v] {
			if dfs(neighbor) {
				return true
			}
		}
		currentPathOfVertices[v] = false
		visited[v] = true

		return false
	}

	result := !dfs(0)

	for k, _ := range visited {
		println(k)
	}

	return result
}
