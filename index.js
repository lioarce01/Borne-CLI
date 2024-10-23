import yargs from "yargs/yargs";
import { hideBin } from "yargs/helpers";
import dotenv from "dotenv";

dotenv.config();

import clone from "./commands/repo-commands/cloneRepo.js";
import createRepo from "./commands/repo-commands/createRepo.js";
import add from "./commands/commit-commands/add.js";
import commit from "./commands/commit-commands/commit.js";
import push from "./commands/push-pull-commands/push.js";
import pull from "./commands/push-pull-commands/pull.js";
import status from "./commands/commit-commands/status.js";
import log from "./commands/commit-commands/log.js";
import createBranch from "./commands/branch-commands/createBranch.js";
import switchBranch from "./commands/branch-commands/switchBranch.js";
import deleteBranch from "./commands/branch-commands/deleteBranch.js";
import setRemote from "./commands/config-commands/setRemote.js";
import setDir from "./commands/config-commands/setDir.js";
import configUser from "./commands/config-commands/configUser.js";

yargs(hideBin(process.argv))
  .command(clone)
  .command(createRepo)
  .command(add)
  .command(commit)
  .command(push)
  .command(pull)
  .command(status)
  .command(log)
  .command(createBranch)
  .command(switchBranch)
  .command(deleteBranch)
  .command(setRemote)
  .command(setDir)
  .command(configUser)
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
