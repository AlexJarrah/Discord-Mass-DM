package files

import (
	"Discord-Mass-DM/internal"
	"encoding/json"
	"errors"
	"os"
)

// Parses the config file and returns a Configuration struct
func ParseConfig() error {
	// Read configuration from file
	f, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	// Parse the JSON contents of the file
	if err = json.Unmarshal(f, &internal.Config); err != nil {
		return err
	}

	return validateConfig()
}

// Ensures the provided configuration is valid
func validateConfig() error {
	if internal.Config.DiscordToken == "" {
		return errors.New("no user token provided")
	} else if len(internal.Config.MessagePool) == 0 {
		return errors.New("no messages provided in pool")
	}

	return nil
}
