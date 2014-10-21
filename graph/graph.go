package graph

import (
  "../queue"
)

type Node struct {
  adjacents []int
  prec int
  distance int
  visited bool
}

type Graph struct {
  nodes map[int]*Node
}

func NewGraph() *Graph {
  return &Graph{
    nodes: make(map[int]*Node),
  }
}

func (graph *Graph) AddNode(key int, adjacents []int) {
  graph.nodes[key] = &Node{adjacents: adjacents, prec: -1, distance: -1, visited: false}
}

func (graph *Graph) BFS(start int) []int {
  q := queue.NewQueue(1)
  e := make([]int, 0)

  graph.nodes[start].visited = true
  graph.nodes[start].distance = 0
  graph.nodes[start].prec = 1
  q.Push(start)

  for q.Count() > 0 {
    key, _ := q.Pop()

    for _, adj := range graph.nodes[key].adjacents {
      if ! graph.nodes[adj].visited {
        graph.nodes[adj].visited = true
        graph.nodes[adj].prec = key
        graph.nodes[adj].distance += 1

        q.Push(adj)
      }
    }

    e = append(e, key)
  }

  return e
}
