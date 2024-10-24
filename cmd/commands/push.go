package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)


var branch string


var PushCmd = &cobra.Command{
	Use: "push [url]",
	Short: "Push a local branch to a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: must provide a URL")
			return
		}
		fmt.Printf("Pushing changes to repository %s in branch %s...\n", args[0], branch)
	},
}

func init() {
	PushCmd.Flags().StringVarP(&branch, "branch", "b", "main", "Specify the branch to push")
}