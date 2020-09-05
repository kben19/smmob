package types

import (
	"strconv"

	"github.com/kben19/smmob/pkg/common"
)

func (p AttackPostPayload) ValidatePayload() error {
	if p.Token == "" || p.APIToken == "" {
		return common.ErrInvalidPayload
	}
	return nil
}

func (p AttackPostPayload) Map() map[string]string {
	mapResult := map[string]string{
		"_token":         p.Token,
		"api_token":      p.APIToken,
		"special_attack": strconv.FormatBool(p.SpecialAttack),
	}
	return mapResult
}

func (h AttackPostHeader) ValidateHeaderAttack() error {
	err := h.ValidateHeader()
	if err != nil {
		return err
	}
	if h.XCSRFToken == "" {
		return common.ErrInvalidHeaderValue
	}
	return nil
}

func (h AttackPostHeader) MapAttack() map[string]string {
	mapResult := h.Map()
	mapResult["X-CSRF-TOKEN"] = h.XCSRFToken
	return mapResult
}
