import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "delete-branch",
  describe: "Delete a branch",
  builder: (yargs) => {
    yargs.positional("branchName", {
      describe: "Name of the branch to delete",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    try {
      await git.deleteLocalBranch(argv.branchName);
      console.log(`Branch "${argv.branchName}" deleted successfully`);
    } catch (error) {
      console.error("Error deleting branch:", error.message);
    }
  },
};
