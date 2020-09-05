package usecase

import (
	"github.com/kben19/smmob/pkg/common"
	"github.com/kben19/smmob/pkg/common/types"
	"log"
	"regexp"
	"strings"
	"time"
)

const (
	logTagUsecaseTravel              = " [UsecaseTravel]"
	logTagUsecaseTravelPerformTravel = logTagUsecaseTravel + "[PerformTravel]"
	logTagUsecaseTravelAttack        = logTagUsecaseTravel + "[TravelAttack]"
)

func (u UsecaseTravel) PerformTravel(token string, apiToken string, hash string) (types.TravelPostResponse, error) {
	header := types.TravelPostHeader{
		HTTPHeader: buildHeaderPost(common.RefererTravel),
		XCSRFToken: token,
	}

	payload := types.TravelPostPayload{
		Token:    token,
		APIToken: apiToken,
		TestData: "testdatacontent",
		Hash:     hash,
	}

	response, err := u.domainTravel.DoTravel(payload, header)
	if err != nil {
		log.Println(err.Error() + logTagUsecaseTravelPerformTravel)
		return response, err
	}
	return response, nil
}

func (u UsecaseTravel) PerformTravelAndAttack(token string, apiToken string, hash string, specialAttack bool) ([]interface{}, error) {
	var allResponse []interface{}
	response, err := u.PerformTravel(token, apiToken, hash)
	if err != nil {
		return nil, err
	}
	allResponse = append(allResponse, response)

	// No monster
	if !strings.Contains(response.Text, common.AttackPath) {
		return allResponse, nil
	}

	// If monster exist, attack until dead
	var responseAtk types.AttackPostResponse
	for !responseAtk.EnemyDeath {
		time.Sleep(1 * time.Second)

		reg := regexp.MustCompile(common.AttackPath + `.*\?`)
		fullAttackPath := reg.FindString(response.Text)
		if fullAttackPath == "" {
			return allResponse, nil
		}
		fullReferer := common.OriginSMMO + fullAttackPath
		header := types.AttackPostHeader{
			HTTPHeader: buildHeaderPost(fullReferer + "new_page=true"),
			XCSRFToken: token,
		}
		payload := types.AttackPostPayload{
			Token:         token,
			APIToken:      apiToken,
			SpecialAttack: specialAttack,
		}

		responseAtk, err = u.domainAttack.DoAttack(payload, header, fullAttackPath[:len(fullAttackPath)-1])
		if err != nil {
			log.Println(err.Error() + logTagUsecaseTravelAttack)
			return nil, err
		}
		allResponse = append(allResponse, responseAtk)
		if responseAtk.UpdatedPlayerHP <= 0 {
			break
		}
	}

	return allResponse, nil
}

func buildHeaderPost(referer string) types.HTTPHeader {
	return types.HTTPHeader{
		ContentType:  common.ContentTypeForm,
		Origin:       common.OriginSMMO,
		Referer:      referer,
		SecFetchMode: "cors",
	}
}
