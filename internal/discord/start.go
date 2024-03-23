package discord

import (
	"Discord-Mass-DM/internal"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Establish Discord connection via the provided token
func Start() error {
	// Create a Discord session
	dg, err := discordgo.New(internal.Config.DiscordToken)
	if err != nil {
		return err
	}

	// Specify required intents for the session
	dg.Identify.Intents = discordgo.IntentsAll

	// Update the session's user agent
	dg.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"

	dg.AddHandler(ready) // Add a ready handler

	// Open a websocket connection
	if err := dg.Open(); err != nil {
		return err
	}
	defer dg.Close()

	// Wait for signal to exit
	select {}
}

// Handler function that starts when a successful Discord connection has been made
func ready(s *discordgo.Session, m *discordgo.Ready) {
	// Retry on error
	for {
		err := execute(s, m)
		if err == nil {
			log.Println("All operations complete successfully. Press CTRL+C to exit.")
			return
		}

		log.Printf("execution error: %v", err)
		time.Sleep(2 * time.Second)
	}
}
