# Borne 🐙

Welcome to Borne This project allows you to interact with Git and GitHub from the terminal in a simple and efficient way. Use this CLI to manage your repositories, branches, and commits without leaving the command line.

## Features 🌟

- **Set the working directory**
- **Create new repositories**
- **Manage branches**: create, switch, and delete
- **Make commits and view history**
- **Set remote repositories**
- **Perform `pull` and `push` operations**

## Installation 🛠️

1. Build the CLI: Open PowerShell and navigate to your project folder. Run the following command to compile the executable:

```shellscript
GOOS=windows GOARCH=amd64 go build -o brn.exe
```

2. Set up the environment variable: Create a system environment variable named GITHUB_TOKEN and assign your GitHub token to it. To do this on Windows:
   -Search for "System Properties" in the start menu.
   -Select "Advanced system settings".
   -Click on "Environment Variables".
   -In "System variables", click on "New..." and add GITHUB_TOKEN as the variable name and your GitHub token as the value.

3. Add the .exe location to the system PATH: In the same "Environment Variables" window, find the Path variable, select it, and click "Edit". Add the path where brn.exe is located.

## Usage 🚀

After installation, use the brn command followed by various subcommands to interact with Git repositories.

Available Commands

```shellscript
brn clone [repoUrl]                 Clone a repository
brn create [repoName]               Create a new Github repository
brn add [files]                     Add files to the Git index (staging area)
brn commit -m [message]             Commit changes with a message
brn push `<origin>` -b `<branch>`   Push changes to a remote repository
brn pull [url]                      Pull changes from a remote repository
brn status                          Show the status of the repository
brn log                             Show commit history
brn branch [branchName]             Create a new branch
brn switch [branchName]             Switch to a branch
brn delete [branchName]             Delete a branch
brn remote [name] [url]             Set a remote repository
brn config                          Configure Git user name and email
brn rebase                          Rebase a branch onto the specified branch
brn merge [branch]                  Merge a branch into the current branch
```

Examples

Here are some examples of how to use the brn commands:

1. Clone a repository:

```shellscript
brn clone [https://github.com/user/repo.git](https://github.com/user/repo.git) my-project
```

2. Create a new repository:

```shellscript
brn create my-new-project
```

3. Add files to staging:

```shellscript
brn add .
```

4. Commit changes:

```shellscript
brn commit "Initial commit"
```

5. Push changes:

```shellscript
brn push origin main
```

6. Create and switch to a new branch:

```shellscript
brn branch feature-branch
brn switch feature-branch
```

7. Show repository status:

```shellscript
brn status
```

8. View commit history:

```shellscript
brn log
```

9. Set a remote repository:

```shellscript
brn remote origin [https://github.com/user/repo.git](https://github.com/user/repo.git)
```

Configuration

Before using Borne, configure your Git user name and email:

```shellscript
brn config
```

You'll be prompted to enter your name and email address.

## Contribution 🤝

We welcome contributions to Borne! If you'd like to contribute:

1. Fork the repository
2. Create your feature branch (git checkout -b feature/AmazingFeature)
3. Commit your changes (git commit -m 'Add some AmazingFeature')
4. Push to the branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

Please make sure to update tests as appropriate and adhere to the project's coding standards.
