package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getInput(filename string) ([][]string, error) {
  input := []string{}
  file, err := os.Open(filename)
  if err != nil {
    return nil, fmt.Errorf("Err: %w", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    input = append(input, line)
  }
  return processInput(input), nil
}

func processInput(input []string) [][]string {
  processed := [][]string{}
  for _, pair := range input {
    p := strings.Split(pair, "-")
    processed = append(processed, p)
  }
  return processed
}

func getConnections(pairs [][]string) map[string]map[string]struct{} {
  conns := map[string]map[string]struct{}{}
  for _, pair := range pairs {
    if v, ok := conns[pair[0]]; ok {
      v[pair[1]] = struct{}{}
    } else {
      conns[pair[0]] = map[string]struct{}{
        pair[1]: {},
      }
    }
    if v, ok := conns[pair[1]]; ok {
      v[pair[0]] = struct{}{}
    } else {
      conns[pair[1]] = map[string]struct{}{
        pair[0]: {},
      }
    }
  }
  return conns 
}

type InterConn struct {
  conns []string
}

func getInterConnected(conns map[string]map[string]struct{}) []string {
  interConnected := []string{}
  unique := map[string]struct{}{}

  for key := range conns {
    for k := range conns[key] {
      for k1 := range conns[k] {
        if _, ok := conns[key][k1]; ok && k1 != k {
          str := []string{key,k,k1}
          sort.Strings(str)
          s := strings.Join(str, ",")
          if _, ok := unique[s]; !ok {
            unique[s] = struct{}{}
            for _, start := range str {
              if string(start[0]) == "t" {
                interConnected = append(interConnected, s)
                break
              }
            }
          } 
        }
      }
    }
  }

  return interConnected
}

func partOne(filename string) int {
  input, err := getInput(filename)
  if err != nil {
    log.Fatal(err)
  }

  conns := getConnections(input)

  interConnected := getInterConnected(conns)

  return len(interConnected) 
}

func bronKerbosch(R, P, X []string, conns map[string]map[string]struct{}, res map[string]struct{}) {
  if len(P) == 0 && len(X) == 0 { 
    str := R
    sort.Strings(str)
    s := strings.Join(str, ",")
    if _, ok := res[s]; !ok {
      res[s] = struct{}{} 
    }
    return
  }
  for _, v := range P {
    dR := R
    dR = append(dR, v)
    dP, dX := []string{}, []string{}
    for _, p := range P {
      if _, ok := conns[v][p]; ok {
        dP = append(dP, p)
      }
    }
    for _, p := range X {
      if _, ok := conns[v][p]; ok {
        dX = append(dX, p)
      }
    }
    bronKerbosch(dR, dP, dX, conns, res) 
  }
  return  
}

func bronKerboschPivot(R, P, X []string, conns map[string]map[string]struct{}, res map[string]struct{}) {
  if len(P) == 0 && len(X) == 0 { 
    str := R
    sort.Strings(str)
    s := strings.Join(str, ",")
    if _, ok := res[s]; !ok {
      res[s] = struct{}{} 
    }
    return
  }
  pivotSet := append(P, X...)
  candidates := []string{}
  u := pivotSet[0]
  for _, v := range pivotSet {
    if _, ok := conns[u][v]; ok { continue }
    candidates = append(candidates, v)
  }
  for _, v := range candidates {
    dR := R
    dR = append(dR, v)
    dP, dX := []string{}, []string{}
    for _, p := range P {
      if _, ok := conns[v][p]; ok {
        dP = append(dP, p)
      }
    }
    for _, p := range X {
      if _, ok := conns[v][p]; ok {
        dX = append(dX, p)
      }
    }
    bronKerboschPivot(dR, dP, dX, conns, res) 
  }
  return  
}

func partTwo(filename string) string {
  input, err := getInput(filename)
  if err != nil {
    log.Fatal(err)
  }
  conns := getConnections(input)
  edges := []string{}
  for edge := range conns {
    edges = append(edges, edge)
  }
  r := map[string]struct{}{} 
  bronKerboschPivot([]string{}, edges, []string{}, conns, r)
  m, res := 0, ""
  for v := range r {
    if len(v) > m {
      res = v 
      m = len(v)
    }
  }

  return res
}

func main() {
  fmt.Println("Part one test: ", partOne("./day23/test.txt"))
  fmt.Println("Part one input: ", partOne("./day23/input.txt"))
  fmt.Println("Part two test: ", partTwo("./day23/test.txt"))
  fmt.Println("Part two input: ", partTwo("./day23/input.txt"))
  return
}
