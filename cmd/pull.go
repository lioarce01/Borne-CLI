package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var PullCmd = &cobra.Command{
	Use: "pull [url]",
	Short: "Pull changes from a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gitPull(); err != nil {
			cmd.Println("Error pulling changes: ", err)
			return
		}

		cmd.Println("Changes pulled successfully from remote repository")

	},
}

func gitPull() error {
	cmd := exec.Command("git", "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}