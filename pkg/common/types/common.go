package types

import "github.com/kben19/smmob/pkg/common"

type HTTPHeader struct {
	ContentType  common.ContentType
	Origin       string
	Referer      string
	SecFetchMode string
}
