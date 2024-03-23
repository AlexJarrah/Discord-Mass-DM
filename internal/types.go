package internal

type Configuration struct {
	DiscordToken string   `json:"discord_token"`
	MessagePool  []string `json:"message_pool"`
	Roles        Roles    `json:"roles"`
}

type Roles struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}
