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

type Sequence struct {
  n1 int
  n2 int
  n3 int
  n4 int
}

func calculateSecretNumTwo(num, steps int, seqMap map[Sequence]int) {
  q := []int{}
  for range steps {
    diff := num % 10
    s1 := ((num * 64) ^ num) % 16777216
    s2 := ((s1 / 32) ^ s1) % 16777216
    s3 := ((s2 * 2048) ^ s2) % 16777216
    num = s3
    diff = (num % 10) - diff
    q = append(q, diff)
    if len(q) == 4 {
      seq := Sequence{q[0],q[1],q[2],q[3]}
      if _, ok := seqMap[seq]; !ok {
        seqMap[seq] = num % 10
      } 
      q = q[1:]
    }
  }
}

func partTwo(filename string) int {
  bananas := 0
  input := getInput(filename)
  scores := map[Sequence]int{}
  for _, num := range input {
    seqMap := map[Sequence]int{}
    calculateSecretNumTwo(num, 2000, seqMap)
    for key, val := range seqMap {
      scores[key] += val
    }
  } 
  for _, val := range scores {
    bananas = max(bananas, val)
  }
  return  bananas
}

func main() {
  fmt.Println("Part one test: ", partOne("./day22/test.txt"))
  fmt.Println("Part one input: ", partOne("./day22/input.txt"))
  fmt.Println("Part two test: ", partTwo("./day22/test-2.txt"))
  fmt.Println("Part two input: ", partTwo("./day22/input.txt"))
  return
}
