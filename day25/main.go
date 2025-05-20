package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInput(filename string) ([][][]string, error) {
  input := [][][]string{{}}
  file, err := os.Open(filename)
  if err != nil {
    return nil, fmt.Errorf("Err: %w", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  row := 0

  for scanner.Scan() {
    line := scanner.Text()
    if len(line) == 0 {
      input = append(input, [][]string{})
      row = 0
    } else {
      input[len(input)-1] = append(input[len(input)-1], []string{})
      for _, val := range line {
        input[len(input)-1][row] = append(input[len(input)-1][row], string(val))
      }
      row++
    }
  }
  return input, err
}

func getKeysAndLocks(input [][][]string) ([][]int, [][]int) {
  keys, locks := [][]int{}, [][]int{}

  for _, obj := range input {
    if obj[0][0] == "#" {
      locks = append(locks, []int{})
      for c := range len(obj[0]) {
        counter := 0
        for r := 1; r < len(obj) - 1; r++ {
          if obj[r][c] == "#" { counter ++ }
        }
        locks[len(locks)-1] = append(locks[len(locks)-1], counter)
      }
    } else {
      keys = append(keys, []int{})
      for c := range len(obj[0]) {
        counter := 0
        for r := 1; r < len(obj) - 1; r++ {
          if obj[r][c] == "#" { counter ++ }
        }
        keys[len(keys)-1] = append(keys[len(keys)-1], counter)
      }
    }
  }
  return keys, locks
}

func tryLocks(keys, locks [][]int) int {
  res := 0
  for _, key := range keys {
    for _, lock := range locks {
      good := true
      for i := range key {
        if key[i] + lock[i] > 5 {
          good = false
          break
        }
      }
      if good {
        res++
      }
    }
  }
  return res 
}

func partOne(filename string) int {
  input, err := getInput(filename)
  if err != nil {
    log.Fatal(err)
  }
  keys, locks := getKeysAndLocks(input)
  fmt.Println(tryLocks(keys, locks))
  return 0
} 

func main() {
  fmt.Println(partOne("./day25/test.txt"))
  fmt.Println(partOne("./day25/input.txt"))
  return
}
