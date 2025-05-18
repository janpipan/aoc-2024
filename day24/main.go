package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput(filename string) ([]string, []string, error) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, nil, fmt.Errorf("Err : %w", err) 
  }  
  defer file.Close()

  scanner := bufio.NewScanner(file)
inputs, gates := []string{}, []string{}

  for scanner.Scan() {
    line := scanner.Text()
    if len(line) == 0 { continue }
    if len(line) < 8 {
      inputs = append(inputs, line)
    } else {
      gates = append(gates, line)
    }

  }
  return inputs, gates, nil 
}

func processInput(inputs []string) map[string]int {
  inp := map[string]int{}
  for _, input := range inputs {
    str := strings.Split(input, ": ")
    if str[1] == "0" { inp[str[0]] = 0 } else { inp[str[0]] = 1 }
  }
  return inp
}

func getDecimal(setBits map[string]int, gates []string) int64 {
  for len(gates) > 0 {
    curr := gates[0]
    line := strings.Split(curr, " ")
    gates = gates[1:]
    if _, ok := setBits[line[0]]; !ok { 
      gates = append(gates, curr)
      continue 
    }
    if _, ok := setBits[line[2]]; !ok { 
      gates = append(gates, curr)
      continue 
    }
    if line[1] == "XOR" {
      setBits[line[4]] = setBits[line[0]] ^ setBits[line[2]]
    } else if line[1] == "AND" {
      setBits[line[4]] = setBits[line[0]] & setBits[line[2]]
    } else if line[1] == "OR" {
      setBits[line[4]] = setBits[line[0]] | setBits[line[2]]
    }
  }
  r := []string{}
  for key:= range setBits {
    if string(key[0]) == "z" {
      r = append(r, key)
    }
  }

  res := make([]string, len(r))
  for _, val := range r {
    num, _ := strconv.Atoi(val[1:])
    v := strconv.Itoa(setBits[val])
    res[len(res)-1-num] = v
  }
  fmt.Println(strings.Join(res, ""))
  n, _ := strconv.ParseInt(strings.Join(res, ""), 2, 64)
  return n 
}

func partOne(filename string) int64 {
  inputs, gates, err := getInput(filename)
  if err != nil {
    log.Fatal(err)
  }
  input := processInput(inputs)
  res := getDecimal(input, gates)
  return res
}


func getWrongGates(gates []string) []string {
  wrong := []string{}
  xorPossible := []string{}
  andPossible := []string{}
  xorInputs := map[string]struct{}{}
  orInputs := map[string]struct{}{}
  g := gates
  for len(g) > 0 {
    curr := g[0]
    line := strings.Split(curr, " ")
    a, gate, b, output := line[0], line[1], line[2], line[4]
    if gate == "XOR" {
      xorInputs[a] = struct{}{}
      xorInputs[b] = struct{}{}
    }
    if gate == "OR" {
      orInputs[a] = struct{}{}
      orInputs[b] = struct{}{}
    }
    g = g[1:]
    if output[0] == 'z' && output != "z45" && gate != "XOR" {
      wrong = append(wrong, curr)
    } 
    if output[0] != 'z' && a[0] != 'x' && a[0] != 'y' && b[0] != 'x' && b[0] != 'y' {
      if gate == "XOR" {
        wrong = append(wrong, curr)
      }
    }
    if (a[0] == 'x' && b[0] == 'y') || (b[0] == 'x' && a[0] == 'y') {
      if (a == "x00" && b == "y00") || (a == "y00" && b == "x00") { continue }
      if gate == "XOR" {
        xorPossible = append(xorPossible, curr)
      }
      if gate == "AND" {
        andPossible = append(andPossible, curr)
      }
    }
  }
  for _, line := range xorPossible {
    output := strings.Split(line, " ")[4]
    if _, ok := xorInputs[output]; !ok {
      wrong = append(wrong, line)
    }
  }
  for _, line := range andPossible {
    output := strings.Split(line, " ")[4]
    if _, ok := orInputs[output]; !ok {
      wrong = append(wrong, line)
    }
  }
  return wrong 
} 

func partTwo(filename string) string {
  _, gates, err := getInput(filename)
  if err != nil {
    log.Fatal(err)
  }
  gate := getWrongGates(gates)
  dist := map[string]struct{}{}
  for _, g := range gate {
    if _, ok := dist[g]; !ok {
      dist[g] = struct{}{}
    }
  }
  res := []string{}
  for k := range dist {
    k := strings.Split(k, " ")
    res = append(res, k[4])
  }
  sort.Strings(res) 
  return strings.Join(res, ",") 
}

func main() {
  fmt.Println(partOne("./day24/test-1.txt"))
  fmt.Println(partOne("./day24/test-2.txt"))
  fmt.Println(partOne("./day24/input.txt"))
  fmt.Println(partTwo("./day24/input.txt"))
  return
}
