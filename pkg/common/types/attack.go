package types

type AttackPostHeader struct {
	HTTPHeader
	XCSRFToken string
}

type AttackPostPayload struct {
	Token         string `json:"token"`
	APIToken      string `json:"api_token"`
	SpecialAttack bool   `json:"special_attack"`
}

type AttackPostResponse struct {
	DamageGiven     string `json:"dmg_given"`
	DamageTaken     string `json:"dmg_taken"`
	UpdatedEnemyHP  int64  `json:"updated_their_hp"`
	UpdatedPlayerHP int64  `json:"updated_your_hp"`
	EnemyDeathText  string `json:"their_death"`
	ItemDrop        string `json:"item_drop"`
	EnemyDeath      bool   `json:"they_are_dead"`
}

type AttackItemDrop struct {
	ItemURL  string `json:"item_url"`
	ItemName string `json:"item_name"`
}
