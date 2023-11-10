use std::{path::PathBuf, env};
use anyhow::{Result, Error, Ok, anyhow};

use crate::opts::ProjectorOpts;

#[derive(Debug)]
pub enum Operation {
    Print(Option<String>),
    Add((String, String)),
    Remove(String),
}

#[derive(Debug)]
pub struct ProjectorConfig {
    pub pwd: Result<PathBuf>,
    pub config: Result<PathBuf>,
    // you do not need args because they are already part of the operation
    pub operation: Operation
}

impl TryFrom<Vec<String>> for Operation {
    type Error = Error;    
    
    fn try_from(mut value: Vec<String>) -> Result<Self, Self::Error> {
        if value.len() == 0 {
            return Ok(Operation::Print(None))
        }

        let term = value.get(0).expect("expect 1 parameter to exist");

        if term == "add" {
            if value.len() != 3 {
                return Err(anyhow!("add expects 2 arguments and received {:?}", value.len() - 1));
            }

            let mut args = value.drain(1..=2);
            return Ok(Operation::Add((args.next().unwrap(), args.next().unwrap())))
        }

        if term == "remove" {
            if value.len() != 2 {
                return Err(anyhow!("remove expects 1 argument and received {:?}", value.len() - 1));
            }

            let arg = value.pop().expect("remove expects 1 argument");

            return Ok(Operation::Remove(arg))
         }


         // assume it is print
         if value.len() != 2 {
            return Err(anyhow!("print expects 1 or 0 arguments and received {:?}", value.len()  -1))
         }

         let arg = value.pop().expect("print expects 1 or 0 arguments");

        return Ok(Operation::Print(Some(arg)))
    }
}

fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(value) = pwd {
        return Ok(value)
    }

    let cwd = std::env::current_dir().expect("expect a current working directory");

    return Ok(cwd);
}

fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(value) = config {
        return Ok(value)
    }

    let home_dir = env::var("HOME").expect("Expect home environment");
    let mut config_dir = PathBuf::from(home_dir);
    config_dir.push("projector");
    config_dir.push("projector.json");
    return Ok(config_dir)
 }

pub fn get_projector_config(opts: ProjectorOpts) -> Result<ProjectorConfig> {
    return Ok(ProjectorConfig {
        operation: opts.args.try_into()?,
        config: get_config(opts.config),
        pwd: get_pwd(opts.pwd)
    })
}
