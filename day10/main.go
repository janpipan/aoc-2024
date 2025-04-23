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

func getTrailheadsScoreAndRating(pos Position, trailMap [][]int) (int, int) {
  directions := []Position{
    {0,1},
    {-1,0},
    {1,0},
    {0,-1},
  }
  stack := []Position{pos}
  score := map[Position]struct{}{} 
  rating := 0
  yMax, xMax := len(trailMap), len(trailMap[0])
  for len(stack) > 0 {
    curr := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    if !isValid(curr, yMax, xMax) {
      continue
    }
    if trailMap[curr.y][curr.x] == 9 {
      rating += 1
      score[Position{curr.y,curr.x}] = struct{}{}
    }
    for _, direction := range directions {
      neigh := Position{curr.y+direction.y, curr.x+direction.x}
      if isValid(neigh, yMax, xMax) && trailMap[neigh.y][neigh.x] == trailMap[curr.y][curr.x]+1 {
        stack = append(stack, Position{neigh.y, neigh.x})
      }
    }
  }
  return len(score), rating 
}

func partOneAndTwo(filename string) (int, int) {
  trailMap, trailheads := getTrailMap(filename)
  resPartOne, resPartTwo := 0, 0

  for k := range trailheads {
    score, rating := getTrailheadsScoreAndRating(k, trailMap)
    resPartOne += score 
    resPartTwo += rating
  }
  return resPartOne, resPartTwo
}

func main() {
  fmt.Println(partOneAndTwo("./day10/test.txt"))
  fmt.Println(partOneAndTwo("./day10/input.txt"))
  return
}
