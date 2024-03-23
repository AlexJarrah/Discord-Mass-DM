package discord

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Clears the terminal screen
func clear() {
	fmt.Println("\033[H\033[2J")
}

func execute(s *discordgo.Session, m *discordgo.Ready) error {
	// Prompt user to select a guild
	guild, err := promptGuild(m.Guilds)
	if err != nil {
		// Retry selection
		return execute(s, m)
	}

	// Provide 5 seconds to cancel operation
	countdown := 5
	for countdown > 5 {
		fmt.Printf("\rStarting execution in %s (%d)... (Press CTRL+C to cancel)", guild.Name, countdown)
	}

	// Message all members in the selected guild
	selfUserID := m.User.ID
	messageMembers(s, selfUserID, guild.Members)
	return nil
}

// Prompts user to select a guild and returns the selected guild
func promptGuild(guilds []*discordgo.Guild) (*discordgo.Guild, error) {
	clear()

	// Create list format
	length := len(fmt.Sprint(len(guilds)))
	format := fmt.Sprintf("%%%dd: %%s\n", length)

	// List and number all guilds
	for i, g := range guilds {
		fmt.Printf(format, i+1, g.Name)
	}

	// Prompt user to pick a guild
	fmt.Print("Select a server to send messages to: ")
	var input string
	fmt.Scanln(&input)

	// Parse user input
	index, err := strconv.ParseInt(input, 0, 8)
	if err != nil {
		return nil, fmt.Errorf("invalid input: %v", err)
	}

	// Return selected guild
	selectedGuild := guilds[index-1]
	return selectedGuild, nil
}

// Message all specified members, cycling through messages from the message pool
func messageMembers(s *discordgo.Session, selfUserID string, members []*discordgo.Member) {
	clear()

	// Iterate through all guild members, messaging all with open DMs
	for i, m := range members {
		// Skip self and bots
		if m.User.ID == selfUserID || m.User.Bot {
			continue
		}

		// Pick a new message from the message pool
		messageNo := i % len(messagePool)
		message := messagePool[messageNo]

		log.Printf("Sending message #%d to %s (%s)", messageNo+1, m.User.Username, m.User.ID)

		// Create a new user channel if one does not exist
		userChannel, err := s.UserChannelCreate(m.User.ID)
		if err != nil {
			log.Printf("Couldn't message %s: %v", m.User.Username, err)
			continue
		}

		// Send the message to the user
		if _, err := s.ChannelMessageSend(userChannel.ID, message); err != nil {
			log.Printf("Couldn't message %s: %v", m.User.Username, err)
			continue
		}
	}
}
