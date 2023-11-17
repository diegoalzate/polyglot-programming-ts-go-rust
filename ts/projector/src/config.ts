import * as path from "path";
import { ProjectorOptions } from "./opts";

export enum Operation {
  Print,
  Add,
  Remove,
}

export type Config = {
  pwd: string;
  config: string;
  operation: Operation;
  args: string[];
};

function getPwd(opts: ProjectorOptions): string {
  if (opts.pwd) {
    return opts.pwd;
  }

  return process.cwd();
}

function getConfig(opts: ProjectorOptions): string {
  if (opts.config) {
    return opts.config;
  }

  return path.join(
    process.env.XDG_CONFIG_HOME ?? "",
    "projector",
    "projector.json"
  );
}

function getOperation(opts: ProjectorOptions): Operation {
  if (opts.arguments?.[0] === "add") {
    return Operation.Add;
  }

  if (opts.arguments?.[0] === "remove") {
    return Operation.Remove;
  }

  return Operation.Print;
}

function getArgs(opts: ProjectorOptions): string[] {
  if (!opts.arguments) {
    return [];
  }

  const operation = getOperation(opts);

  if (operation === Operation.Add) {
    if (opts.arguments.length !== 3) {
      throw new Error(
        `add expects 2 arguments and received ${opts.arguments.length - 1}`
      );
    }

    return opts.arguments.slice(1);
  }

  if (operation === Operation.Print) {
    if (opts.arguments.length !== 2) {
      throw new Error(
        `print expects 1 argument and received ${opts.arguments.length - 1}`
      );
    }

    return [opts.arguments[1]];
  }

  if (opts.arguments.length !== 2) {
    throw new Error(
      `remove expects 1 argument and received ${opts.arguments.length - 1}`
    );
  }

  // remove operation
  return [opts.arguments[1]];
}

/**
 *
 * Transforms opts into config
 */
export default function config(opts: ProjectorOptions): Config {
  return {
    pwd: getPwd(opts),
    args: getArgs(opts),
    config: getConfig(opts),
    operation: getOperation(opts),
  };
}
