package types

import "github.com/kben19/smmob/pkg/common"

func (p TravelPostPayload) ValidatePayload() error {
	if p.Token == "" || p.APIToken == "" || p.Hash == "" {
		return common.ErrInvalidPayload
	}
	return nil
}

func (p TravelPostPayload) Map() map[string]string {
	mapResult := map[string]string{
		"_token":       p.Token,
		"api_token":    p.APIToken,
		"testdata":     p.TestData,
		"hash_fj8n3u7": p.Hash,
	}
	return mapResult
}

func (h TravelPostHeader) ValidateHeaderTravel() error {
	err := h.ValidateHeader()
	if err != nil {
		return err
	}
	if h.XCSRFToken == "" {
		return common.ErrInvalidHeaderValue
	}
	return nil
}

func (h TravelPostHeader) MapTravel() map[string]string {
	mapResult := h.Map()
	mapResult["X-CSRF-TOKEN"] = h.XCSRFToken
	return mapResult
}
