package main

import (
  "fmt"
  "./queue"
)

type Node struct {
  adjacents []int
  prec int
  distance int
  visited bool
}

func main() {
  graph := make(map[int]*Node)
  graph[1] = &Node{adjacents: []int{2,8}, prec: -1, distance: -1, visited: false}
  graph[2] = &Node{adjacents: []int{1,3}, prec: -1, distance: -1, visited: false}
  graph[3] = &Node{adjacents: []int{2,4,5,7,8}, prec: -1, distance: -1, visited: false}
  graph[4] = &Node{adjacents: []int{3,5}, prec: -1, distance: -1, visited: false}
  graph[5] = &Node{adjacents: []int{3,4,6}, prec: -1, distance: -1, visited: false}
  graph[6] = &Node{adjacents: []int{5}, prec: -1, distance: -1, visited: false}
  graph[7] = &Node{adjacents: []int{3,8}, prec: -1, distance: -1, visited: false}
  graph[8] = &Node{adjacents: []int{1,3,7,9}, prec: -1, distance: -1, visited: false}
  graph[9] = &Node{adjacents: []int{8}, prec: -1, distance: -1, visited: false}

  distance := 0
  fmt.Println(BFS(graph, 1, 1, distance))
}

func BFS(graph map[int]*Node, current int, prec int, distance int) []int {
  q = queue.NewQueue(0)
  e = make([]int, 0)

  graph[current].visited = true
  graph[current].distance = distance
  graph[current].prec = 1
  q.Push(current)

  for q.Count() > 0 {
    key, _ := q.Pop()

    for _, adj := range graph[key].adjacents {
      if ! graph[adj].visited {
        graph[adj].visited = true
        graph[adj].prec = key
        graph[adj].distance += 1

        q.Push(adj)
      }
    }

    e = append(e, key)
  }

  return e
}
