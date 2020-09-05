package usecase

import (
	"github.com/kben19/smmob/pkg/common/types"
	"github.com/kben19/smmob/pkg/domain/attack"
	"github.com/kben19/smmob/pkg/domain/travel"
)

type UsecaseTravel struct {
	domainTravel travel.DomainTravelItf
	domainAttack attack.DomainAttackItf
}

type UsecaseTravelItf interface {
	PerformTravel(token string, apiToken string, hash string) (types.TravelPostResponse, error)
	PerformTravelAndAttack(token string, apiToken string, hash string, specialAttack bool) ([]interface{}, error)
}

func InitUsecaseTravel(domainTravel travel.DomainTravelItf, domainAttack attack.DomainAttackItf) UsecaseTravelItf {
	return &UsecaseTravel{
		domainTravel: domainTravel,
		domainAttack: domainAttack,
	}
}
