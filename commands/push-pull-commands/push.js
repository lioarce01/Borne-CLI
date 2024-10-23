import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "push",
  describe: "Push changes to a remote repository",
  builder: () => {
    yargs.positional("origin", {
      describe: "repository to push to",
      type: "string",
      demandOption: true,
    });
    yargs.positional("branch", {
      describe: "branch to push",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    const { origin, branch } = argv;

    try {
      if (!origin || !branch) {
        console.error("Error: origin and branch are required");
      }
      await git.push(origin, branch);
      console.log("Changes pushed to a remote repository");
    } catch (error) {
      console.error("Error pushing changes:", error.message);
    }
  },
};
