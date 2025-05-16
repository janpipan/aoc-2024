package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput(filename string) []int {
  input := []int{}

  file, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    num, _ := strconv.Atoi(line)
    input = append(input, num)
  }
  
  return input
}

func calculateSecretNum(num, steps int) int {
  for range steps {
    s1 := ((num * 64) ^ num) % 16777216
    s2 := ((s1 / 32) ^ s1) % 16777216
    s3 := ((s2 * 2048) ^ s2) % 16777216
    num = s3
  }
  return num 
}

func partOne(filename string) int {
  input := getInput(filename)
  res := 0
  for _, num := range input {
    res += calculateSecretNum(num, 2000) 
  }
  return res 
}

func main() {
  fmt.Println("Part one test: ", partOne("./day22/test.txt"))
  fmt.Println("Part one input: ", partOne("./day22/input.txt"))
  return
}
