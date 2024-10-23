import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "log",
  describe: "Show commit history",
  builder: () => {},
  handler: async () => {
    try {
      const log = await git.log();
      console.log("Commit history:");
      console.log(log);
    } catch (error) {
      console.error("Error retrieving commit history:", error.message);
    }
  },
};
