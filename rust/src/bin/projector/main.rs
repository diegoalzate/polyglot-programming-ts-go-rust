use clap::Parser;
use polyglot::config::ProjectorConfig;
use polyglot::opts::ProjectorOpts;

use anyhow::{Result, Ok};

fn main() -> Result<()>{
    let config: ProjectorConfig = ProjectorOpts::parse().try_into()?;
    print!("{:?}", config);

    return Ok(())
}