package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)






func main() {
  file, err := os.Open("./day3/input.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()  

  r, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")
  r2, _ := regexp.Compile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
  scanner := bufio.NewScanner(file)
  sum := 0
  sum2 := 0
  enabled := true
  for scanner.Scan() {
    line := scanner.Text() 
    multipliers := r.FindAllString(line, -1) 
    dos := r2.FindAllString(line, -1)
    // part one
    for _, v := range multipliers {
      r, _ := regexp.Compile("\\d+")
      digits := r.FindAllString(v, 2)
      digit1, err := strconv.Atoi(digits[0])
      if err != nil {
        panic(err)
      }
      digit2, err := strconv.Atoi(digits[1])
      if err != nil {
        panic(err)
      }
      sum += digit1 * digit2
    }
    // part two
    for _, v := range dos {
      if enabled && v != "don't()" &&  v != "do()" {
        r, _ := regexp.Compile("\\d+")
        digits := r.FindAllString(v, 2)
        digit1, err := strconv.Atoi(digits[0])
        if err != nil {
          panic(err)
        }
        digit2, err := strconv.Atoi(digits[1])
        if err != nil {
          panic(err)
        }
        sum2 += digit1 * digit2
      }
      if v == "don't()" {
        enabled = false
      }
      if v == "do()" {
        enabled = true
      }
    }
  }
  fmt.Println(sum)
  fmt.Println(sum2)
  return 
}
