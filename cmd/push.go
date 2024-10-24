package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)


var branch string


var PushCmd = &cobra.Command{
	Use: "push [url]",
	Short: "Push a local branch to a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gitPush(branch); err != nil {
			cmd.Println("Error pushing branch: ", err)
			return
		}
		cmd.Println("Branch pushed successfully to branch:", branch)
	},
}

func gitPush(branch string) error {
	cmd := exec.Command("git", "push", "origin", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	PushCmd.Flags().StringVarP(&branch, "branch", "b", "main", "Specify the branch to push")
}