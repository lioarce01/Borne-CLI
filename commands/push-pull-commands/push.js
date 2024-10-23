import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "push <origin> <branch>",
  describe: "Push changes to a remote repository",
  builder: (yargs) => {
    yargs.positional("origin", {
      describe: "The name of the remote repository (e.g., origin)",
      type: "string",
      demandOption: true,
    });
    yargs.positional("branch", {
      describe: "The name of the branch to push",
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
      const result = await git.push(origin, branch);
      console.log(`Changes successfully pushed to ${origin}/${branch}`);
    } catch (error) {
      console.error("Error pushing changes:", error.message);
    }
  },
};
