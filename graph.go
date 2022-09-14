package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Graph struct {
	isDirected bool
	Nodes      []string
	Edges      []Edge
}

// type Node string

type Edge struct {
	A string
	B string
}

func NewGraph(nodes []string, edges [][]string) *Graph {
	var Nodes []string
	for _, v := range nodes {
		node := v
		Nodes = append(Nodes, node)
	}
	sort.Strings(Nodes)
	var Edges []Edge
	for _, v := range edges {
		edge := new(Edge)
		edge.A = v[0]
		edge.B = v[1]
		Edges = append(Edges, *edge)
	}
	g := new(Graph)
	g.Nodes = Nodes
	g.Edges = Edges
	return g
}

func (g *Graph) GetAdjacencyMatrix() map[string]map[string]int {
	matrix := map[string]map[string]int{}
	// init the matrix with 0s
	for _, v1 := range g.Nodes {
		matrix[v1] = map[string]int{}
		for _, v2 := range g.Nodes {
			matrix[v1][v2] = 0
		}
	}
	// fill the connection indices
	for _, v := range g.Edges {
		matrix[v.A][v.B] = 1
		if !g.isDirected {
			matrix[v.B][v.A] = 1
		}
	}
	return matrix
}

func (g *Graph) PrintMatrix(matrix map[string]map[string]int) {
	fmt.Println("")
	fmt.Println("Adjacency Matrix")
	var header string = "    "
	for _, v := range g.Nodes {
		header += v + "  "
	}
	fmt.Println(header)
	lines := map[string]string{}
	for _, v := range g.Nodes {
		subM := matrix[v]
		for _, v2 := range g.Nodes {
			lines[v2] += strconv.Itoa(subM[v2]) + "  "
		}
	}
	for _, v := range g.Nodes {
		fmt.Println(v, ":", lines[v])
	}
}

func (g *Graph) GetAdjacencyList() map[string][]string {
	list := make(map[string][]string, len(g.Nodes))
	for _, v := range g.Nodes {
		list[v] = make([]string, 0)
	}
	for _, v := range g.Edges {
		list[v.A] = append(list[v.A], v.B)
	}
	return list
}

func (g *Graph) PrintList(list map[string][]string) {
	fmt.Println("")
	fmt.Println("Adjacency List")
	for _, node := range g.Nodes {
		line := node + ":"
		if len(list[node]) > 0 {
			for _, element := range list[node] {
				line += " -> " + element
			}
		} else {
			line += " [] "
		}

		fmt.Println(line)
	}
}
