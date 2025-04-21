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

func main() {
  fmt.Println("First part test: ", getUnique(getPositions("./day8/test.txt")))
  fmt.Println("First part input: ", getUnique(getPositions("./day8/input.txt")))
  return
}
