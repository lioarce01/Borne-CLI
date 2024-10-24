package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "Configure Git user name and email for github",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if username == "" || email == "" {
			fmt.Println("Error: You must provide a Git user name and email")
			return
		}

		 _, err := exec.Command("git", "config", "--global", "user.name", username).Output()
		 if err != nil {
			 log.Fatalf("Error al configurar el nombre de usuario: %v\n", err)
			 return
		 }
 
		 _, err = exec.Command("git", "config", "--global", "user.email", email).Output()
		 if err != nil {
			 log.Fatalf("Error al configurar el correo electrónico: %v\n", err)
			 return
		 }

		 fmt.Printf("Git user name set to '%s' and email set to '%s'.\n", username, email)
	},
}

func init() {
    ConfigCmd.Flags().StringP("name", "n", "", "Git user name")
    ConfigCmd.Flags().StringP("email", "e", "", "Git user email")
}