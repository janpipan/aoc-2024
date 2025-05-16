package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func getInput(filename string) []string {
	file, err := os.Open(filename)
  if err != nil {
    log.Panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  res := []string{}

  for scanner.Scan() {
    line := scanner.Text()
    res = append(res, line)
  }
  return res
}

type Coordinate struct {
  y int
  x int
}

type Position struct {
  start string
  end string
}

type Combination struct {
  cord Coordinate
  sequence string
}

var numKeypad [][]string = [][]string{
  {"7", "8", "9"},
  {"4", "5", "6"},
  {"1", "2", "3"},
  {"-1", "0", "A"},
}

var dirKeypad [][]string = [][]string{
  {"-1", "^", "A"},
  {"<", "v", ">"},
}

func computeSequences(keypad [][]string) map[Position][]string {
  pos := map[string]Coordinate{} 
  for i := range keypad {
    for j := range keypad[i] {
      if keypad[i][j] != "-1" { pos[keypad[i][j]] = Coordinate{i,j}}
    }
  }
  seqs := map[Position][]string{}
  for p1 := range pos {
    for p2 := range pos {
      if p1 == p2 {
        seqs[Position{p1,p2}] = []string{"A"}
        continue
      }
      possibilities := []string{}
      q := []Combination{{pos[p1], ""}}
      optimal := 100
      for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        comb, seq := curr.cord, curr.sequence 
        found := false
        for _, c := range []Combination{{Coordinate{comb.y-1, comb.x}, "^"}, {Coordinate{comb.y+1, comb.x}, "v"}, {Coordinate{comb.y, comb.x-1}, "<"}, {Coordinate{comb.y, comb.x+1}, ">"}} {
          if c.cord.y < 0 || c.cord.x < 0 || c.cord.y >= len(keypad) || c.cord.x >= len(keypad[0]) { continue } 
          if keypad[c.cord.y][c.cord.x] == "-1" { continue } 
          if keypad[c.cord.y][c.cord.x] == p2 {
            if optimal < len(seq) + 1 { 
              found = true
              break 
            }
            optimal = len(seq) + 1
            possibilities = append(possibilities, seq + c.sequence + "A")
          } else {
            q = append(q, Combination{c.cord, seq + c.sequence})
          }
        }
        if !found {
          continue
        }
        break
      }
      seqs[Position{p1,p2}] = possibilities
    }
  }
  return seqs
}

func solve(input string, sequences map[Position][]string) []string {
  return combs(sequences, input) 
}

func combs(sequences map[Position][]string, input string) []string {
  text := "A" + input
  options := [][]string{}
  for i := 0; i < len(text) - 1; i++ {
    c1, c2 := text[i], text[i+1]
    options = append(options, []string{})
    for _, seq := range sequences[Position{string(c1),string(c2)}] {
      options[len(options)-1] = append(options[len(options)-1], seq)
    }
  }
  lenOptions := len(options)
  for n := 0; n < lenOptions; n++ {
    if len(options) > 1 {
      o1, o2 := options[0], options[1]
      sub := []string{}
      for _, i := range o1 {
        for _, j := range o2 {
          sub = append(sub, i + j)
        }
      }
      options = options[1:]
      options[0] = sub
    }
  }
  return options[0] 
}

func partOne(filename string) int {
  input := getInput(filename)
  res := 0
  numSequences := computeSequences(numKeypad)
  dirSequences := computeSequences(dirKeypad)
  for _, i := range input {
    firstRobot := solve(i, numSequences)
    //fmt.Println(firstRobot)
    m := math.MaxInt
    m2 := math.MaxInt
    for _, sol := range firstRobot {
      //fmt.Println(sol)
      secondRobot := solve(sol, dirSequences)
      for _, sol2 := range secondRobot {
        if len(sol2) < m {
          m = len(sol2)
        }
      }
      for _, sol2 := range secondRobot {
        if len(sol2) == m {
          person := solve(sol2, dirSequences)
          for _, p := range person {
            if len(p) < m2 {
              m2 = len(p)
            }
          }
        }
      }
    }
    num, _ := strconv.Atoi(i[:len(i)-1])
    res += m2 * num 
  }
  return res
}

type CacheComb struct {
  start string
  end string
  depth int
}

var cache = map[CacheComb]int{} 

var numSequences = computeSequences(numKeypad)
var dirSequences = computeSequences(dirKeypad)

func recSolve(start, end string, depth, maxDepth int) int {
  res := 0
  if val, ok := cache[CacheComb{start,end,depth}]; ok {
    return val
  }
  var seqs []string
  if depth == 1 {
    seqs = numSequences[Position{start,end}]
  } else {
    seqs = dirSequences[Position{start,end}]
  }
  if depth == maxDepth {
    m := math.MaxInt
    for _, seq := range seqs {
      m = min(m, len(seq))
    }
    return m
  } else { 
    m := math.MaxInt
    for _, seq := range seqs {
      sum := 0
      seq = "A" + seq
      for i := 0; i < len(seq) -1; i++ {
        sum += recSolve(string(seq[i]), string(seq[i+1]), depth+1, maxDepth)
      }
      m = min(m, sum)
    }
    res = m
  }
  cache[CacheComb{start,end,depth}] = res
  return res 
}

func partTwo(filename string) int {
  input := getInput(filename)
  result := 0
  
  for _, inp := range input {
    num, _ := strconv.Atoi(inp[:len(inp)-1])
    inp = "A" + inp 
    res := 0
    for i := 0; i < len(inp) - 1; i++ {
      r := recSolve(string(inp[i]), string(inp[i+1]), 1, 26)
      res += r 
    }
    result += res * num 
  }
  return result
}




func main() {
  fmt.Println("Part one test: ", partOne("./day21/test.txt"))
  fmt.Println("Part one input: ", partOne("./day21/input.txt"))
  fmt.Println("Part two test: ", partTwo("./day21/test.txt"))
  fmt.Println("Part two input: :", partTwo("./day21/input.txt"))
  return
}
