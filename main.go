package main

import (
	"container/list"
	"math"
	"strconv"
	"fmt"
)

type edge [2]int 

type node struct {
	value int
	children []node
	parent *node
}

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

func makeEdgeMatrix(n int, edges [][]int) []*list.List {
	edgeMatrix := make([](*list.List), n)
	for i := 0; i < n; i++ {
		edgeMatrix[i] = list.New()
	}
	for _, e := range edges {
		edgeMatrix[e[1]].PushBack(e[0])
		edgeMatrix[e[0]].PushBack(e[1])

	}
	return edgeMatrix
}

func isVisited(visited map[int]struct{}, node interface{}) bool {
	_, ok := visited[node.(int)]
	return ok
}

func printPath(n *node) {
	curr := n
	path := []int{}
	for curr != nil {
		path = append(path, curr.value)
		curr = curr.parent
	}
	for i := len(path)-1; i >= 0; i-- {
		fmt.Printf("%v ", path[i])
	}
	fmt.Println()
}

func calcTime(nodesWithApple []*node) int {
	edges := map[string]struct{}{}
	edgeString := ""
	for _, n := range nodesWithApple {
		
		for n != nil {
			if (n.parent != nil) {
				a := math.Min(float64(n.value), float64(n.parent.value))
				b := math.Max(float64(n.value), float64(n.parent.value))
				edgeString = strconv.Itoa(int(a)) + "-" + strconv.Itoa(int(b))
				edges[edgeString] = struct{}{}
			}
			n = n.parent
			
		}
	}
	return 2 * len(edges)
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	if (n == 0) {
		return 0
	}

	edgeMatrix := makeEdgeMatrix(n, edges)

	stack := []*node{}

	root := new(node)
	root = new(node)
	root.parent = nil
	root.children = make([]node, 0)
	root.value =  0
	stack = append(stack, root)

	var neighbours *list.List

	nodesWithApple := []*node{}

	for len(stack) != 0{

		curr := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if hasApple[curr.value] {
			nodesWithApple = append(nodesWithApple, curr)
			hasApple[curr.value] = false
		}
		
		neighbours = edgeMatrix[curr.value]
		for n := neighbours.Front(); n != nil; n = n.Next() {
			if curr.parent == nil ||
			   curr.parent.value != n.Value.(int) { 
				
				next := new(node)
				next.parent = curr
				next.value = n.Value.(int)
				next.children = make([]node, 0)

				stack = append(stack, next)
				
			}
		}
		
	}

	
    return calcTime(nodesWithApple)
}

func main() {
	n := 7
	edges := [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}
	hasApple := []bool{false,false,true,false,true,true,false}
	fmt.Println(minTime(n, edges, hasApple))
}