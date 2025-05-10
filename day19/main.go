package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func getInput(filename string) (map[string]struct{}, []string) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  towels := map[string]struct{}{}
  patterns := []string{}

  i := 0
  for scanner.Scan() {
    line := scanner.Text()
    if i == 0 {
      t := strings.Split(line, ",")
      for _, towel := range t {
        towels[strings.Trim(towel, " ")] = struct{}{}
      }
      i++
    } else if len(line) > 0 {
      patterns = append(patterns, line)
    }
  }
  return towels, patterns 
}

func isPossible(pattern string, towels map[string]struct{}) bool {
  if pattern == "" { return true }
  for i := 1; i < len(pattern)+1; i++ {
    if _, ok := towels[pattern[:i]]; ok && isPossible(pattern[i:], towels) { 
      return true 
    }
  }
  return false
}

func partOne(filename string) int {
  towels, patterns := getInput(filename)
  res := 0
  for _, pattern := range patterns {
    if isPossible(pattern, towels) { res++ }
  }
  return res 
}

func main() {
  fmt.Println(partOne("./day19/test.txt"))
  fmt.Println(partOne("./day19/input.txt"))
  return 
}
