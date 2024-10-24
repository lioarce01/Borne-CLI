package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short:  "Commit changes to your local repository",
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")


		if message == "" {
			fmt.Println("Error: You must provide a commit message")
			return
		}

		err := exec.Command("git", "commit", "-m", message).Run()
		if err != nil {
			log.Fatalf("Error committing changes: %v", err)
		}

		fmt.Printf("Changes committed with message: '%s'\n", message)
	},
}

func init() {
	CommitCmd.Flags().StringP("message", "m", "", "Commit message")
}