use std::path::PathBuf;

use clap::Parser;

#[derive(Debug)]
#[derive(Parser)]
#[clap()]
pub struct Opts {
    #[clap(default_value = "")]    
    pub args: Vec<String>,
    #[clap(short = 'c', long = "config")]
    pub config: Option<PathBuf>,
    #[clap(short = 'p', long = "pwd")]
    pub pwd: Option<PathBuf>,
}

