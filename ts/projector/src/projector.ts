import fs from "fs";
import * as path from "path";
import { Config } from "./config";

type Data = {
  // todo: if we had other top level items, we could put them here
  // such as settings or links
  projector: {
    [key: string]: {
      [key: string]: string;
    };
  };
};

type Value = string | undefined;

const DEFAULT_VALUE: Data = { projector: {} };

export class Projector {
  constructor(private config: Config, private data: Data = DEFAULT_VALUE) {}

  getValue(key: string): Value {
    // pwd
    // dirname(pwd) until empty
    let prev: Value = undefined;
    let curr = this.config.pwd;

    let out: Value = undefined;
    do {
      let val = this.data.projector[curr]?.[key];
      if (val !== undefined) {
        out = val;
        break;
      }

      prev = curr;
      curr = path.dirname(curr);
    } while (prev !== curr);

    return out;
  }

  setValue(key: string, value: string) {
    let pwd = this.config.pwd;
    if (!this.data.projector[pwd]) {
      this.data.projector[pwd] = {};
    }

    this.data.projector[pwd][key] = value;
  }

  deleteValue(key: string) {
    delete this.data.projector[this.config.pwd]?.[key];
  }

  static fromConfig(config: Config): Projector {
    let data: Data | undefined = undefined;
    try {
      if (fs.existsSync(config.config)) {
        data = JSON.parse(fs.readFileSync(config.config).toString()) as Data;
      }
    } catch {
      data = undefined;
    }

    return new Projector(config, data);
  }
}
