import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "push",
  describe: "Push changes to a remote repository",
  builder: () => {},
  handler: async () => {
    try {
      await git.push("origin", "main");
      console.log("Changes pushed to a remote repository");
    } catch (error) {
      console.error("Error pushing changes:", error.message);
    }
  },
};
