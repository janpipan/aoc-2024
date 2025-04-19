package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPages(fileName string) (map[int]map[int]struct{}, [][]int){
  file, err := os.Open(fileName)
  if err != nil {
    return nil, nil
  }
  defer file.Close()


  scanner := bufio.NewScanner(file)
  order := map[int]map[int]struct{}{} 
  update := [][]int{}
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      break
    }
    lineSplit := strings.Split(line,"|")
    first, err := strconv.Atoi(lineSplit[0])
    if err != nil {
      return nil, nil
    }
    second, err := strconv.Atoi(lineSplit[1])
    if err != nil {
      return nil, nil
    }
    if _, ok := order[first]; ok {
      order[first][second] = struct{}{}
    } else {
      order[first] = map[int]struct{}{second: {}}
    }
  }
  for scanner.Scan() {
    line := scanner.Text()
    lineSplit := strings.Split(line, ",")
    upd := []int{}
    for _, v := range lineSplit {
      num, err := strconv.Atoi(v) 
      if err != nil {
        return nil, nil
      }
      upd = append(upd, num)
    }
    update = append(update, upd)
  }
  return order, update 
}

func isCorrect(order map[int]map[int]struct{}, page []int) (bool, int, int) {
  stack := []int{}
  for iv, v := range page {
    stack = append(stack, v)
    for is, s := range stack {
      if _, ok := order[v][s]; ok {
        return false, iv, is
      }
    }
  }
  return true, -1, -1
}


func partOne(order map[int]map[int]struct{}, update [][]int) (int, [][]int) {
  res := 0
  inc := [][]int{}

  for _, line := range update {
    good, _, _ := isCorrect(order, line)
    if good {
      res += line[int(len(line)/2)]
    } else {
      inc = append(inc, line)
    }
  }
  
  return res, inc
}

func partTwo(order map[int]map[int]struct{}, incorrect [][]int) int {
  res := 0
  for _, line := range incorrect {
    for correct, iv, is := isCorrect(order, line); !correct; correct, iv, is = isCorrect(order, line) {
      line[iv], line[is] = line[is], line[iv]
    }
    res += line[int(len(line)/2)]
  }
  return res 
}


func main() {
  order, update := getPages("./day5/test.txt")
  resTest, incorrectTest := partOne(order, update)
  fmt.Println("Result part one test:", resTest)
  fmt.Println("Result part two:", partTwo(order, incorrectTest))
  orderInput, updateInput := getPages("./day5/input.txt")
  resPartOne, incorrect := partOne(orderInput, updateInput)
  fmt.Printf("Puzzle result: %d\n", resPartOne)
  fmt.Println("Result part two:", partTwo(orderInput, incorrect))
  return
}
