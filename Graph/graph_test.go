package main

import "testing"

/*
	creating graph like this
	 	  0
		/	\
       1 	 4
		\	 /
		2 -  3
*/

func TestGetVerticesEdges(t *testing.T) {

	adjMatrix := getVerticesEdgesMatrix()
	printGraph(adjMatrix, 0)

}
