package attack

import (
	"encoding/json"
	"log"

	"github.com/kben19/smmob/pkg/common"
	"github.com/kben19/smmob/pkg/common/types"
)

const (
	logTagDomainAttack         = " [DomainAttack]"
	logTagDomainAttackDoAttack = logTagDomainAttack + "[DoAttack]"
)

func (d DomainAttack) DoAttack(payload types.AttackPostPayload, header types.AttackPostHeader, path string) (types.AttackPostResponse, error) {
	var response types.AttackPostResponse
	err := payload.ValidatePayload()
	if err != nil {
		log.Println(err.Error() + logTagDomainAttackDoAttack)
		return response, err
	}
	err = header.ValidateHeaderAttack()
	if err != nil {
		log.Println(err.Error() + logTagDomainAttackDoAttack)
		return response, err
	}

	url := common.OriginAPISMMO + path
	bytes, err := d.rscAPI.PostHTTPRequest(url, nil, header.MapAttack(), payload.Map())
	if err != nil {
		log.Println(err.Error() + logTagDomainAttackDoAttack)
		return response, err
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Println(err.Error() + logTagDomainAttackDoAttack)
		return response, err
	}
	return response, nil
}
