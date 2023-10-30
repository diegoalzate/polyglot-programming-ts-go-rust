import cli from "command-line-args";

export type ProjectorOptions = {
  pwd?: string; // projector --pwd ...
  config?: string; // projector --config ...
  arguments?: string[];
};

export default function getOptions(): ProjectorOptions {
  return cli([
    { name: "config", type: String },
    { name: "pwd", type: String },
    { name: "arguments", type: String, defaultOption: true, multiple: true },
  ]) as ProjectorOptions;
}
