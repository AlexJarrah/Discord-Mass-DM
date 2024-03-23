package main

import (
	"Discord-Mass-DM/internal/discord"
	"Discord-Mass-DM/internal/files"
	"log"
)

func main() {
	// Ensure necessary files exist and are properly initialized
	if err := files.InitializeFiles(); err != nil {
		log.Fatalf("failed to initialize files: %v", err)
	}

	// Parse & validate the configuration file
	if err := files.ParseConfig(); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	// Start program
	if err := discord.Start(); err != nil {
		log.Fatalf("error establishing connection: %v", err)
	}
}
