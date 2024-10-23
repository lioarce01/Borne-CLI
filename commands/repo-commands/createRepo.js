import { Octokit } from "@octokit/rest";

const TOKEN = process.env.GITHUB_TOKEN;
const octokit = new Octokit({ auth: TOKEN });

export default {
  command: "create-repo [repoName]",
  describe: "Create a new repository",
  builder: (yargs) => {
    yargs.positional("repoName", {
      describe: "Name of the new repository",
      type: "string",
      demandOption: true,
    });
  },
  handler: async (argv) => {
    try {
      const response = await octokit.repos.createForAuthenticatedUser({
        name: argv.repoName,
        private: false,
      });
      console.log(`Repository ${argv.repoName} created successfully`);
      console.log(`URL: ${response.data.html_url}`);
    } catch (error) {
      console.error("Error creating repository:", error.message);
    }
  },
};
