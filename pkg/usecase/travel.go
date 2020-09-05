package usecase

import (
	"log"

	"github.com/kben19/smmob/pkg/common"
	"github.com/kben19/smmob/pkg/common/types"
)

const (
	logTagUsecaseTravel              = " [UsecaseTravel]"
	logTagUsecaseTravelPerformTravel = logTagUsecaseTravel + "[PerformTravel]"
)

func (u UsecaseTravel) PerformTravel(token string, apiToken string, hash string) (types.TravelPostResponse, error) {
	var header types.TravelPostHeader
	header.XCSRFToken = token
	header.ContentType = common.ContentTypeForm
	header.Origin = common.OriginSMMO
	header.Referer = common.RefererTravel
	header.SecFetchMode = "cors"

	var payload types.TravelPostPayload
	payload.Token = token
	payload.APIToken = apiToken
	payload.Hash = hash
	payload.TestData = "testdatacontent"

	response, err := u.domain.DoTravel(payload, header)
	if err != nil {
		log.Println(err.Error() + logTagUsecaseTravelPerformTravel)
		return response, err
	}
	return response, nil
}
