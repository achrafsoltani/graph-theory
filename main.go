package main

import "fmt"

func main() {
	fmt.Println("test Graphs")
	Nodes := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	Edges := [][]string{{"A", "B"}, {"B", "C"}}
	graph := NewGraph(Nodes, Edges)
	x := graph.GetAdjencancyMatrix()
	graph.PrintMatrix(x)
	//fmt.Println(graph.Nodes, x)

}
