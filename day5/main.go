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


func partOne(order map[int]map[int]struct{}, update [][]int) int {
  res := 0

  for _, line := range update {
    stack := []int{}
    good := true
    for _, v := range line {
      stack = append(stack, v)
      if !good {
        break
      }
      for _, s := range stack {
        if _, ok := order[v][s]; ok {
          good = false
          break
        }
      }
    }
    if good {
      res += line[int(len(line)/2)]
    }
  }
  
  return res
}



func main() {
  order, update := getPages("./day5/test.txt")
  fmt.Println(order)
  fmt.Println(update)
  fmt.Println(partOne(order, update))
  orderInput, updateInput := getPages("./day5/input.txt")
  fmt.Printf("Puzzle result: %d", partOne(orderInput, updateInput))
  return
}
