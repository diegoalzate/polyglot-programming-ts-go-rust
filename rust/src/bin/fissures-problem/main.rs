use std::str::FromStr;

use anyhow::{Result, anyhow};

fn get_input() -> &'static str {
return "0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2";
}

#[derive(Debug)]
struct Point {
    x: i32,
    y: i32
}

#[derive(Debug)]
struct Line {
    p1: Point,
    p2: Point
}

impl FromStr for Point {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        let parts = s.split_once(",");
        if parts.is_none() {
            return Err(anyhow!("failed to parse point"));
        }
        let values: (&str, &str) = parts.unwrap();
        let x: i32 = str::parse(values.0)?;
        let y: i32 = str::parse(values.1)?;

        return Ok(Point{ x, y });
    }
}


impl FromStr for Line {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        let parts = s.split_once(" -> ");
        if parts.is_none() {
            return Err(anyhow!("failed to parse line"));
        }
        let values: (&str, &str) = parts.unwrap();
        let p1: Point = str::parse(values.0)?;
        let p2: Point = str::parse(values.1)?;

        return Ok(Line { p1, p2 });
    }
}


impl Line {
    fn has_zero_delta(&self) -> bool {
        return self.p1.x == self.p2.x || self.p1.y == self.p2.y;
    }
}

fn main() {
     let lines = get_input().lines();

     let final_lines: Vec<Line> = lines
     .flat_map(|line| str::parse(line))
     .filter(|x: &Line| x.has_zero_delta())
     .collect();

    println!("{:?}", final_lines) 
}