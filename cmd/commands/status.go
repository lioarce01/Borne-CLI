package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of your local repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Repository status: clean")
	},
}