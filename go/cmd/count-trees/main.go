package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

func main() {
	count := 0
	for index, line := range strings.Split(getInput(), "\n") {
		char := line[index * 3 % len(line)]
		if string(char) == "#" {
			count++;
		}
	}

	fmt.Print(count);
}