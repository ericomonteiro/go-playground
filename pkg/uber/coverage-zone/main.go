package main

import (
	"fmt"
)

type Graph = map[Coordinate][]Coordinate

type Coordinate struct {
	X int
	Y int
}

func isPossibleToConnect(graph Graph, pointA, pointB Coordinate) bool {
	if _, exists := graph[pointA]; !exists {
		return false
	}

	if _, exists := graph[pointB]; !exists {
		return false
	}

	visited := make(map[Coordinate]bool)
	var fnDfs func(pointA, pointB Coordinate) bool

	var path string
	fnDfs = func(pointA, pointB Coordinate) bool {
		visited[pointA] = true
		for _, edge := range graph[pointA] {
			if visited[edge] {
				continue
			}

			if edge == pointB {
				path = fmt.Sprintf("[%d %d]", edge.X, edge.Y) + path
				return true
			}

			if fnDfs(edge, pointB) {
				path = fmt.Sprintf("[%d %d] --> ", edge.X, edge.Y) + path
				return true
			}
		}

		return false
	}

	result := fnDfs(pointA, pointB)

	fmt.Println()
	fmt.Println("Result.:")
	if result {
		fmt.Println("Is possible to connect")
		path = fmt.Sprintf("[%d %d] --> ", pointA.X, pointA.Y) + path
		fmt.Println(path)
	} else {
		fmt.Println("No path found")
	}

	return result
}

func buildGraph(coverageZone [][]int) Graph {
	coordinatorMin := Coordinate{
		X: coverageZone[0][0],
		Y: coverageZone[0][1],
	}

	coordinatorMax := Coordinate{
		X: coverageZone[0][0],
		Y: coverageZone[0][1],
	}

	// Create empty graph
	graph := make(Graph)
	for _, coordinate := range coverageZone {
		node := Coordinate{
			X: coordinate[0],
			Y: coordinate[1],
		}

		if node.X < coordinatorMin.X {
			coordinatorMin.X = node.X
		}

		if node.Y < coordinatorMin.Y {
			coordinatorMin.Y = node.Y
		}

		if node.X > coordinatorMax.X {
			coordinatorMax.X = node.X
		}

		if node.Y > coordinatorMax.Y {
			coordinatorMax.Y = node.Y
		}

		graph[node] = []Coordinate{}
	}

	for node, _ := range graph {
		edges := possibleEdges(node, coordinatorMin, coordinatorMax)

		for _, edge := range edges {
			if _, exists := graph[edge]; exists {
				graph[node] = append(graph[node], edge)
			}
		}
	}

	printGraph(graph)

	return graph
}

func possibleEdges(node Coordinate, coordinateMin, coordinateMax Coordinate) []Coordinate {
	adjacent := make([]Coordinate, 0)
	if node.X > coordinateMin.X {
		adjacent = append(adjacent, Coordinate{node.X - 1, node.Y})
	}
	if node.X < coordinateMax.X {
		adjacent = append(adjacent, Coordinate{node.X + 1, node.Y})
	}

	if node.Y > coordinateMin.Y {
		adjacent = append(adjacent, Coordinate{node.X, node.Y - 1})
	}
	if node.Y < coordinateMax.Y {
		adjacent = append(adjacent, Coordinate{node.X, node.Y + 1})
	}

	return adjacent
}

func printGraph(graph Graph) {
	for node, edges := range graph {
		fmt.Printf("Node: [%d, %d]\n", node.X, node.Y)
		fmt.Printf("   Edges: ")
		for _, edge := range edges {
			fmt.Printf("[%d, %d] ", edge.X, edge.Y)
		}
		fmt.Println()
	}
}

func main() {
	var coverageZone [][]int

	pointA := Coordinate{0, 0}
	pointB := Coordinate{3, 0}

	coverageZone = [][]int{
		{0, 0},
		{0, 1},
		{0, 2},

		{1, 2},

		{2, 0},
		{2, 1},
		{2, 2},
		{2, 3},

		{3, 0},
		{3, 3},
	}

	graph := buildGraph(coverageZone)
	_ = isPossibleToConnect(graph, pointA, pointB)
}
