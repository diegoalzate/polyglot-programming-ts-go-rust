fn get_input() -> &'static str {
	return "..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#"
}

fn main() {
    let lines = get_input().lines();
    let count = lines
    .enumerate()
    .flat_map(|(idx, line)| line.chars().nth(idx * 3 % line.len()))
    .filter(|char| char == &'#')
    .count();

    println!("{:?}", count) 
}