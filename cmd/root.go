package cmd

import (
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "brn",
	Short: "Borne is a CLI for managing your Github repositories",
	Long: `Borne is a  CLI for managing your Github repositories. It allows you to create, delete, and list repositories.`,
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(CloneCmd)
	rootCmd.AddCommand(ConfigCmd)
	rootCmd.AddCommand(CommitCmd)
	rootCmd.AddCommand(CreateRepoCmd)
	rootCmd.AddCommand(PullCmd)
	rootCmd.AddCommand(PushCmd)
	rootCmd.AddCommand(StatusCmd)
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(RemoteCmd)
	rootCmd.AddCommand(LogCmd)
	rootCmd.AddCommand(CreateBranchCmd)
	rootCmd.AddCommand(SwitchBranchCmd)
	rootCmd.AddCommand(DeleteBranchCmd)
	rootCmd.AddCommand(MergeCmd)
	rootCmd.AddCommand(RebaseCmd)
}