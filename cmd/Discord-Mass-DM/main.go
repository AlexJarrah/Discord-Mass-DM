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

	// Parse the configuration file
	config, err := files.ParseConfig()
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	// Verify configuration validity
	if err := files.ValidateConfig(config); err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	// Start program
	if err := discord.Start(config); err != nil {
		log.Fatalf("error establishing connection: %v", err)
	}
}
