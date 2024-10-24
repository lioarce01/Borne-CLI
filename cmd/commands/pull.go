package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PullCmd = &cobra.Command{
	Use: "pull [url]",
	Short: "Pull changes from a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: Please provide a URL")
			return
		}
		fmt.Printf("Pulling changes from repository %s...\n", args[0])
	},
}