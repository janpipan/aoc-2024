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
      if warehouse[i][j] == "O" || warehouse[i][j] == "[" {
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

func getBiggerWarehouse(warehouse [][]string) ([][]string, Pos) {
  newWarehouse := [][]string{}
  var robot Pos
  for i := range warehouse {
    newWarehouse = append(newWarehouse, []string{})
    line := ""
    for j := range warehouse[i] {
      if warehouse[i][j] == "#"{ 
        line += "##"
      } else if warehouse[i][j] == "O" {
        line += "[]"
      } else if warehouse[i][j] == "." {
        line += ".."
      } else if warehouse[i][j] == "@" {
        line += "@."
      }
    }
    for x, v := range strings.Split(line, "") {
      if v == "@" {
        robot = Pos{len(newWarehouse)-1, x} 
      }
      newWarehouse[len(newWarehouse)-1] = append(newWarehouse[len(newWarehouse)-1], v)
    }
  }
  return newWarehouse, robot
}


func partTwo(filename string) int {
  warehouse, movements, _ := getInput(filename)
  bigWarehouse, robot := getBiggerWarehouse(warehouse)

  for _, move := range strings.Split(movements, "") {
    dy, dx := robot.y + moves[move].y, robot.x + moves[move].x
    if bigWarehouse[dy][dx] == "." {
      bigWarehouse[dy][dx] = "@"
      bigWarehouse[robot.y][robot.x] = "."
      robot.y, robot.x = dy, dx
    } else if bigWarehouse[dy][dx] == "#" {
      continue
    } else if bigWarehouse[dy][dx] == "[" || bigWarehouse[dy][dx] == "]" {
      if move == "^" || move == "v" {
        stack := []Pos{{dy,dx}} 
        moveStack := []Pos{}
        seen := map[Pos]struct{}{}
        if bigWarehouse[dy][dx] == "[" {
          stack = append(stack, Pos{dy, dx+1})
        } else {
          stack = append(stack, Pos{dy, dx-1})
        }
        for len(stack) > 0 {
          curr := stack[0]
          stack = stack[1:]
          my, mx := curr.y+moves[move].y, curr.x+moves[move].x
          if bigWarehouse[my][mx] != "#" {
            if bigWarehouse[my][mx] == "[" || bigWarehouse[my][mx] == "]" {
              stack = append(stack, Pos{my,mx})
              if bigWarehouse[my][mx] == "[" {
                stack = append(stack, Pos{my, mx+1})
              } else {
                stack = append(stack, Pos{my, mx-1})
              }
            }
            if _, ok := seen[curr]; !ok {
              moveStack = append(moveStack, curr) 
              seen[curr] = struct{}{}
            }
          } else {
            moveStack = []Pos{}
            break
          }
        }
        if len(moveStack) == 0 {
          continue
        }
        for len(moveStack) > 0 {
          curr := moveStack[len(moveStack)-1]
          moveStack = moveStack[:len(moveStack)-1]
          bigWarehouse[curr.y+moves[move].y][curr.x+moves[move].x], bigWarehouse[curr.y][curr.x] = bigWarehouse[curr.y][curr.x], bigWarehouse[curr.y+moves[move].y][curr.x+moves[move].x] 
        }
        bigWarehouse[robot.y][robot.x], bigWarehouse[dy][dx] = bigWarehouse[dy][dx], bigWarehouse[robot.y][robot.x]
        robot.y, robot.x = dy, dx
      } else {
        stack := []string{}
        for bigWarehouse[dy][dx] == "[" || bigWarehouse[dy][dx] == "]" {
          stack = append(stack, bigWarehouse[dy][dx])
          dy, dx = dy + moves[move].y, dx + moves[move].x
        }
        if bigWarehouse[dy][dx] == "#" {
          continue
        } else {
          for len(stack) > 0 {
            bigWarehouse[dy][dx] = stack[len(stack)-1]
            dy, dx = dy - moves[move].y, dx - moves[move].x
            stack = stack[:len(stack)-1]
          }
          bigWarehouse[dy][dx] = "@"
          bigWarehouse[robot.y][robot.x] = "."
          robot.y, robot.x =  dy, dx
        }
      }
    }
  }

  return gpsCoord(bigWarehouse) 
}

func main() {
  fmt.Println(partOne("./day15/test.txt"))
  fmt.Println(partOne("./day15/input.txt"))
  fmt.Println(partTwo("./day15/test.txt"))
  fmt.Println(partTwo("./day15/input.txt"))
  return
}
