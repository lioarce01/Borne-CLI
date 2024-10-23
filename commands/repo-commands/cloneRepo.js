import path from "path";
import fs from "fs";
import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "clone [repoUrl] [dir]",
  describe: "Clone a repository",
  builder: (yargs) => {
    yargs
      .positional("repoUrl", {
        describe: "URL of the repository to clone",
        type: "string",
        demandOption: true,
      })
      .positional("dir", {
        describe: "Directory to clone the repository into",
        type: "string",
        default: ".",
      });
  },
  handler: async (argv) => {
    try {
      const destinationDir = path.resolve(argv.dir);
      const files = fs.readdirSync(destinationDir);
      if (destinationDir === process.cwd() && files.length > 0) {
        console.error("Error: The current directory is not empty.");
        return;
      }
      console.log(`Cloning repository from ${argv.repoUrl} into ${argv.dir}`);
      await git.clone(argv.repoUrl, destinationDir);
      console.log(`Repository cloned into ${destinationDir}`);
    } catch (error) {
      console.error("Error cloning repository:", error.message);
    }
  },
};
