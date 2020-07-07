package main

import "fmt"

func main() {
  graph := *Graph{numberOfNodes: 2, adjecentList: map[int][]int{0: [1, 2], 1: [0, 1],}}
}

type Graph struct {
	numberOfNodes int
	adjecentList  map[interface{}][]interface{}
}

func New() *Graph {
  return &Graph{
    numberOfNodes: 0,
    adjecentList: make(map[interface{}][]interface{}),
  }
}

func (g *Graph) AddVertex(node int) {
	if _, ok := g.adjecentList[node]; !ok {
		g.adjecentList[node] = []int{}
	}
	g.numberOfNodes++
}

func (g *Graph) AddConnection(node1, node2 int) {
	if _, ok1 := g.adjecentList[node1]; ok1 {
		if _, ok2 := g.adjecentList[node2]; ok2 {
			g.adjecentList[node1] = append(g.adjecentList[node1], node2)
			g.adjecentList[node2] = append(g.adjecentList[node2], node1)
		}
	}
}

func (g *Graph) ShowConnections() {
	for node, connections := range g.adjecentList {
		fmt.Println(node, connections)
	}
}
