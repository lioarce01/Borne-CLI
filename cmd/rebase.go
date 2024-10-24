package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var RebaseCmd = &cobra.Command{
	Use: "rebase [branch]",
	Short: "Rebase a branch onto the specified branch",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetBranch := args[0]

		rebaseCmd := exec.Command("git", "rebase", targetBranch)

		var out bytes.Buffer
		rebaseCmd.Stdout = &out
		rebaseCmd.Stderr = &out

		err := rebaseCmd.Run()
		if err != nil {
			log.Fatalf("Error rebasing branch onto '%s': %v\nOutput %s", targetBranch, err, out.String())
		}

		fmt.Printf("Rebased branch onto '%s' successfully.\n", targetBranch)
		fmt.Println(out.String())

		if checkRebaseConflicts() {
			fmt.Println("Rebase conflicts detected. Please resolve them and continue the rebase.")
		}
	},
}

func checkRebaseConflicts() bool {
    statusCmd := exec.Command("git", "status", "--porcelain")
    output, err := statusCmd.Output()
    if err != nil {
        log.Fatalf("Error checking rebase conflicts: %v", err)
    }

    return bytes.Contains(output, []byte("UU"))
}