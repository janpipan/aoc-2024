package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// returns 0 if the reports are not safe
// returns 1 if the reports are decreasing or increasing
func isSafe(reports []int) bool {
  decreasing := 1
  if reports[0] > reports[1] {
    decreasing = -1
  }
  for i := 0; i < len(reports) - 1; i++ {
    left := decreasing * (reports[i+1] - reports[i]) 
    if left < 1 || left > 3 {
      return false 
    }
  }
  return true 
}

func isSafePartTwo(reports []int) bool {
  if isSafe(reports) {
    return true
  }
  for i := 0; i < len(reports); i++ { 
    report := slices.Clone(reports)
    report = slices.Delete(report, i, i+1) 
    if isSafe(report) {
      return true
    }
  }
  return false
}

func getReports(inputFile string) [][]int {
  res := [][]int{}
  file, err := os.Open(inputFile)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    report := []int{} 
    for _, v := range strings.Split(line, " ") {
      num, err := strconv.Atoi(v)
      if err != nil {
        panic(err)
      }
      report = append(report, num) 
    }
    res = append(res, report)
  } 
  return res 
}

func getSafeReports(inputFile string) (int, int) {
  
  res1, res2 := 0, 0
  for _, report := range getReports(inputFile) {
    if isSafe(report) {
      res1++
    }
  }
  for _, report := range getReports(inputFile) {
    if isSafePartTwo(report) {
      res2++
    }
  }

  return res1, res2
}

func main() {
  safe1, safe2 := getSafeReports("./day2/test.txt")
  fmt.Println(safe1, safe2)
  if safe1 == 2 {
    fmt.Println(getSafeReports("./day2/input.txt"))
  } else {
    fmt.Println(safe1)
  }
  if safe2 == 4 {
    fmt.Println(getSafeReports("./day2/input.txt"))
  } else {
    fmt.Println(safe2)
  }
}
