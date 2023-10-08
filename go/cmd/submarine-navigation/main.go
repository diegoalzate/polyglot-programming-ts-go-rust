package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`;
}

type Point struct {
  x int
  y int
}

func parseLine(line string) Point {
  parts := strings.Split(line, " ")
  dir := parts[0]
  amount, err := strconv.Atoi(parts[1])

  if err != nil {
    // this should never happen
    panic(err)
  }
  
  if dir == "forward" {
    return Point{x: amount, y: 0};
  } else if dir == "down" {
    return Point{x: 0 ,y: amount};
  } else {
    return Point{x: 0, y: -amount};
  }
}

func main() {
  lines := strings.Split(getInput(), "\n")

  initialPoint := Point{x: 0, y: 0}
  for _, line := range lines {
    tempPoint := parseLine(line);
    initialPoint.x += tempPoint.x;
    initialPoint.y += tempPoint.y;
  }

  fmt.Printf("point: %+v", initialPoint)
}