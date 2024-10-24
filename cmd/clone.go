package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var CloneCmd = &cobra.Command{
    Use:   "clone [url]",
    Short: "Clona un repositorio de Git",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Error: You must provide a repository URL to clone")
            return
        }
        repoURL := args[0]

        err := exec.Command("git", "clone", repoURL).Run()
        if err != nil {
            log.Fatalf("Error cloning repository: %v", err)
        }

        fmt.Printf("Repository cloned from %s successfully.\n", repoURL)
    },
}

func init() {
    //Flags
}
