package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use: "log",
	Short: "Show the commit logs",
	Run: func(cmd *cobra.Command, args []string) {
		var out bytes.Buffer
		command := exec.Command("git", "log")
		command.Stdout = &out
		err := command.Run()
		if err != nil {
			log.Fatalf("Error getting log: %v", err)
		}

		fmt.Println(out.String())
	},
}

func init() {
	//No flags
}