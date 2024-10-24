package cmd

import (
	"borne/cmd/commands"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "borne",
	Short: "A CLI for managing your Github repositories",
	Long: `Borne is a  CLI for managing your Github repositories. It allows you to create, delete, and list repositories.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commands.PushCmd)
	rootCmd.AddCommand(commands.PullCmd)
	rootCmd.AddCommand(commands.StatusCmd)
	rootCmd.AddCommand(commands.CommitCmd)
	rootCmd.AddCommand(commands.CloneCmd)
	rootCmd.AddCommand(commands.ConfigCmd)
}


