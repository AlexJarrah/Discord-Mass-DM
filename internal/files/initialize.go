package files

import "os"

// Ensures all necessary files exist and are properly initialized
func InitializeFiles() error {
	// Check if the config file already exists
	if _, err := os.Stat(configFilePath); err == nil {
		return nil // File already exists, no action needed
	}

	// Create the data directory if it doesn't exist
	if err := os.Mkdir("data", 0755); err != nil {
		return err
	}

	// Create the config file with default data
	f, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write default data to the config file
	_, err = f.Write([]byte(`{"discord_token":"","message_pool":[]}`))
	return err
}
