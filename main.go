package main

import (
	"container/list"
	// "reflect"
	"fmt"
)

type edge [2]int 

func equals(e1 edge, e2 edge) bool {
	if (e1[0] == e2[0]) && 
		(e1[1] == e2[1]){
		return true
	} else if (e1[0] == e2[1]) && 
			  (e1[1] == e2[0]){
		return true
	}
	return false
}

func makeEdgeMatrix(n int, edges [][]int) []list.List {
	edgeMatrix := make([](list.List), n)
	for i := 0; i < n; i++ {
		edgeMatrix[i] = *list.New()
	}
	for _, e := range edges {
		fmt.Printf("%v, %T\n", e, e)
		edgeMatrix[e[1]].PushBack(e[0])
		edgeMatrix[e[0]].PushBack(e[1])

	}
	return edgeMatrix
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	if (n == 0) {
		return 0
	}

	// edgematrix := makeEdgeMatrix(n, edges)
	
    return 0
}

func main() {
	fmt.Println("hello world")
	n := 7
	edges := [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}
	hasApple := []bool{false,false,true,false,true,true,false}
	minTime(n, edges, hasApple)
}