package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Block struct {
  idx int
  size int
}

func getInput(fileName string) string {
  file, err := os.Open(fileName)
  if err != nil {
    return "" 
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var res string
  for scanner.Scan() {
    line := scanner.Text() 
    return line 
  }
  return res
}

func getFilesystem(input string) []int {
  i := strings.Split(input, "")
  id := 0
  filesystem := []int{}
  for i, c := range i {
    t, err := strconv.Atoi(c)
    if err != nil {
      return nil
    }
    if i % 2 == 0 {
      for j := 0; j < t; j++ {
        filesystem = append(filesystem, id)
      }
      id++
    } else {
      for j := 0; j < t; j++ {
        filesystem = append(filesystem, -1)
      }
    }
  }
  return filesystem 
}

func partOne(filesystem []int) []int {
  l, r := 0, len(filesystem)-1
  for l < r {
    if filesystem[l] == -1 {
      if filesystem[r] == -1 {
        for filesystem[r] == -1 {
          r--
        }
      }
      filesystem[l], filesystem[r] = filesystem[r], filesystem[l]
      l++
      r--
    } else {
      l++
    } 
  }
  return filesystem
}

func getFilesystemPartTwo(input string) ([]int, []Block, []Block) {
  i := strings.Split(input, "")
  id := 0
  filesystem := []int{}
  space, file := []Block{}, []Block{}
  for i, c := range i {
    t, err := strconv.Atoi(c)
    if err != nil {
      return nil, nil, nil
    }
    if i % 2 == 0 {
      file = append(file, Block{len(filesystem), t})
      for j := 0; j < t; j++ {
        filesystem = append(filesystem, id)
      }
      id++
    } else {
      space = append(space, Block{len(filesystem), t})
      for j := 0; j < t; j++ {
        filesystem = append(filesystem, -1)
      }
    }
  }
  return filesystem, file, space 
}

func partTwo(filesystem []int, file []Block, space []Block) []int {
  for r := len(file) - 1; r >= 0; r-- {
    for l, v := range space {
      if v.idx < file[r].idx && v.size >= file[r].size {
        for i := range file[r].size {
          filesystem[v.idx+i] = filesystem[file[r].idx]
        }
        for i := range file[r].size {
          filesystem[file[r].idx+i] = -1 
        }
        space[l].size -= file[r].size
        space[l].idx += file[r].size
        break  
      }
    }
  }
  return filesystem
}

func getCheckSum(filesystem []int) int {
  sum := 0
  for i := range filesystem { 
    if filesystem[i] != -1 {
      sum += i * filesystem[i]
    }
  }
  return sum
}




func main () {
  filesystem := getFilesystem(getInput("./day9/test-1.txt"))
  fmt.Println(filesystem)
  fmt.Println("First part test-1:", getCheckSum(partOne(filesystem)))
  fmt.Println("First part test-2:", getCheckSum(partOne(getFilesystem(getInput("./day9/test-2.txt")))))
  fmt.Println("First part input:", getCheckSum(partOne(getFilesystem(getInput("./day9/input.txt")))))
  filesystemPartTwo, file, space := getFilesystemPartTwo(getInput("./day9/test-2.txt"))
  fmt.Println("Second part test-2:", getCheckSum(partTwo(filesystemPartTwo, file, space)))
  fmt.Println("Second part input:", getCheckSum(partTwo(getFilesystemPartTwo(getInput("./day9/input.txt")))))
  return 
}
