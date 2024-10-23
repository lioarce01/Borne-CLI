import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "commit [message]",
  describe: "Commit changes with a message",
  builder: (yargs) => {
    yargs.positional("message", {
      describe: "Commit message",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    try {
      await git.commit(argv.message);
      console.log(`Changes committed with message: "${argv.message}"`);
    } catch (error) {
      console.error("Error committing changes:", error.message);
    }
  },
};
