package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func getInput(filename string) [][]string {
  file, err := os.Open(filename) 
  if err != nil {
    return nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  regions := [][]string{}
  idx := 0
  for scanner.Scan() {
    line := scanner.Text()
    regions = append(regions, []string{})
    for _, plant := range strings.Split(line, "") {
      regions[idx] = append(regions[idx], plant) 
     }
    idx++
  }
  return regions 
}

type Position struct {
  y int
  x int
}


func partOne(regions [][]string) int {
  res := 0
  directions := [][]int{
    {0,1},
    {1,0},
    {0,-1},
    {-1,0},
  }
  seen := map[Position]struct{}{}
  yMax, xMax := len(regions), len(regions[0])
  for i := range regions {
    for j := range regions[i] {
      pos := Position{i, j}
      if _, ok := seen[pos]; ok {
        continue
      }
      area, perimiter := 0, 0 
      q := []Position{{i,j}}
      region := map[Position]struct{}{}
      for len(q) > 0 {
        a, p := 1, 4
        curr := q[0]
        seen[curr] = struct{}{}
        region[Position{curr.y,curr.x}] = struct{}{}
        for _, dir := range directions {
          dy, dx := curr.y + dir[0], curr.x + dir[1]
          if dy < 0 || dx < 0 || dy >= yMax || dx >= xMax {
            continue
          } else if regions[dy][dx] == regions[curr.y][curr.x] {
            p--
            if _, ok := region[Position{dy,dx}]; !ok {
              region[Position{dy,dx}] = struct{}{}
              q = append(q, Position{dy,dx})
            }
          } 
        }
        q = q[1:]
        area += a
        perimiter += p
      }
      res += area * perimiter
    }
  }
  return res 
}

func main() {
  fmt.Println(partOne(getInput("./day12/test.txt")))
  fmt.Println(partOne(getInput("./day12/input.txt")))
  return
}
