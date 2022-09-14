package main

import "fmt"

func main() {
	Nodes := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	Edges := [][]string{{"A", "B"}, {"B", "C"}, {"E", "E"}}
	graph := NewGraph(Nodes, Edges)
	x := graph.GetAdjacencyMatrix()
	graph.PrintMatrix(x)
	L := graph.GetAdjacencyList()
	graph.PrintList(L)
	fmt.Println("")
}
