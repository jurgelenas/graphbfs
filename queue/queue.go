/* 
  Queue implementation taken from https://gist.github.com/moraes/2141121

  I have only simplified it by removing unnecessary complexity of Node struct and
  added Count() method.
*/

package queue

import (
  "errors"
)

type Queue struct {
  nodes []int
  size  int
  head  int
  tail  int
  count int
}

func NewQueue(size int) *Queue {
  return &Queue{
    nodes: make([]int, size),
    size:  size,
  }
}

func (q *Queue) Count() int {
  return q.count
}

func (q *Queue) Push(n int) {
  if q.head == q.tail && q.count > 0 {
    nodes := make([]int, len(q.nodes)+q.size)
    copy(nodes, q.nodes[q.head:])
    copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
    q.head = 0
    q.tail = len(q.nodes)
    q.nodes = nodes
  }

  q.nodes[q.tail] = n
  q.tail = (q.tail + 1) % len(q.nodes)
  q.count++
}
 
func (q *Queue) Pop() (int, error) {
  if q.count == 0 {
    return 0, errors.New("Queue is empty")
  }

  node := q.nodes[q.head]
  q.head = (q.head + 1) % len(q.nodes)
  q.count--

  return node, nil
}
