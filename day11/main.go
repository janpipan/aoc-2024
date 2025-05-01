package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) []int {
  stones := []int{}
  
  file, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    input := strings.Split(line, " ")
    for _, num := range input {
      n, err := strconv.Atoi(num)
      if err != nil {
        return nil 
      }
      stones = append(stones, n)
    }
  }

  return stones
}

func partOne(stones []int, times int) int {
  for i := 0; i < times; i++ {
    temp := []int{}
    for _, stone := range stones {
      s := strconv.Itoa(stone)
      digits := strings.Split(s, "")
      if stone == 0 {
        temp = append(temp, 1)
      } else if len(digits) % 2 == 1 {
        dig := strings.Join(digits, "")
        d, err := strconv.Atoi(dig) 
        if err != nil {
          return -1 
        }
        temp = append(temp, d * 2024) 
      } else if len(digits) % 2 == 0 {
        first := strings.Join(digits[:int(len(digits)/2)], "")
        second := strings.Join(digits[int(len(digits)/2):], "")
        f, err := strconv.Atoi(first)
        if err != nil {
          return -1 
        }
        s, err := strconv.Atoi(second)
        if err != nil {
          return -1 
        }
        temp = append(temp, f)
        temp = append(temp, s)
      }
    }
    stones = temp
  }
  return len(stones)
}

func main() {
  fmt.Println(partOne(getInput("./day11/test.txt"), 25))
  fmt.Println(partOne(getInput("./day11/input.txt"), 25))
  return 
}
