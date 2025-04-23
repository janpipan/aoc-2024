package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
  y int
  x int
}

func isValid(pos Position, yMax int, xMax int) bool {
  if pos.y >= yMax || pos.x >= xMax || pos.y < 0 || pos.x < 0 {
    return false
  }
  return true    
}

func getTrailMap(filename string) ([][]int, map[Position]struct{}) {

  trailMap := [][]int{}
  
  file, err := os.Open(filename)
  if err != nil {
    return nil, nil
  }

  scanner := bufio.NewScanner(file)

  trailheads := map[Position]struct{}{}
  for row:= 0; scanner.Scan(); row++ {
    line := scanner.Text()
    nums := strings.Split(line, "")
    trailMap = append(trailMap, []int{})
    for col, num := range nums {
      n, err := strconv.Atoi(num)
      if err != nil {
        return nil, nil
      }
      trailMap[row] = append(trailMap[row], n)
      if n == 0 {
        trailheads[Position{row,col}] = struct{}{}
      }
    }
  }
  return trailMap, trailheads
}

func getTrailheadsScore(pos Position, trailMap [][]int) int {
  directions := []Position{
    {0,1},
    {-1,0},
    {1,0},
    {0,-1},
  }
  stack := []Position{pos}
  heads := map[Position]struct{}{} 
  yMax, xMax := len(trailMap), len(trailMap[0])
  for len(stack) > 0 {
    curr := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    if !isValid(curr, yMax, xMax) {
      continue
    }
    if trailMap[curr.y][curr.x] == 9 {
      heads[Position{curr.y,curr.x}] = struct{}{}
    }
    for _, direction := range directions {
      neigh := Position{curr.y+direction.y, curr.x+direction.x}
      if isValid(neigh, yMax, xMax) && trailMap[neigh.y][neigh.x] == trailMap[curr.y][curr.x]+1 {
        stack = append(stack, Position{neigh.y, neigh.x})
      }
    }
  }
  return len(heads) 
}

func partOne(filename string) int {
  trailMap, trailheads := getTrailMap(filename)
  res := 0

  for k := range trailheads {
    res += getTrailheadsScore(k, trailMap)
  }
  return res
}

func main() {
  fmt.Println(getTrailMap("./day10/test.txt"))
  fmt.Println(partOne("./day10/test.txt"))
  fmt.Println(partOne("./day10/input.txt"))
  return
}
