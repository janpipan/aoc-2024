package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Position struct {
  y int
  x int
}

func getInput(filename string) ([][]int, Position) {
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  racetrack := [][]int{}
  start := Position{}
  

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    racetrack = append(racetrack, []int{})
    l := strings.Split(line, "")
    for i, p := range l {
      var num int
      if p == "#" {
        num = -2
      } else if p == "." { 
        num = -1
      } 
      if p == "S" {
        start.y, start.x = len(racetrack)-1, i 
        num = 0
      } else if p == "E" {
        num = -1
      }
      racetrack[len(racetrack)-1] = append(racetrack[len(racetrack)-1], num)
    }
  }
  return racetrack, start
}

func printRacetrack(racetrack [][]int) {
  for i := range len(racetrack) {
    fmt.Println(racetrack[i])
  }
  return
}

func calculateCost(racetrack *[][]int, start Position) {
  q := []Position{start}
  for len(q) > 0 {
    curr := q[0]
    q = q[1:]
    for _, pos := range []Position{{curr.y+1,curr.x}, {curr.y-1, curr.x}, {curr.y, curr.x+1}, {curr.y, curr.x-1}} {
      if pos.y < 0 || pos.x < 0 || pos.y >= len(*racetrack) || pos.x >= len(*racetrack) { continue }
      if (*racetrack)[pos.y][pos.x] == -2 { continue }
      if (*racetrack)[pos.y][pos.x] != -1 { continue }
      (*racetrack)[pos.y][pos.x] = (*racetrack)[curr.y][curr.x] + 1
      q = append(q, pos)
    }
  }
  return
}

func savings(racetrack *[][]int) int {
  res := 0
  for i := range len(*racetrack) {
    for j := range len((*racetrack)[i]) {
      if (*racetrack)[i][j] == -2 { continue }
      for _, pos := range []Position{{i+2, j}, {i+1,j+1}, {i,j+2}, {i-1,j+1}} {
        if pos.y < 0 || pos.x < 0 || pos.y >= len(*racetrack) || pos.x >= len((*racetrack)[i]) { continue }
        if (*racetrack)[pos.y][pos.x] == -2 { continue }
        if int(math.Abs(float64((*racetrack)[i][j]) - float64((*racetrack)[pos.y][pos.x]))) >= 102 { res++ } 
      }
    }
  }
  return res 
}

func partOne(filename string) int {
  racetrack, start := getInput(filename)
  //printRacetrack(racetrack)
  calculateCost(&racetrack, start)
  //printRacetrack(racetrack)
  return savings(&racetrack) 
}

func main() {
  fmt.Println(partOne("./day20/test.txt"))
  fmt.Println(partOne("./day20/input.txt"))
  return
}
