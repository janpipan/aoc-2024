package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
  y int
  x int
  yvel int
  xvel int
}

func getInput(filename string) []Robot {
  robots:= []Robot{}
  file, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  r, _ := regexp.Compile("[-]?\\d+") 
  for scanner.Scan() {
    line := scanner.Text()
    matches := r.FindAllString(line, -1)
    nums := []int{}
    for _, m := range matches {
      num, _ := strconv.Atoi(m)
      nums = append(nums, num)
    }
    robots = append(robots, Robot{nums[1], nums[0], nums[3], nums[2]}) 
  }

  return robots
}


func simulate(robot Robot, yMax, xMax, times int) int {
  for i := 0; i < times; i++ {
    dy, dx := robot.y + robot.yvel, robot.x + robot.xvel
    if dy < 0 {
      robot.y = yMax + dy
    } else {
      robot.y = dy % yMax
    }
    if dx < 0 {
      robot.x = xMax + dx
    } else {
      robot.x = dx % xMax
    }
  }
  ybound := (yMax - 1) / 2
  xbound := (xMax - 1) / 2 
  if robot.y < ybound && robot.x < xbound {
    return 1
  } else if robot.y < ybound && robot.x > xbound {
    return 2
  } else if robot.y > ybound && robot.x < xbound {
    return 3
  } else if robot.y > ybound && robot.x > xbound {
    return 4
  }
  return 0
}

func partOne(filename string, yMax, xMax, times int) int {
  robots := getInput(filename)
  counter := map[int]int{
    1: 0,
    2: 0,
    3: 0,
    4: 0,
  }
  for _, robot := range robots {
    counter[simulate(robot, yMax, xMax, times)]++ 
  }
  res := 1
  fmt.Println(counter)
  for k, v := range counter {
    if k != 0 {
      res *= v 
    } 
  }
  return res
}

func main() {
  fmt.Println(partOne("./day14/test.txt", 7, 11, 100))
  fmt.Println(partOne("./day14/input.txt", 103,101, 100))  
  return 
}
