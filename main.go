package main

import (
	"log"

	"github.com/lioarce01/Borne-CLI/cmd"
)

func main() {
    // loadEnv()

    if err := cmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

// func loadEnv() {
//     err := godotenv.Load("./.env")
//     if err != nil {
//         log.Println("Error loading .env file")
//     }
// }
