package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use: "add",
	Short: "Add files to the staging area",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: You must provide at least one file to add")
			return
		}

		for _, file := range args {
			err := exec.Command("git", "add", file).Run()
			if err != nil {
				log.Fatalf("Error adding file '%s': %v", file, err)
			}
			fmt.Printf("File '%s' added to staging area\n", file)
		}
	},
}

func init() {
	//No flags
}