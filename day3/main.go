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
  scanner := bufio.NewScanner(file)
  sum := 0
  for scanner.Scan() {
    line := scanner.Text() 
    multipliers := r.FindAllString(line, -1) 
    for _, v := range multipliers {
      r, _ := regexp.Compile("\\d+")
      digits := r.FindAllString(v, 2)
      //fmt.Println(digits)
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
  }
  fmt.Println(sum)
  return 
}
