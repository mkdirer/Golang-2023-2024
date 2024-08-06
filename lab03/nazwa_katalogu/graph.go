package main

import (
	"fmt"
)

type Node struct {
	ID  int
	In  []int
	Out []int
}

type Graph struct {
	Nodes []Node
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddNode(id int) {
	g.Nodes = append(g.Nodes, Node{ID: id})
}

func (g *Graph) AddEdge(src, dest int) {
	g.Nodes[src].Out = append(g.Nodes[src].Out, dest)
	g.Nodes[dest].In = append(g.Nodes[dest].In, src)
}

func (g *Graph) Print() {
	for _, node := range g.Nodes {
		fmt.Printf("Node %d:\n", node.ID)
		fmt.Printf("  In:  %v\n", node.In)
		fmt.Printf("  Out: %v\n", node.Out)
	}
}

func (g *Graph) DegreeDistribution() (inDegrees, outDegrees []int) {
	numNodes := len(g.Nodes)
	inDegrees = make([]int, numNodes)
	outDegrees = make([]int, numNodes)

	for _, node := range g.Nodes {
		inDegrees[len(node.In)]++
		outDegrees[len(node.Out)]++
	}

	return inDegrees, outDegrees
}

func (g *Graph) ShortestPaths() [][]int {
	numNodes := len(g.Nodes)
	dist := make([][]int, numNodes)

	for i := range dist {
		dist[i] = make([]int, numNodes)
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
		dist[i][i] = 0
	}

	for _, node := range g.Nodes {
		src := node.ID
		for _, dest := range node.Out {
			dist[src][dest] = 1
		}
	}

	for k := 0; k < numNodes; k++ {
		for i := 0; i < numNodes; i++ {
			for j := 0; j < numNodes; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}
