package usecase

import (
	"github.com/kben19/smmob/pkg/common/types"
	"github.com/kben19/smmob/pkg/domain/travel"
)

type UsecaseTravel struct {
	domain travel.DomainTravelItf
}

type UsecaseTravelItf interface {
	PerformTravel(token string, apiToken string, hash string) (types.TravelPostResponse, error)
}

func InitUsecaseTravel(domainTravel travel.DomainTravelItf) UsecaseTravelItf {
	return &UsecaseTravel{domain: domainTravel}
}
