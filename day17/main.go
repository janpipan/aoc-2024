package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInput(filename string) (map[string]int, []int, string) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, nil, "" 
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  registers := map[string]int{} 
  program := []int{}
  var s string

  for scanner.Scan() {
    line := scanner.Text()
    if len(line) == 0 { continue }
    l := strings.Split(line, " ")
    if l[0] == "Register" {
      num, _ := strconv.Atoi(l[2])
      registers[strings.Split(l[1], ":")[0]] = num 
    } else {
      s = l[1]
      nums := strings.Split(l[1], ",")
      for _, num := range nums {
        n, _ := strconv.Atoi(num)
        program = append(program, n) 
      }
    }
  }
  return registers, program, s
}

func runProgram(registers map[string]int, program []int, num int) []string {
  output := []string{}
  if num > 0 {
    registers["A"] = num  
  }
  for i := 0; i < len(program) - 1; {
    inst, literal := program[i], program[i+1]
    combo := literal
    switch combo {
    case 4:
      combo = registers["A"]
    case 5:
      combo = registers["B"]
    case 6:
      combo = registers["C"]
    }
    switch inst {
    case 0:
      registers["A"] = int(registers["A"] / int(math.Pow(2,float64(combo))))
    case 1:
      registers["B"] = registers["B"] ^ literal 
    case 2:
      registers["B"] = combo % 8
    case 3:
      if registers["A"] == 0 { 
        i += 2
        continue 
      }
      i = literal 
      continue
    case 4:
      registers["B"] = registers["C"] ^ registers["B"]
    case 5:
      out := strconv.Itoa(combo % 8)
      output = append(output, out) 
    case 6:
      registers["B"] = int(registers["A"] / int(math.Pow(2,float64(combo))))
    case 7:
      registers["C"] = int(registers["A"] / int(math.Pow(2,float64(combo))))
    }
    i += 2
  }
  return output 
}


func partOne(filename string) string {
  registers, program, _ := getInput(filename)
  return strings.Join(runProgram(registers, program, 0), ",")
}


func partTwo(filename string) int {
  register, program, _ := getInput(filename)
  possible := []int{}
  sub := []int{0}
  for i := len(program) - 1; i >= 0; i-- {
    inst := strconv.Itoa(program[i])
    possible = sub
    sub = []int{}
    for len(possible) > 0 {
      p := possible[0]
      possible = possible[1:]

      p *= 8

      for num := range 8 {
        res := runProgram(register, program, p + num)
        if len(res) > 0 && res[0] == inst {
          sub = append(sub, p+num)
        }
      }
    }
  }
  return slices.Min(sub) 
}

// A = 8
//2,4 B = A % 8  B = 0 
//1,1 B = B ^ 1  B = 1 
//7,5 C = A // 2^B C = 4 
//4,4 B = B ^ C B = 5
//1,4 B = B ^ 4 B = 1 
//0,3 A = A // 2^3 A = 1
//5,5 out B 
//3,0 end if A == 0 


func main() {
  fmt.Println(partOne("./day17/test.txt"))
  fmt.Println(partOne("./day17/input.txt"))
  fmt.Println(partTwo("./day17/test.txt"))
  fmt.Println(partTwo("./day17/input.txt"))
  return
}
