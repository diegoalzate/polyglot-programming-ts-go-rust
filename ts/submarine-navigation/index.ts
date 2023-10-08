function getInput(): string {
  return `forward 5
down 5
forward 8
up 3
down 8
forward 2`;
}

function parseLine(line: string): [number, number] {
  const parts = line.trim().split(" ");
  const amount = +parts[1];
  const dir = parts[0];

  if (dir === "forward") {
    return [amount, 0];
  } else if (dir === "down") {
    return [0, amount];
  } else {
    return [0, -amount];
  }
}

function main() {
  return getInput()
    .split("\n")
    .map(parseLine)
    .reduce(
      (acc, curr) => {
        acc[0] += curr[0];
        acc[1] += curr[1];
        return acc;
      },
      [0, 0]
    );
}

console.log(main());
