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
	sort.Slice(Nodes, func(i, j int) bool {
		return Nodes[i] < Nodes[j]
	})
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

func (g *Graph) GetAdjencancyMatrix() map[string]map[string]int {
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
	var header string = "    "
	keys := make([]string, len(matrix))
	i := 0
	for k, _ := range matrix {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, v := range keys {
		header += v + "  "
	}
	fmt.Println(header)
	lines := map[string]string{}
	for _, v := range keys {
		subM := matrix[v]
		for _, v2 := range keys {
			lines[v2] += strconv.Itoa(subM[v2]) + "  "
		}
	}
	for k, v := range lines {
		fmt.Println(k, ":", v)
	}
}
