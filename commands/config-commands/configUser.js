import simpleGit from "simple-git";

const git = simpleGit();

export default {
  command: "config-user",
  describe: "Configure Git user name and email",
  builder: (yargs) => {
    return yargs
      .option("name", {
        describe: "Git user name",
        type: "string",
        demandOption: true,
      })
      .option("email", {
        describe: "Git user email",
        type: "string",
        demandOption: true,
      });
  },
  handler: async (argv) => {
    const { name, email } = argv;

    try {
      await git.addConfig("user.name", name);
      await git.addConfig("user.email", email);

      console.log(`Git user configured: \nName: ${name}\nEmail: ${email}`);
    } catch (error) {
      console.error("Error configuring Git user:", error.message);
    }
  },
};
