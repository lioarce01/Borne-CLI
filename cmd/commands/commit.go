package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short:  "Commit changes to your local repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a commit message")
			return
		}
		fmt.Printf("Committing changes with message: %s...\n", args[0])

	},

}