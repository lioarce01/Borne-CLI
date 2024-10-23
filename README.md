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

1. Install Borne globally using npm:

```shellscript
npm install -g borne
```

## Usage 🚀

After installation, use the brn command followed by various subcommands to interact with Git repositories.

Available Commands

```shellscript
brn clone [repoUrl] [dir]       Clone a repository
brn create-repo [repoName]      Create a new repository
brn add [files]                 Add files to the Git index (staging area)
brn commit [message]            Commit changes with a message
brn push `<origin>` `<branch>`      Push changes to a remote repository
brn pull                        Pull changes from a remote repository
brn status                      Show the status of the repository
brn log                         Show commit history
brn create-branch [branchName]  Create a new branch
brn switch-branch [branchName]  Switch to a branch
brn delete-branch [branchName]  Delete a branch
brn set-remote [name] [url]     Set a remote repository
brn set-dir [directory]         Set working directory for Git commands
brn config-user                 Configure Git user name and email
brn exit                        Exit the program
```

Examples

Here are some examples of how to use the brn commands:

1. Clone a repository:

```shellscript
brn clone [https://github.com/user/repo.git](https://github.com/user/repo.git) my-project
```

2. Create a new repository:

```shellscript
brn create-repo my-new-project
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
brn create-branch feature-branch
brn switch-branch feature-branch
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
brn set-remote origin [https://github.com/user/repo.git](https://github.com/user/repo.git)
```

10. Set working directory:

```shellscript
brn set-dir /path/to/your/project
```

Configuration

Before using Borne, configure your Git user name and email:

```shellscript
brn config-user
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

## [License](LICENSE)

Your feedback and contributions help make Borne better for everyone. Don't hesitate to reach out!

Happy coding with Borne! May your repositories always be in order and your commits crystal clear. 🚀🐙
