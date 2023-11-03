use clap::Parser;

fn main() {
    let opts = polyglot::opts::Opts::parse();
    print!("{:?}", opts)

}