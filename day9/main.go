package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(fileName string) string {
  file, err := os.Open(fileName)
  if err != nil {
    return "" 
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var res string
  for scanner.Scan() {
    line := scanner.Text() 
    return line 
  }
  return res
}

func getCheckSum(input string) int {
  sum := 0
  i := strings.Split(input, "")
  id := 0
  b := []int{}
  for i, c := range i {
    t, err := strconv.Atoi(c)
    if err != nil {
      return 0
    }
    if i % 2 == 0 {
      for j := 0; j < t; j++ {
        b = append(b, id)
      }
      id++
    } else {
      for j := 0; j < t; j++ {
        b = append(b, -1)
      }
    }
  }
  l, r := 0, len(b)-1
  for l < r {
    if b[l] == -1 {
      if b[r] == -1 {
        for b[r] == -1 {
          r--
        }
      }
      b[l], b[r] = b[r], b[l]
      l++
      r--
    } else {
      l++
    } 
  }
  for i := 0; b[i] != -1; i++ {
    sum += i * b[i]
  }
  return sum
}


func main () {
  fmt.Println("First part test-1: ", getCheckSum(getInput("./day9/test-1.txt")))
  fmt.Println("First part test-2: ", getCheckSum(getInput("./day9/test-2.txt")))
  fmt.Println("First part input:", getCheckSum(getInput("./day9/input.txt")))
  return 
}
