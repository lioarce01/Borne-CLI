package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var DeleteBranchCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete a branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: You  must specify a branch name.")
			return
		}

		branchName := args[0]

		err := exec.Command("git", "branch", "-d", branchName).Run()
		if err != nil {
			log.Fatalf("Error deleting branch '%s': %v", branchName, err)
		}
		fmt.Printf("Branch '%s' successfully deleted.\n", branchName)
	},
}

func init() {
	//No flags
}