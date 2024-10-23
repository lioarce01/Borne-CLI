# Git CLI 🐙

Welcome to the Git CLI! This project allows you to interact with Git and GitHub from the terminal in a simple and efficient way. Use this CLI to manage your repositories, branches, and commits without leaving the command line.

## Features 🌟

- **Set the working directory**
- **Create new repositories**
- **Manage branches**: create, switch, and delete
- **Make commits and view history**
- **Set remote repositories**
- **Perform `pull` and `push` operations**


## Installation 🛠️

1. Clone the repository:

```shellscript
git clone <repository_url>
```


2. Navigate to the project directory:

```shellscript
cd <directory_name>
```


3. Install dependencies:

```shellscript
npm install
```


4. Create a .env file and add your GitHub token:

```plaintext
GITHUB_TOKEN=your_github_token
```




## Usage 🚀

1. Set the working directory

```shellscript
node lioarce01.js set-dir <directory>
```

Description: Sets the directory where Git commands will be executed.


2. Clone a repository

```shellscript
node lioarce01.js clone <repository_url> [directory]
```

Description: Clones a repository into the specified directory or the current directory if none is provided.


3. Create a new repository

```shellscript
node lioarce01.js create-repo <repository_name>
```

Description: Creates a new repository in the current working directory.


4. Create a new branch

```shellscript
node lioarce01.js create-branch <branch_name>
```

Description: Creates a new branch in the current repository.


5. Switch branches

```shellscript
node lioarce01.js switch-branch <branch_name>
```

Description: Switches to the specified branch.


6. Show files pending commit

```shellscript
node lioarce01.js status
```

Description: Displays the files that are ready to be committed.


7. Make a commit

```shellscript
node lioarce01.js commit "<commit_message>"
```

Description: Commits all modified files with the provided message.


8. Show commit history

```shellscript
node lioarce01.js log
```

Description: Displays the commit history of the current repository.


9. Delete a branch

```shellscript
node lioarce01.js delete-branch <branch_name>
```

Description: Deletes the specified branch.


10. Set a remote repository

```shellscript
node lioarce01.js set-remote <name> <repository_url>
```

Description: Sets a remote repository with the specified name and URL.


11. Pull changes

```shellscript
node lioarce01.js pull
```

Description: Fetches and merges changes from the remote repository into the current directory.


12. Push changes

```shellscript
node lioarce01.js push
```

Description: Sends local commits to the remote repository.


13. Exit

```shellscript
node lioarce01.js exit
```

Description: Exits the program.




## Contribution 🤝

If you would like to contribute, please open an issue or submit a pull request.
