package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var MergeCmd = &cobra.Command{
	Use: "merge [branch]",
	Short: "Merge a branch into the current branch",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]

		mergeCmd := exec.Command("git", "merge", branchName)

		var out bytes.Buffer
		mergeCmd.Stdout = &out
		mergeCmd.Stderr = &out

		err :=  mergeCmd.Run()
		if err != nil {
			log.Fatalf("Error merging branch '%s': %v\nOutput %s", branchName, err, out.String())
		}

		fmt.Printf("Merged branch '%s' successfully.\n", branchName)
		fmt.Println(out.String())

		if checkMergeConflicts() {
			fmt.Println("Merge conflicts detected. Please resolve them and commit.")
		}

	},
}

func checkMergeConflicts() bool {
	statusCmd := exec.Command("git", "status", "--porcelain")
	output, err := statusCmd.Output()
	if err != nil {
		log.Fatalf("Error checking merge conflicts: %v", err)
	}

	return bytes.Contains(output, []byte("UU"))
}