package main

import (
  "fmt"
  "./graph"
)

func main() {
  graph := graph.NewGraph()
  graph.AddNode(1, []int{2,8})
  graph.AddNode(2, []int{1,3})
  graph.AddNode(3, []int{2,4,5,7,8})
  graph.AddNode(4, []int{3,5})
  graph.AddNode(5, []int{3,4,6})
  graph.AddNode(6, []int{5})
  graph.AddNode(7, []int{3,8})
  graph.AddNode(8, []int{1,3,7,9})
  graph.AddNode(9, []int{8})

  fmt.Println(graph.BFS(1))
}
