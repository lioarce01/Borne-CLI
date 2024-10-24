package main

import (
	"log"

	"github.com/lioarce01/Borne-CLI/cmd"
)

func main() {

    if err := cmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
