package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type RepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

var CreateRepoCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: Error loading .env file; using environment variables instead")
		}
		
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("Error: GITHUB_TOKEN environment variable is not set")
		}

		repoName, _ := cmd.Flags().GetString("name")
		isPrivate, _ := cmd.Flags().GetBool("private")
		description, _ := cmd.Flags().GetString("description")

		if repoName == "" {
			fmt.Println("Error: You must provide a repository name")
			return
		}

		repo := RepoRequest{
			Name:        repoName,
			Description: description,
			Private:     isPrivate,
		}

		jsonData, err := json.Marshal(repo)
		if err != nil {
			log.Fatalf("Error converting payload to JSON: %v", err)
		}

		url := "https://api.github.com/user/repos"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatalf("Error creating the request: %v", err)
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making the HTTP request: %v", err)
		}
		defer resp.Body.Close()

		// Manejar la respuesta
		if resp.StatusCode == http.StatusCreated {
			fmt.Printf("Repository '%s' created successfully.\n", repoName)
		} else {
			var responseBody map[string]interface{}
			json.NewDecoder(resp.Body).Decode(&responseBody)

			fmt.Printf("Error creating the repository. Status code: %d\n", resp.StatusCode)
			if message, exists := responseBody["message"]; exists {
				fmt.Printf("GitHub message: %s\n", message)
			}
		}
	},
}

func init() {
	CreateRepoCmd.Flags().StringP("name", "n", "", "New repository name")
	CreateRepoCmd.Flags().BoolP("private", "p", false, "Make the repository private")
	CreateRepoCmd.Flags().StringP("description", "d", "", "Repository description (optional)")
}