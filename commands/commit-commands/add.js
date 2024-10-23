import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "add [files]",
  describe: "Add files to the Git index (staging area)",
  builder: (yargs) => {
    yargs.positional("files", {
      describe: "Files to add (use '.' to add all files)",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    try {
      await git.add(argv.files);
      console.log(`Files added: ${argv.files}`);
    } catch (error) {
      console.error("Error adding files:", error.message);
    }
  },
};
