package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`;
}

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func parsePoint(point string) (*Point, error) {
	parts := strings.Split(point, ",")
	x, xError := strconv.Atoi(parts[0])
	if xError != nil {
		log.Fatal("x can not be parsed to number")
		return nil, xError
	}

	y, yError := strconv.Atoi(parts[1])
	if yError != nil {
		log.Fatal("y can not be parsed to number")
		return nil,yError
	}

	return &Point{
		x: x,
		y: y,
	}, nil
}

func parseLine(line string) (*Line, error) {
	parts := strings.Split(line, " -> ")
	p1, p1Error := parsePoint(parts[0])
	if p1Error != nil {
		log.Fatal("p1Error can not be parsed")
		return nil, p1Error
	}

	p2, p2Error := parsePoint(parts[1])
	if p2Error != nil {
		log.Fatal("p2Error can not be parsed")
		return nil,p2Error
	}

	return &Line{
		p1: *p1,
		p2: *p2,
	}, nil
}

func hasZeroDelta(p1 Point, p2 Point) bool {
	return p1.x == p2.x || p1.y == p2.y
}

func main() {
	lines := []*Line{}
	
	for _, line := range strings.Split(getInput(), "\n") {
		parsedLine, err := parseLine(line)

		if err != nil {
			panic("failed to parseLine")
		}

		if (hasZeroDelta(parsedLine.p1, parsedLine.p2)) {
			lines = append(lines, parsedLine)
		}
	}

	fmt.Print(lines)
}