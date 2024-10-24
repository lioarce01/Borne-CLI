package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var RemoteCmd = &cobra.Command{
	Use: "remote",
	Short: "Add a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Error: You must provide a remote name and URL")
			return
		}

		remoteName := args[0]
		remoteURL := args[1]

		err := exec.Command("git", "remote", "add", remoteName, remoteURL).Run()
		if err != nil {
			log.Fatalf("Error adding remote repository: '%s': %v", remoteName,err)
		}
		fmt.Printf("Remote repository '%s' added with URL '%s'.\n", remoteName, remoteURL)
	},
}

func init() {
	//No flags
}