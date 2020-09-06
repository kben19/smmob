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
	Text                    string  `json:"text"`
	ResultText              string  `json:"resultText"`
	RewardType              string  `json:"rewardType"`
	RewardAmount            float64 `json:"rewardAmount"`
	CurrentEXP              int64   `json:"currentEXP"`
	CurrentGold             int64   `json:"currentGold"`
	Level                   int64   `json:"level"`
	NextWait                int64   `json:"nextwait"`
	IsNpc                   bool    `json:"is_npc"`
	NpcName                 string  `json:"npc_name"`
	NpcAvatar               string  `json:"npc_avatar"`
	UserAmount              int64   `json:"userAmount"`
	GuildUserAmount         int64   `json:"guild_userAmount"`
	GuildDropRatePercentage float64 `json:"guild_droprate_percentage"`
	GuildGoldIncrease       float64 `json:"guild_gold_increase"`
	GuildExpIncrease        float64 `json:"guild_exp_increase"`
}

type TravelAttackLoop struct {
	Journey []interface{}           `json:"journey"`
	Summary TravelAttackLoopSummary `json:"summary"`
}

type TravelAttackLoopSummary struct {
	TotalEXP          float64          `json:"total_exp"`
	TotalGold         float64          `json:"total_gold"`
	TotalItem         int              `json:"total_item"`
	TotalEvent        int              `json:"total_event"`
	TotalEventGold    int              `json:"total_event_gold"`
	TotalEventEXP     int              `json:"total_event_exp"`
	TotalEventItem    int              `json:"total_event_item"`
	TotalEventMonster int              `json:"total_event_monster"`
	TotalEventBattle  int              `json:"total_event_battle"`
	ItemList          []AttackItemDrop `json:"item_list"`
}
