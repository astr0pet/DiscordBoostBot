package Utils

type guildin struct {
	Type      int    `json:"type"`
	Code      string `json:"code"`
	ExpiresAt any    `json:"expires_at"`
	Guild     struct {
		ID                       string   `json:"id"`
		Name                     string   `json:"name"`
		Splash                   any      `json:"splash"`
		Banner                   any      `json:"banner"`
		Description              any      `json:"description"`
		Icon                     any      `json:"icon"`
		Features                 []string `json:"features"`
		VerificationLevel        int      `json:"verification_level"`
		VanityURLCode            string   `json:"vanity_url_code"`
		NsfwLevel                int      `json:"nsfw_level"`
		Nsfw                     bool     `json:"nsfw"`
		PremiumSubscriptionCount int      `json:"premium_subscription_count"`
	} `json:"guild"`
	GuildID string `json:"guild_id"`
	Channel struct {
		ID   string `json:"id"`
		Type int    `json:"type"`
		Name string `json:"name"`
	} `json:"channel"`
	ApproximateMemberCount   int `json:"approximate_member_count"`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
}

type WebsocketResponse struct {
	D struct {
		SessionID string `json:"session_id,omitempty"`
	} `json:"d,omitempty"`
}

type ConfigFile struct {
	Proxyless       bool   `json:"proxyless"`
	Timeout         int    `json:"timeout"`
	CapService      string `json:"capService"`
	CapKey          string `json:"capKey"`
	InviteFieldName string `json:"inviteFieldName"`
	Port            string `json:"port"`
	DiscordSettings struct {
		Token         string   `json:"token"`
		BotStatus     string   `json:"botStatus"`
		BotActivity   string   `json:"botActivity"`
		SupportServer string   `json:"supportServer"`
		GuildID       string   `json:"guildID"`
		Owners        []string `json:"owners"`
		EmbedColor    string   `json:"embedColor"`
		LogsChannel   string   `json:"logsChannel"`
	} `json:"discordSettings"`
	CustomPersonalization struct {
		Onliner      bool     `json:"onliner"`
		DisplayName  string   `json:"displayName"`
		CustomBio    string   `json:"customBio"`
		CustomPfp    []string `json:"customPfp"`
		CustomBanner []string `json:"customBanner"`
		Status       []string `json:"status"`
		StatusEmoji  []string `json:"statusEmoji"`
	} `json:"customPersonalization"`
}

// Example function that takes guildin struct as input
func ProcessGuildData(data guildin) {
	// Process guild data here
}

// Example function that returns a guildin struct
func CreateGuildData() guildin {
	// Create and return guild data
	return guildin{}
}
