package travel

import (
	"github.com/kben19/smmob/pkg/common/types"
	"github.com/kben19/smmob/pkg/resource/api"
)

type DomainTravel struct {
	rscAPI api.ResourceAPIItf
}

type DomainTravelItf interface {
	DoTravel(payload types.TravelPostPayload, header types.TravelPostHeader) (types.TravelPostResponse, error)
}

func InitDomainTravel(rscAPI api.ResourceAPIItf) DomainTravelItf {
	return &DomainTravel{rscAPI: rscAPI}
}
