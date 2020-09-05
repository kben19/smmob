package travel

import (
	"encoding/json"
	"github.com/kben19/smmob/pkg/common"
	"log"

	"github.com/kben19/smmob/pkg/common/types"
)

const (
	logTagDomainTravel         = " [DomainTravel]"
	logTagDomainTravelDoTravel = logTagDomainTravel + "[DoTravel]"

	UrlDoTravel = common.OriginAPISMMO + "/travel/perform"
)

func (d DomainTravel) DoTravel(payload types.TravelPostPayload, header types.TravelPostHeader) (types.TravelPostResponse, error) {
	var response types.TravelPostResponse
	err := payload.ValidatePayload()
	if err != nil {
		log.Println(err.Error() + logTagDomainTravelDoTravel)
		return response, err
	}
	err = header.ValidateHeaderTravel()
	if err != nil {
		log.Println(err.Error() + logTagDomainTravelDoTravel)
		return response, err
	}

	bytes, err := d.rscAPI.PostHTTPRequest(UrlDoTravel, nil, header.MapTravel(), payload.Map())
	if err != nil {
		log.Println(err.Error() + logTagDomainTravelDoTravel)
		return response, err
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Println(err.Error() + logTagDomainTravelDoTravel)
		return response, err
	}
	return response, nil
}
