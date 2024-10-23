import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "pull",
  describe: "Pull changes from a remote repository",
  builder: () => {},
  handler: async () => {
    try {
      await git.pull("origin", "main");
      console.log("Changes pulled from a remote repository");
    } catch (error) {
      console.error("Error pulling changes:", error.message);
    }
  },
};
