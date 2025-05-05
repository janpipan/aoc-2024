package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
  y int
  x int
  dy int
  dx int
}

type Node struct {
  position Pos
  cost int 
  index int
}

type Heap []*Node

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h Heap) Swap(i, j int) { 
  h[i], h[j] = h[j], h[i] 
  h[i].index = i
  h[j].index = j
}

func (h *Heap) Push(x any) {
  n := len(*h)
  item := x.(*Node)
  item.index = n
  *h = append(*h, item) 
}

func (h *Heap) Pop() any {
  old := *h
  n := len(old)
  x := old[n-1]
  old[n-1] = nil
  x.index = -1
  *h = old[:n-1]
  return x
}

func getInput(filename string) ([][]string, Pos) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, Pos{}
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  puzzleMap := [][]string{}
  start := Pos{} 

  for scanner.Scan() {
    line := scanner.Text()
    puzzleMap = append(puzzleMap, []string{})
    for x, l := range strings.Split(line, "") {
      if l == "S" {
        start = Pos{len(puzzleMap)-1, x, 0, 1}
      }
      puzzleMap[len(puzzleMap)-1] = append(puzzleMap[len(puzzleMap)-1], l)
    }
  }
  return puzzleMap, start
}

func partOne(filename string) int {
  maze, start := getInput(filename)
  return getSolution(maze, start) 
}


func getSolution(maze [][]string, start Pos) int {
  pq := Heap{}
  seen := map[Pos]struct{}{start:{}}
  heap.Init(&pq)
  heap.Push(&pq, &Node{start, 0, 0})
  for pq.Len() > 0 {
    n := *heap.Pop(&pq).(*Node)
    cost, y, x, dy, dx := n.cost, n.position.y, n.position.x, n.position.dy, n.position.dx
    seen[n.position] = struct{}{}
    if maze[n.position.y][n.position.x] == "E" {
      return n.cost
    }
    for _, node := range []Node{{position: Pos{y+dy, x+dx, dy,dx},cost: cost + 1},{position: Pos{y, x, dx,-dy}, cost: cost + 1000},{position: Pos{y, x, -dx,dy},cost: cost + 1000}} {
      if maze[node.position.y][node.position.x] == "#" { continue }
      if _, ok := seen[node.position]; ok { continue }
      heap.Push(&pq, &node)
    }
  }
  return 0 
}


func main() {
  fmt.Println(partOne("./day16/test.txt"))
  fmt.Println(partOne("./day16/input.txt"))
  return 
}
