package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var CreateBranchCmd = &cobra.Command{
	Use: "branch",
	Short: "Create a new branch",
	Run:  func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: You must provide a branch name")
			return
		}

		branchName := args[0]

		err := exec.Command("git", "branch", branchName).Run()
		if err != nil {
			log.Fatalf("Error creating branch '%s': %v", branchName, err)
		}
		fmt.Printf("Branch '%s' created successfully.\n", branchName)
	},
}

func init() {
	//No flags
}
