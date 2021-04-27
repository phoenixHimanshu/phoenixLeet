package main

import "fmt"

func getVerticesEdgesMatrix() [5][5]int {
	var vertex int
	var edges int
	vertex = 5
	edges = 5
	fmt.Printf("vertex:'%d' ", vertex)
	fmt.Printf("Edges: '%d' ", edges)

	var adjMatrix [5][5]int

	adjMatrix[0][1] = 1
	adjMatrix[1][2] = 1
	adjMatrix[2][3] = 1
	adjMatrix[3][4] = 1
	adjMatrix[4][0] = 1

	return adjMatrix

}

func printGraph(adjMatrix [5][5]int, StartVertex int) {
	v := len(adjMatrix)
	visited := make([]bool, v)
	for i := 0; i < v; i++ {
		if visited[i] == false {
			printHelper(adjMatrix, visited, StartVertex)
		}
	}

}

func printHelper(adjMat [5][5]int, visit []bool, sv int) {
	fmt.Printf("Starting Vertex = %d", sv)
	visit[sv] = true
	v := len(adjMat)
	for i := 0; i < v; i++ {
		if adjMat[sv][i] == 1 && visit[i] == false {
			printHelper(adjMat, visit, i)
		}
	}

}
