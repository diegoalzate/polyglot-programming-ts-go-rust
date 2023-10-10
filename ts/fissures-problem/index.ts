function getInput() {
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

type Point = {
  x: number;
  y: number;
};

type Line = {
  p1: Point;
  p2: Point;
};

function parsePoint(point: string): Point {
  const [x, y] = point.split(",");
  return {
    x: +x,
    y: +y,
  };
}

function parseLine(line: string): Line {
  const [p1, p2] = line.split(" -> ");
  const parsedP1 = parsePoint(p1);
  const parsedP2 = parsePoint(p2);

  return {
    p1: parsedP1,
    p2: parsedP2,
  };
}

function hasZeroDelta(p1: Point, p2: Point) {
  return p1.x === p2.x || p1.y === p2.y;
}

function main() {
  const lines = getInput().split("\n");
  const res = lines
    .map(parseLine)
    .filter((line) => hasZeroDelta(line.p1, line.p2));

  console.log(res);
}

main();
