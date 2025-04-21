package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pos struct {
  y int
  x int
}

func getPositions(fileName string) (map[string][]Pos, int, int) {
  file, err := os.Open(fileName)
  if err != nil {
    return nil, 0, 0
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  y := -1 
  var xMax int
  positions := map[string][]Pos{}
  for scanner.Scan() {
    y++
    line := scanner.Text()
    l := strings.Split(line, "")
    xMax = len(l)
    for x, v := range l {
      if v != "." {
        if _, ok := positions[v]; !ok {
          positions[v] = []Pos{{y,x}}
        } else {
          positions[v] = append(positions[v], Pos{y,x})
        }
      }
    }
  }
  
  return positions, y, xMax  
}

func getUnique(positions map[string][]Pos, yMax int, xMax int) int {
  res := map[Pos]struct{}{} 
  for _, v := range positions {
    for i := range v {
      for j := i + 1; j < len(v); j++ {
        yDist, xDist := int(math.Abs(float64(v[i].y - v[j].y))), int(math.Abs(float64(v[i].x - v[j].x)))
        yAnti1, yAnti2 := v[i].y - yDist, v[j].y + yDist
        var xAnti1, xAnti2 int
        if v[i].x > v[j].x {
          xAnti1, xAnti2 = v[i].x + xDist, v[j].x - xDist 
        } else {
          xAnti1, xAnti2 = v[i].x - xDist, v[j].x + xDist 
        }
        if yAnti1 >= 0 && yAnti1 <= yMax && xAnti1 >= 0 && xAnti1 < xMax {
          res[Pos{yAnti1, xAnti1}]=struct{}{}
        }
        if yAnti2 >= 0 && yAnti2 <= yMax && xAnti2 >= 0 && xAnti2 < xMax {
          res[Pos{yAnti2, xAnti2}]=struct{}{}
        }
      }
    }
  }
  return len(res) 
}


func partTwo(positions map[string][]Pos, yMax int, xMax int) int {
  res := map[Pos]struct{}{}
  for _, v := range positions {
    for _, pos := range v { 
      res[pos] = struct{}{}
    }
    for i := range v {
      for j := i + 1; j < len(v); j++ {
        yDist, xDist := int(math.Abs(float64(v[i].y - v[j].y))), int(math.Abs(float64(v[i].x - v[j].x)))
        yAnti1, yAnti2 := v[i].y - yDist, v[j].y + yDist
        var xAnti1, xAnti2 int
        if v[i].x > v[j].x {
          xAnti1, xAnti2 = v[i].x + xDist, v[j].x - xDist 
        } else {
          xAnti1, xAnti2 = v[i].x - xDist, v[j].x + xDist 
        }
        for yAnti1 >= 0 && yAnti1 <= yMax && xAnti1 >= 0 && xAnti1 < xMax {
          res[Pos{yAnti1, xAnti1}]=struct{}{}
          yAnti1 -= yDist
          if v[i].x > v[j].x {
            xAnti1 += xDist
          } else {
            xAnti1 -= xDist
          }
        }
        for yAnti2 >= 0 && yAnti2 <= yMax && xAnti2 >= 0 && xAnti2 < xMax {
          res[Pos{yAnti2, xAnti2}]=struct{}{}
          yAnti2 += yDist
          if v[i].x > v[j].x {
            xAnti2 -= xDist
          } else {
            xAnti2 += xDist
          }
        }
      }
    }
  }
  //printSol(res, positions, yMax, xMax)
  return len(res) 
}

func printSol(solutions map[Pos]struct{}, positions map[string][]Pos, yMax int, xMax int) {
  resMap := [][]string{}
  for i := range yMax+1 { 
    resMap = append(resMap, []string{}) 
    for j := range xMax+1 { 
      fmt.Print(j)
      resMap[i] = append(resMap[i], ".") 
    }
  }
  for s := range solutions {
    resMap[s.y][s.x] = "#"
  }
  for c, pos := range positions {
    for _, p := range pos {
      resMap[p.y][p.x] = c
    }
  }
  fmt.Println()
  for i := range yMax {
    fmt.Println(resMap[i])
  }
}

func main() {
  fmt.Println("First part test: ", getUnique(getPositions("./day8/test.txt")))
  fmt.Println("First part input: ", getUnique(getPositions("./day8/input.txt")))
  fmt.Println("Second part test: ", partTwo(getPositions("./day8/test.txt")))
  fmt.Println("Second part input: ", partTwo(getPositions("./day8/input.txt")))
  return
}
