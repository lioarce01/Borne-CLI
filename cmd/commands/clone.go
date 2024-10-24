package commands

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var CloneCmd = &cobra.Command{
    Use:   "clone [url]",
    Short: "Clona un repositorio de Git",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        url := args[0]

        fmt.Printf("Clonando el repositorio desde %s...\n", url)
        err := exec.Command("git", "clone", url).Run()
        if err != nil {
            fmt.Printf("Error al clonar el repositorio: %v\n", err)
            return
        }

        fmt.Println("Clonación completa.")
    },
}

func init() {
    //FLAGS
}
