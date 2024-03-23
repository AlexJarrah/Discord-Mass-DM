package files

import (
	"Discord-Mass-DM/internal"
	"encoding/json"
	"os"
)

// Ensures all necessary files exist and are properly initialized
func InitializeFiles() error {
	// Check if the config file already exists
	if _, err := os.Stat(configFilePath); err == nil {
		return nil // File already exists, no action needed
	}

	// Create the data directory if it doesn't exist
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if err := os.Mkdir("data", 0755); err != nil {
			return err
		}
	}

	// Create the config file with default data
	f, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Define default configuration
	defaultConfig := internal.Configuration{
		DiscordToken: "",
		MessagePool:  []string{},
		Roles: internal.Roles{
			Include: []string{"*"},
			Exclude: []string{""},
		},
	}

	// Marshal the config with indentation
	jsonData, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON to the config file
	_, err = f.Write(jsonData)
	return err
}
