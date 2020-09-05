package types

import "github.com/kben19/smmob/pkg/common"

func (h HTTPHeader) ValidateHeader() error {
	if h.ContentType.String() == "" || h.Origin == "" || h.Referer == "" || h.SecFetchMode == "" {
		return common.ErrInvalidHeaderValue
	}
	return nil
}

func (h HTTPHeader) Map() map[string]string {
	mapResult := map[string]string{
		common.HeaderContentType:  h.ContentType.String(),
		common.HeaderOrigin:       h.Origin,
		common.HeaderReferer:      h.Referer,
		common.HeaderSecFetchMode: h.SecFetchMode,
	}
	return mapResult
}
