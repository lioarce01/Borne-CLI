import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "create-branch [branchName]",
  describe: "Create a new branch",
  builder: (yargs) => {
    yargs.positional("branchName", {
      describe: "Name of the new branch",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    try {
      await git.checkoutLocalBranch(argv.branchName);
      console.log(`Branch "${argv.branchName}" created successfully`);
    } catch (error) {
      console.error("Error creating branch:", error.message);
    }
  },
};
