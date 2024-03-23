package files

import (
	"Discord-Mass-DM/internal"
	"encoding/json"
	"errors"
	"os"
)

// Parses the config file and returns a Configuration struct
func ParseConfig() (internal.Configuration, error) {
	// Read configuration from file
	f, err := os.ReadFile(configFilePath)
	if err != nil {
		return internal.Configuration{}, err
	}

	// Parse the JSON contents of the file
	configuration := internal.Configuration{}
	if err = json.Unmarshal(f, &configuration); err != nil {
		return internal.Configuration{}, err
	}

	return configuration, nil
}

// Ensures the provided configuration is valid
func ValidateConfig(config internal.Configuration) error {
	if config.DiscordToken == "" {
		return errors.New("no user token provided")
	} else if len(config.MessagePool) == 0 {
		return errors.New("no messages provided in pool")
	}

	return nil
}
