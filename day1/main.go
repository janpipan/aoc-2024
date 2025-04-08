package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getNums(inputFile string) ([]int, []int) {
  
  file, err := os.Open(inputFile)
  if err != nil {
    panic(err) 
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  left := []int{}
  right := []int{}
  for scanner.Scan() {
    line := scanner.Text()
    val, err := strconv.Atoi(strings.Split(line, "   ")[0])
    if err != nil {
      panic(err)
    }
    left = append(left, val)
    val, err = strconv.Atoi(strings.Split(line, "   ")[1])
    if err != nil {
      panic(err)
    }
    right = append(right, val)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return left, right
}

func partOne(inputFile string) int { 
  left, right := getNums(inputFile) 
  slices.Sort(right)
  slices.Sort(left)
  sum := 0
  for i := range right {
    sum += int(math.Abs(float64(right[i] - left[i])))
  }
  return sum
}


func partTwo(inputFile string) int {
  simScore := map[int]int{}
  left, right := getNums(inputFile) 
  res := 0
  for _, val := range left {
    simScore[val] = 0
  }
  for _, val := range right {
    if _, ok := simScore[val]; ok {
      simScore[val]++
    }
  }
  for _, val := range left {
    res += val * simScore[val]
  }
  return res
}

func main() {
  if one := partOne("./test.txt"); one == 11 {
    fmt.Println(partOne("./input.txt"))
  } else {
    fmt.Println(one)
  }
  if two := partTwo("./test.txt"); two == 31 {
    fmt.Println(partTwo("./input.txt"))
  } else {
    fmt.Print(two) 
  }
  return 
}
