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

func partTwo(regions [][]string) int {
  res := 0
  directions := [][]int{
    {0,1},
    {1,0},
    {0,-1},
    {-1,0},
  }
  corners := [][]int{
    {0,1,1,0,1,1},
    {1,0,0,-1,1,-1},
    {0,-1,-1,0,-1,-1},
    {-1,0,0,1,-1,1},
  }
  seen := map[Position]struct{}{}
  yMax, xMax := len(regions), len(regions[0])
  for i := range regions {
    for j := range regions[i] {
      pos := Position{i, j}
      if _, ok := seen[pos]; ok {
        continue
      }
      area, corner := 0, 0 
      q := []Position{{i,j}}
      region := map[Position]struct{}{}
      for len(q) > 0 {
        a, c := 1, 0 
        curr := q[0]
        seen[curr] = struct{}{}
        region[Position{curr.y,curr.x}] = struct{}{}
        // check corners  
        for _, cor := range corners {
          fy, fx, sy, sx, ty, tx := curr.y + cor[0], curr.x + cor[1], curr.y + cor[2], curr.x + cor[3], curr.y + cor[4], curr.x +cor[5]
          // check if valid
          firstDiff, secondDiff := false, false
          if fy < 0 || fx < 0 || fy >= yMax || fx >= xMax {
            firstDiff = true 
          } else {
            if regions[curr.y][curr.x] != regions[fy][fx] {
              firstDiff = true
            }
          }
          if sy < 0 || sx < 0 || sy >= yMax || sx >= xMax {
            secondDiff = true 
          } else {
            if regions[curr.y][curr.x] != regions[sy][sx] {
              secondDiff = true
            }
          }
          if !firstDiff && !secondDiff {
            if (ty < 0 || tx < 0 || ty >= yMax || tx >= xMax) {
              c++
            } else if regions[curr.y][curr.x] != regions[ty][tx] {
              c++
            }
          } else if firstDiff && secondDiff {
            c++
          }
        }
        for _, dir := range directions {
          dy, dx := curr.y + dir[0], curr.x + dir[1]
          if dy < 0 || dx < 0 || dy >= yMax || dx >= xMax {
            continue
          } else if regions[dy][dx] == regions[curr.y][curr.x] {
            if _, ok := region[Position{dy,dx}]; !ok {
              region[Position{dy,dx}] = struct{}{}
              q = append(q, Position{dy,dx})
            }
          } 
        }
        q = q[1:]
        area += a
        corner += c  
      }
      res += area * corner 
    }
  }
  return res 
}

func main() {
  fmt.Println(partOne(getInput("./day12/test.txt")))
  fmt.Println(partOne(getInput("./day12/input.txt")))
  fmt.Println(partTwo(getInput("./day12/test.txt")))
  fmt.Println(partTwo(getInput("./day12/input.txt")))
  return
}
