package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
) 

var SwitchBranchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: You must provide a branch name to switch to")
			return
		}

		branchName := args[0]

		err := exec.Command("git", "checkout", branchName).Run()
		if err != nil {
			log.Fatalf("Error switching to branch '%s': %v", branchName, err)
		}
		fmt.Printf("Switched to branch '%s' successfully\n", branchName)
	},
}

func init() {
	//No flags
}