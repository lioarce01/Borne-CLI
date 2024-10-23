import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "switch-branch [branchName]",
  describe: "Switch to a branch",
  builder: (yargs) => {
    yargs.positional("branchName", {
      describe: "Name of the branch to switch to",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    try {
      await git.checkout(argv.branchName);
      console.log(`Switched to branch "${argv.branchName}".`);
    } catch (error) {
      console.error("Error switching branches:", error.message);
    }
  },
};
