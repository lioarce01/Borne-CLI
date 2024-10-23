import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "set-remote [name] [url]",
  describe: "Set a remote repository",
  builder: (yargs) => {
    yargs
      .positional("name", {
        describe: "Name of the remote (e.g., origin)",
        type: "string",
        demandOption: true,
      })
      .positional("url", {
        describe: "URL of the remote repository",
        type: "string",
        demandOption: true,
      });
  },
  handler: async (argv) => {
    try {
      await git.addRemote(argv.name, argv.url);
      console.log(`Remote "${argv.name}" set to ${argv.url}`);
    } catch (error) {
      console.error("Error setting remote repository:", error.message);
    }
  },
};
