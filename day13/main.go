package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)


func getInput(filename string) [][]int{
  file, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file) 

  machines := [][]int{{}}

  for scanner.Scan() {
    line := scanner.Text()
    r, _ := regexp.Compile("\\d+")
    matches := r.FindAllString(line, -1)
    if len(matches) == 0 {
      machines = append(machines, []int{})
    }
    for _, m := range matches {
      num, _ := strconv.Atoi(m)
      machines[len(machines)-1] = append(machines[len(machines)-1], num)
    }

  }
  return machines 
}

func partOne(machines [][]int) int {
  res := 0
  for _, machine := range machines {
    a1, a2, b1, b2, s1, s2 := machine[0], machine[1], machine[2], machine[3], machine[4], machine[5] 
    up := s1 * b2 - s2 * b1
    bot := b2 * a1 - b1 * a2
    if bot != 0 {
      x := float64(up) / float64(bot)
      y := (float64(s1) - float64(a1) * x) / float64(b1) 
      if x == math.Trunc(x) && y == math.Trunc(y) {
        res += 3 * int(x) + int(y)
      }
    }
  }
  return res
}


func partTwo(machines [][]int) int {
  res := 0
  for _, machine := range machines {
    a1, a2, b1, b2, s1, s2 := machine[0], machine[1], machine[2], machine[3], machine[4] + 10000000000000, machine[5] + 10000000000000 
    up := s1 * b2 - s2 * b1
    bot := b2 * a1 - b1 * a2
    if bot != 0 {
      x := float64(up) / float64(bot)
      y := (float64(s1) - float64(a1) * x) / float64(b1) 
      if x == math.Trunc(x) && y == math.Trunc(y) {
        res += 3 * int(x) + int(y)
      }
    }
  }
  return res
}

func main() {
  fmt.Println(partOne(getInput("./day13/test.txt")))
  fmt.Println(partOne(getInput("./day13/input.txt")))
  fmt.Println(partTwo(getInput("./day13/test.txt")))
  fmt.Println(partTwo(getInput("./day13/input.txt")))
  return
}
