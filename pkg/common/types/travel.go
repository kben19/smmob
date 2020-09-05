package types

type TravelPostHeader struct {
	HTTPHeader
	XCSRFToken string
}

type TravelPostPayload struct {
	Token    string `json:"token"`
	APIToken string `json:"api_token"`
	TestData string `json:"test_data"`
	Hash     string `json:"hash"`
}

type TravelPostResponse struct {
	Text                    string `json:"text"`
	ResultText              string `json:"resultText"`
	RewardType              string `json:"rewardType"`
	RewardAmount            int64  `json:"rewardAmount"`
	CurrentEXP              int64  `json:"currentEXP"`
	CurrentGold             int64  `json:"currentGold"`
	Level                   int64  `json:"level"`
	NextWait                int64  `json:"nextwait"`
	IsNpc                   bool   `json:"is_npc"`
	NpcName                 string `json:"npc_name"`
	NpcAvatar               string `json:"npc_avatar"`
	UserAmount              int64  `json:"userAmount"`
	GuildUserAmount         int64  `json:"guild_userAmount"`
	GuildDropRatePercentage int    `json:"guild_droprate_percentage"`
	GuildGoldIncrease       int    `json:"guild_gold_increase"`
	GuildExpIncrease        int    `json:"guild_exp_increase"`
}
