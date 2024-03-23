package internal

type Configuration struct {
	DiscordToken string   `json:"discord_token"`
	MessagePool  []string `json:"message_pool"`
}
