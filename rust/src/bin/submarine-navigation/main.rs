fn get_input() -> &'static str {
  return "forward 5
down 5
forward 8
up 3
down 8
forward 2";
}

#[derive(Debug)]
struct Point {
    x: i32,
    y: i32
}

fn parse_line(line: &str) -> Point {
  let parts = line.split_once(" ").expect("Expect string to have 2 parts");
  let dir: &str = parts.0;
  let amount: i32 = str::parse(parts.1).expect("Expect the parse to go through");

  if dir == "forward" {
    return Point{x: amount, y: 0};
  } else if dir == "down" {
    return Point{x: 0 ,y: amount};
  } else {
    return Point{x: 0, y: -amount};
  }
}

fn main() {
  let lines = get_input().lines();
  let final_point = lines
    .map(parse_line)
    .fold(Point{x: 0, y: 0}, |mut acc, point| {
      acc.x += point.x;
      acc.y += point.y;
      return acc
    });

    println!("{:?}", final_point)
}
