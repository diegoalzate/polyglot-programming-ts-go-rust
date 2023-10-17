function getInput(): string {
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
.#..#...#.#`;
}

const tree = "#";

function main() {
  let count = 0;
  const lines = getInput().split("\n");
  let index = 0;
  for (let line of lines) {
    const char = line.at((index * 3) % lines.length);
    if (char === tree) count++;
    index++;
  }

  return count;
}

console.log(main());
