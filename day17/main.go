package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) (map[string]int, []int) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  registers := map[string]int{} 
  program := []int{}

  for scanner.Scan() {
    line := scanner.Text()
    if len(line) == 0 { continue }
    l := strings.Split(line, " ")
    if l[0] == "Register" {
      num, _ := strconv.Atoi(l[2])
      registers[strings.Split(l[1], ":")[0]] = num 
    } else {
      nums := strings.Split(l[1], ",")
      for _, num := range nums {
        n, _ := strconv.Atoi(num)
        program = append(program, n) 
      }
    }
  }
  return registers, program
}

func runProgram(registers map[string]int, program []int) []string {
  output := []string{}
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
  registers, program := getInput(filename)
  return strings.Join(runProgram(registers, program), ",")
}



func main() {
  fmt.Println(partOne("./day17/test.txt"))
  fmt.Println(partOne("./day17/input.txt"))
  return
}
