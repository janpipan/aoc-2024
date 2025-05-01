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

type State struct {
  num int
  times int
}

var cache = make(map[State]int)

func partTwo(stones []int, times int) int {
  res := 0
  for _, stone := range stones {
    res += blink(stone, times)
  }
  return res 
}

func blink(num int, times int) int {
  if times == 0 {
    return 1 
  }
  state := State{num, times}
  if v, ok := cache[state]; ok {
    return v
  }
  if num == 0 {
    res := blink(1, times - 1)
    cache[state] = res
    return res 
  } else {
    s := strconv.Itoa(num)
    digits := strings.Split(s, "")
    if len(digits) % 2 == 1 {
      dig := strings.Join(digits, "")
      d, err := strconv.Atoi(dig) 
      if err != nil {
        return -1 
      }
      res := blink(d * 2024, times - 1)
      cache[state] = res
      return res 
    } else {
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
      res := blink(f, times - 1) + blink(s, times - 1) 
      cache[state] = res
      return res 
    }
  } 
}

func main() {
  fmt.Println(partOne(getInput("./day11/test.txt"), 25))
  fmt.Println(partOne(getInput("./day11/input.txt"), 25))
  //fmt.Println(partOne(getInput("./day11/input.txt"), 75))
  fmt.Println(partTwo(getInput("./day11/test.txt"),6))
  fmt.Println(partTwo(getInput("./day11/test.txt"),25))
  fmt.Println(partTwo(getInput("./day11/input.txt"),75))
  return 
}
