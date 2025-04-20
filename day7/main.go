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

func parseEquation(equation string) (int, []int) {
  equ := strings.Split(equation, ":")
  sum, err := strconv.Atoi(equ[0])
  if err != nil {
    return 0, nil 
  }
  nums := []int{}
  for _, n := range strings.Split(strings.Trim(equ[1], " "), " ") {
    num, err := strconv.Atoi(n)
    if err != nil {
      return 0, nil
    }
    nums = append(nums, num)
  }
  return sum, nums 
}

func evalEquation(sum int, nums []int) int {

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
func evalEquationTwo(sum int, nums []int) int {

  operators := []string{"+", "*", "||"}
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
      if o == "||" {
        first := strconv.Itoa(s)
        second := strconv.Itoa(nums[i])
        conc := first + second
        c, err := strconv.Atoi(conc)
        if err != nil {
          return false
        }
        if backtrack(c, i+1){
          return true
        }
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
    s, n := parseEquation(e)
    sum += evalEquation(s, n) 
  }
  return sum
}
func partTwo(equations []string) int {
  sum := 0
  for _, e := range equations {
    s, n := parseEquation(e)
    sum += evalEquationTwo(s, n) 
  }
  return sum
}



func main() {
  testEquations := getInput("./day7/test.txt")
  fmt.Println("Part one test:", partOne(testEquations))
  inputEquations := getInput("./day7/input.txt")
  fmt.Println("Part one result:", partOne(inputEquations))
  fmt.Println("Part two test:", partTwo(testEquations))
  fmt.Println("Part two result:", partTwo(inputEquations))
  return 
}
