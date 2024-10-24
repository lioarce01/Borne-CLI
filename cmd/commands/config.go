package commands

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "Configure user and email for github",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")

		if username == "" || email == "" {
			fmt.Println("Please provide both username and email")
			return
		}

		 _, err := exec.Command("git", "config", "--global", "user.name", username).Output()
		 if err != nil {
			 fmt.Printf("Error al configurar el nombre de usuario: %v\n", err)
			 return
		 }
 
		 _, err = exec.Command("git", "config", "--global", "user.email", email).Output()
		 if err != nil {
			 fmt.Printf("Error al configurar el correo electrónico: %v\n", err)
			 return
		 }

		fmt.Printf("Configuring user %s and email %s...\n", username, email)
	},
}

func init() {
    ConfigCmd.Flags().StringP("username", "u", "", "Nombre de usuario de GitHub")
    ConfigCmd.Flags().StringP("email", "e", "", "Correo electrónico de GitHub")
}