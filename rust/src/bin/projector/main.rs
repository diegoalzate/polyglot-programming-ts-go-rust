use clap::Parser;

fn main() {
    let opts = polyglot::opts::ProjectorOpts::parse();
    print!("{:?}", opts);

    let config = polyglot::config::get_projector_config(opts);
    print!("{:?}", config.expect(""));

}