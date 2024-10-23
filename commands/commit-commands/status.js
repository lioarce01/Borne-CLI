import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "status",
  describe: "Show the status of the repository",
  builder: () => {},
  handler: async () => {
    try {
      const status = await git.status();
      console.log("Repository status:");
      console.log(
        status.isClean() === true ? "Clean" : "Modified files:",
        status.modified
      );
    } catch (error) {
      console.error("Error retrieving repository status:", error.message);
    }
  },
};
