package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
  y int
  x int
}

var moves = map[string]Pos {
  ">": {0,1},
  "<": {0,-1},
  "v": {1,0},
  "^": {-1,0},
}

func getInput(filename string) ([][]string, string, Pos) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, "", Pos{} 
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  warehouse := [][]string{}
  movements := ""
  var robot Pos 
  
  for scanner.Scan() {
    line := scanner.Text()
    l := strings.Split(line, "")
    if len(l) == 0 {
      continue
    }
    if l[0] == "#" {
      warehouse = append(warehouse, []string{})
      for i, v := range l {
        if v == "@" {
          robot = Pos{len(warehouse)-1, i} 
        }
        warehouse[len(warehouse)-1] = append(warehouse[len(warehouse)-1], v)
      }
    } else {
      movements += line
    }
  }
  return warehouse, movements, robot 
}

func printWarehouse(warehouse [][]string) {
  for _, line := range warehouse {
    fmt.Println(line)
  }
  return
}

func gpsCoord(warehouse [][]string) int{
  res := 0

  for i := range warehouse {
    for j := range warehouse[i] {
      if warehouse[i][j] == "O"{
        res += 100 * i + j
      }
    }
  }

  return res 
}

func partOne(filename string) int{
  warehouse, movements, robot := getInput(filename)
  
  for _, move := range strings.Split(movements, "") {
    dy, dx := robot.y + moves[move].y, robot.x + moves[move].x
    if warehouse[dy][dx] == "." {
      warehouse[dy][dx] = "@"
      warehouse[robot.y][robot.x] = "."
      robot.y, robot.x = dy, dx

    } else if warehouse[dy][dx] == "#" {
      continue
    } else if warehouse[dy][dx] == "O" {
      for warehouse[dy][dx] == "O" {
        dy, dx = dy + moves[move].y, dx + moves[move].x
      }
      if warehouse[dy][dx] == "#" {
        continue
      } else {
        warehouse[dy][dx] = "O"
        dy, dx := robot.y + moves[move].y, robot.x + moves[move].x
        warehouse[dy][dx] = "@"
        warehouse[robot.y][robot.x] = "."
        robot.y, robot.x =  dy, dx
      }
    }
  }

  return gpsCoord(warehouse) 
}

func main() {
  fmt.Println(partOne("./day15/test.txt"))
  fmt.Println(partOne("./day15/input.txt"))
  return
}
