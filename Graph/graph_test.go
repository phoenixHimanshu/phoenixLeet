package main

import "testing"

/*
	creating graph like this
	 	  0
		/	\
       1 	 4
		\	 /
		2 -  3
					5-6
Note: vertex 5 and 6 have and edge and this is disconnected
graph

*/

func TestGetVerticesEdges(t *testing.T) {

	adjMatrix := getVerticesEdgesMatrix()
	//DFS
	println("=====================")
	printGraphDFS(adjMatrix, 0)

	//BFS

	println("=====================")
	printGraphBFS(adjMatrix)

}
