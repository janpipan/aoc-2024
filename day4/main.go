package main

import (
	"bufio"
	"os"
	"strings"
  "fmt"
)



func getMap(fileName string) [][]string {
  file, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  

  scanner := bufio.NewScanner(file)
  xmasMap := [][]string{}
  for scanner.Scan() {
    line := scanner.Text()
    xmasMap = append(xmasMap, strings.Split(line, ""))
  }

  return xmasMap 
}


func partOne(xmasMap [][]string) int {
  directions := [][]int {
    // {x, y}
    {1,0},
    {1,1},
    {0,1},
    {-1,1},
    {-1,0},
    {-1,-1},
    {0,-1},
    {1,-1},
  } 
  curr := ""
  res := 0
  for y, v := range xmasMap {
    for x, c := range v {
      if c == "X" {
        curr = "X"
        for _, d := range directions {
          dy, dx := y+d[1], x+d[0]
          for dy > -1 && dy < len(xmasMap) && dx > -1 && dx < len(v) { 
            if curr == "XMA" && xmasMap[dy][dx] == "S" {
              res++
              break
            } else if curr == "X" && xmasMap[dy][dx] == "M" {
              curr = "XM"
            } else if curr == "XM" && xmasMap[dy][dx] == "A" {
              curr = "XMA"
            } else {
              break
            } 
            dy += d[1]
            dx += d[0]
          }
          curr = "X"
        }
      }
    }
    curr = ""
  }
  return res 
}


func partTwo(xmasMap [][]string) int{
  res := 0
  for y := 1; y < len(xmasMap) - 1; y++ {
    for x := 1; x < len(xmasMap[y]) - 1; x++ {
      if xmasMap[y][x] == "A" {
        if (xmasMap[y-1][x-1] == "M" &&  xmasMap[y-1][x+1] == "M" &&  xmasMap[y+1][x-1] == "S" && xmasMap[y+1][x+1] == "S") || 
        (xmasMap[y-1][x-1] == "M" &&  xmasMap[y-1][x+1] == "S" &&  xmasMap[y+1][x-1] == "M" && xmasMap[y+1][x+1] == "S") || 
        (xmasMap[y-1][x-1] == "S" &&  xmasMap[y-1][x+1] == "M" &&  xmasMap[y+1][x-1] == "S" && xmasMap[y+1][x+1] == "M") || 
        (xmasMap[y-1][x-1] == "S" &&  xmasMap[y-1][x+1] == "S" &&  xmasMap[y+1][x-1] == "M" && xmasMap[y+1][x+1] == "M") {
          res++
        }
      }
    }
  }
  return res
}

func main() {
  xmasMapTest := getMap("./day4/test.txt")
  fmt.Println(partOne(xmasMapTest))
  fmt.Println(partTwo(xmasMapTest))
  xmasMap := getMap("./day4/input.txt")
  fmt.Println(partOne(xmasMap))
  fmt.Println(partTwo(xmasMap))
  return 
}
