package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of your local repository",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gitStatus(); err != nil {
			cmd.Println("Error executing 'status' command: ", err)
			return
		}

		cmd.Println("Repository is clean")
	},
}

func gitStatus() error {
	cmd := exec.Command("git", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}