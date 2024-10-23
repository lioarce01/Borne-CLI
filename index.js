import yargs from "yargs/yargs";
import { hideBin } from "yargs/helpers";
import dotenv from "dotenv";
import simpleGit from "simple-git";
import path from "path";
import fs from "fs";

dotenv.config();
const TOKEN = process.env.GITHUB_TOKEN;
let git = simpleGit();

import setDir from "./commands/setDir.js";
import clone from "./commands/clone.js";
import commit from "./commands/commit.js";
import branch from "./commands/branch.js";

yargs(hideBin(process.argv))
  .command(setDir)
  .command(clone)
  .command(commit)
  .command(branch)
  //MAS COMANDOS ACA
  .command(
    "exit",
    "Exit the program",
    () => {},
    () => {
      console.log("Exiting the program. Goodbye!");
      process.exit(0);
    }
  )
  .help().argv;
