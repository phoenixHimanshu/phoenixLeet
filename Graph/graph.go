package main

import (
	"container/list"
	"fmt"
)

func getVerticesEdgesMatrix() [7][7]int {
	var vertex int
	var edges int
	vertex = 7
	edges = 6
	fmt.Printf("vertex:'%d' ", vertex)
	fmt.Printf("Edges: '%d' ", edges)

	var adjMatrix [7][7]int
	//Multi directional graph
	adjMatrix[0][1] = 1
	adjMatrix[1][0] = 1

	adjMatrix[1][2] = 1
	adjMatrix[2][1] = 1

	adjMatrix[2][3] = 1
	adjMatrix[3][2] = 1

	adjMatrix[3][4] = 1
	adjMatrix[4][3] = 1

	adjMatrix[4][0] = 1
	adjMatrix[0][4] = 1

	adjMatrix[5][6] = 1
	adjMatrix[6][5] = 1

	return adjMatrix

}

//=================== DFS===============
// checking visited node otherwise it will end up in
// an infinite loop between 1 - 0 and 0 - 1

func printGraphDFS(adjMatrix [7][7]int, StartVertex int) {
	v := len(adjMatrix)
	visited := make([]bool, v)
	//This loop is to check edge like 5-6
	for i := 0; i < v; i++ {
		if visited[i] == false {
			printDfsHelper(adjMatrix, visited, StartVertex)
		}
	}

}

func printDfsHelper(adjMat [7][7]int, visit []bool, sv int) {
	println(sv)
	visit[sv] = true
	v := len(adjMat)
	for i := 0; i < v; i++ {
		if adjMat[sv][i] == 1 && visit[i] == false {
			printDfsHelper(adjMat, visit, i)
		}
	}

}

//=================== BFS===============
//The key thing to remember in BFS is the useage of queue
// Iterative solution with usage of stack works best for BFS
func printGraphBFS(adjMatrix [7][7]int) {
	v := len(adjMatrix)
	visited := make([]bool, v)
	//This loop ensures working on edge between vertices 5-6
	for i := 0; i < v; i++ {
		if visited[i] == false {
			printBfsHelper(adjMatrix, visited, i)
		}
	}
}

func printBfsHelper(adjmatrix [7][7]int, visited []bool, sv int) {
	queue := list.New()
	queue.PushBack(sv)
	visited[sv] = true
	numOfVertex := len(adjmatrix)
	//The outer loop below is to traverse through queue
	// ex 0-4 and 4-0 has like , 0 is starting vertex
	// so queue in first iteration would have = 1 ,4 = inside queue
	for 0 < queue.Len() {
		//polling queue ie take out the first element from queue
		front := queue.Front().Value.(int)
		println(front)
		queue.Remove(queue.Front())
		//This inner loop is for checking unvisited edges between
		// vertices
		for i := 0; numOfVertex > i; i++ {
			if adjmatrix[front][i] == 1 && visited[i] == false {
				queue.PushBack(i)
				visited[i] = true
			}
		}
	}
}
