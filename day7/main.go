package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(fileName string) []string {
  
  file, err := os.Open(fileName)
  if err != nil {
    return nil
  }
  defer file.Close()

  equations := []string{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    equations = append(equations, line)
  }
  return equations 
}

func evalEquation(equation string) int {
  equ := strings.Split(equation, ":")
  sum, err := strconv.Atoi(equ[0])
  if err != nil {
    return 0 
  }
  nums := []int{}
  for _, n := range strings.Split(strings.Trim(equ[1], " "), " ") {
    num, err := strconv.Atoi(n)
    if err != nil {
      return 0
    }
    nums = append(nums, num)
  }
  operators := []string{"+", "*"}
  var backtrack func(s int, i int) bool
  backtrack = func(s int, i int) bool {
    if s == sum && i == len(nums) {
      return true
    } else if i == len(nums) {
      return false
    }
    for _, o := range operators {
      if o == "+" && backtrack(s+nums[i],i+1) {
        return true 
      }
      if o == "*" && backtrack(s*nums[i],i+1) {
        return true
      }
    }
    return false 
  }
  
  if backtrack(nums[0], 1) {
    return sum
  }
  return 0
}

func partOne(equations []string) int {
  sum := 0
  for _, e := range equations {
    sum += evalEquation(e)
  }
  return sum
}

func main() {
  testEquations := getInput("./day7/test.txt")
  fmt.Println(partOne(testEquations))
  inputEquations := getInput("./day7/input.txt")
  fmt.Println(partOne(inputEquations))
  return 
}
