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

func printGraphDFS(adjMatrix [7][7]int, StartVertex int) {
	v := len(adjMatrix)
	visited := make([]bool, v)
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

func printGraphBFS(adjMatrix [7][7]int) {
	v := len(adjMatrix)
	visited := make([]bool, v)
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

	for 0 < queue.Len() {
		//polling queue
		front := queue.Front().Value.(int)
		println(front)
		queue.Remove(queue.Front())

		for i := 0; numOfVertex > i; i++ {
			//queue.PushBack(i)
			if adjmatrix[front][i] == 1 && visited[i] == false {
				//println(i)
				queue.PushBack(i)
				visited[i] = true
			}
		}
	}
}
