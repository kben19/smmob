package server

import (
	"net/http"
	"time"

	"github.com/kben19/smmob/pkg/domain/attack"
	"github.com/kben19/smmob/pkg/domain/travel"
	"github.com/kben19/smmob/pkg/resource/api"
	"github.com/kben19/smmob/pkg/usecase"
)

type UseCases struct {
	travel usecase.UsecaseTravelItf
}

// Init function from usecases to resource
func Init() UseCases {
	useCase := UseCases{}

	resourceAPI := api.InitResourceAPI(&http.Client{
		Timeout: 5 * time.Second,
	})

	domainTravel := travel.InitDomainTravel(resourceAPI)
	domainAttack := attack.InitDomainAttack(resourceAPI)
	useCase.travel = usecase.InitUsecaseTravel(domainTravel, domainAttack)

	return useCase
}
