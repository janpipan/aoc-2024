package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
  y int
  x int
}

type Node struct {
  pos Position
  dist int
}
  

func getInput(filename string) []Position {
  res := []Position{}

  file, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    nums := strings.Split(line, ",")
    x, _ := strconv.Atoi(nums[0])
    y, _ := strconv.Atoi(nums[1])
    res = append(res, Position{y:y,x:x})
  }
  return res 
}

func makeMap(bytePositions []Position, x, y, steps int) [][]int {
  byteMap := [][]int{}
  for i := range y {
    byteMap = append(byteMap, []int{})
    for j := 0; j < x; j++ {
      byteMap[i] = append(byteMap[i], 0)
    }
  }
  for i := range steps {
    byteMap[bytePositions[i].y][bytePositions[i].x] = 1 
  }

  return byteMap
}

func printMap(byteMap [][]int) {
  for i := range len(byteMap){
    fmt.Println(byteMap[i])
  }
  return
}

func partOne(filename string, x, y, steps int) int{
  positions := getInput(filename)
  byteMap := makeMap(positions, x, y, steps)
  q := []Node{{Position{0,0},0}} 
  seen := map[Position]struct{}{}
  m := int(math.MaxInt)

  for len(q) > 0 {
    curr := q[0]
    q = q[1:]
    for _, pos := range []Node{{Position{curr.pos.y, curr.pos.x+1}, curr.dist+1,}, {Position{curr.pos.y+1,curr.pos.x}, curr.dist+1}, {Position{curr.pos.y, curr.pos.x-1}, curr.dist+1}, {Position{curr.pos.y-1,curr.pos.x},curr.dist+1}} {
      if pos.pos.y < 0 || pos.pos.y >= len(byteMap) || pos.pos.x < 0 || pos.pos.x >= len(byteMap[0]) { continue }
      if byteMap[pos.pos.y][pos.pos.x] == 1 { continue }
      if _, ok := seen[pos.pos]; ok { continue }
      if pos.pos.y == len(byteMap) - 1 && pos.pos.x == len(byteMap[0]) - 1 { m = min(m, pos.dist) }
      seen[pos.pos] = struct{}{}
      q = append(q, pos)
    }
  }
  return m 
}

func isConnected(positions []Position, x, y, steps int) bool {
  byteMap := makeMap(positions, x, y, steps)
  q := []Position{{0,0}}
  seen := map[Position]struct{}{}

  for len(q) > 0 {
    curr := q[0]
    q = q[1:]
    for _, pos := range []Position{{curr.y, curr.x+1}, {curr.y+1,curr.x}, {curr.y, curr.x-1}, {curr.y-1,curr.x}} {
      if pos.y < 0 || pos.y >= len(byteMap) || pos.x < 0 || pos.x >= len(byteMap[0]) { continue }
      if byteMap[pos.y][pos.x] == 1 { continue }
      if _, ok := seen[pos]; ok { continue }
      if pos.y == len(byteMap) - 1 && pos.x == len(byteMap[0]) - 1 { return true}
      seen[pos] = struct{}{}
      q = append(q, pos)
    }
  }
  return false 
}

func partTwo(filename string, x, y, start int) Position {
  positions := getInput(filename)
  end := len(positions)
  for start < end {
    mid := int((start + end) / 2)
    r := isConnected(positions, x, y, mid+1)
    if r {
      start = mid + 1 
    } else {
      end = mid 
    }
  }
  return positions[start] 
}

func main() {
  fmt.Println(partOne("./day18/test.txt", 7, 7, 12))
  fmt.Println(partOne("./day18/input.txt", 71, 71, 1024))
  fmt.Println(partTwo("./day18/test.txt", 7, 7, 12))
  fmt.Println(partTwo("./day18/input.txt", 71, 71, 1024))
  return 
}
