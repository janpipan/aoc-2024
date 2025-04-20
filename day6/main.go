package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pos struct {
  y int
  x int
  dir int
}

func getMap(fileName string) ([][]string, int, int) {
  guardMap := [][]string{}
  x, y := -1, -1
  file, err := os.Open(fileName)
  if err != nil {
    return nil, x, y 
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  col := -1 
  for scanner.Scan() {
    col++
    line := scanner.Text()
    splitLine := strings.Split(line, "")
    if slices.Contains(splitLine, "^") {
      for i := range (splitLine) {
        if splitLine[i] == "^" {
          x = i
          y = col
        }
      }
    }
    guardMap = append(guardMap, splitLine)
  }
  return guardMap , x, y 
}

func partOne(guardMap [][]string, x int, y int) int {
  res := 1
  d := 0
  direction := [][]int{
    {-1,0},
    {0,1},
    {1,0},
    {0,-1},
  }
  for x < len(guardMap[0])  && x >= 0 && y < len(guardMap) && y >= 0 {
    if guardMap[y][x] == "." {
      res++
    }
    guardMap[y][x] = "X"
    dy, dx := y+direction[d][0], x+direction[d][1]
    if dx < len(guardMap[0])  && dx >= 0 && dy < len(guardMap) && dy >= 0 {
      if guardMap[dy][dx] == "#" {
        if d+1 > 3{
          d = 0
        } else {
          d++
        }
      } else {
        y, x = dy, dx
      }
    } else {
      y, x = dy, dx
    }
  }
  return res
}

func printMap(guardMap [][]string) {
  fmt.Print(" ")
  for i := range guardMap {
    fmt.Print(i, " ")
  }
  fmt.Println()
  for i, r := range guardMap {
    fmt.Println(r, i)
  }
}

func isLoop(guardMap [][]string, d int, start []int) bool {
  direction := [][]int{
    {-1,0},
    {0,1},
    {1,0},
    {0,-1},
  }
  visited := map[Pos]struct{}{}
  y, x := start[0], start[1]
  for x < len(guardMap[0])  && x >= 0 && y < len(guardMap) && y >= 0 {
    if _, ok := visited[Pos{y, x, d}]; ok {
      return true
    }
    visited[Pos{y, x, d}] = struct{}{}
    dy, dx := y+direction[d][0], x+direction[d][1]
    if dx < len(guardMap[0])  && dx >= 0 && dy < len(guardMap) && dy >= 0 {
      if guardMap[dy][dx] == "#" {
        if d+1 > 3{
          d = 0
        } else {
          d++
        }
      } else {
        y, x = dy, dx
      }
    } else {
      y, x = dy, dx
    }
  }
  return false 
}

func partTwo(guardMap [][]string, start []int) int {
  res := map[Pos]struct{}{}
  y, x := start[0], start[1]
  d := 0
  direction := [][]int{
    {-1,0},
    {0,1},
    {1,0},
    {0,-1},
  }
  for x < len(guardMap[0])  && x >= 0 && y < len(guardMap) && y >= 0 {
    if guardMap[y][x] != "^" {
      guardMap[y][x] = "X"
    }
    dy, dx := y+direction[d][0], x+direction[d][1]
    if dx < len(guardMap[0])  && dx >= 0 && dy < len(guardMap) && dy >= 0 {
      if guardMap[dy][dx] == "#" {
        if d+1 > 3{
          d = 0
        } else {
          d++
        }
      } else {
        if guardMap[dy][dx] != "^" {
          tmp := guardMap[dy][dx]
          guardMap[dy][dx] = "#"
          if isLoop(guardMap, 0, start){
            res[Pos{dy,dx,0}] = struct{}{}
          }
          guardMap[dy][dx] = tmp
        }
        y, x = dy, dx
      }
    } else {
      y, x = dy, dx
    }
  }
  return len(res)
}

func main() {
  guardMapTest, xTest, yTest := getMap("./day6/test.txt")
  guardMapTestTwo, _, _ := getMap("./day6/test.txt")
  fmt.Println("Test result:", partOne(guardMapTest, xTest, yTest))
  fmt.Println("Test result two:", partTwo(guardMapTestTwo, []int{yTest, xTest}))
  guardMap, x, y := getMap("./day6/input.txt")
  guardMapTwo, _, _ := getMap("./day6/input.txt")
  fmt.Println("Part one result:", partOne(guardMap, x, y))
  fmt.Println("Part two result:", partTwo(guardMapTwo, []int{y, x}))
  return 
}
