export default {
  command: "set-dir [directory]",
  describe: "Set working directory for Git commands",
  builder: (yargs) => {
    yargs.positional("directory", {
      describe: "Path to the directory",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    try {
      const dir = path.resolve(argv.directory);
      if (!fs.existsSync(dir)) {
        console.error("The specified directory does not exist.");
        return;
      }

      process.chdir(argv.directory);
      console.log(`Working directory changed to: ${process.cwd()}`);
    } catch (error) {
      console.error("Error setting working directory", error.message);
    }
  },
};
