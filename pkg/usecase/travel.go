package usecase

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/kben19/smmob/pkg/common"
	"github.com/kben19/smmob/pkg/common/types"
	"github.com/kben19/smmob/pkg/common/utils"
)

const (
	logTagUsecaseTravel              = " [UsecaseTravel]"
	logTagUsecaseTravelPerformTravel = logTagUsecaseTravel + "[PerformTravel]"
	logTagUsecaseTravelAttack        = logTagUsecaseTravel + "[TravelAttack]"
	logTagUsecaseTravelAttackLoop    = logTagUsecaseTravel + "[TravelAttackLoop]"

	DefaultTimeTravel = 15
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

		fullAttackPath := utils.RegexFindMultipleString([]string{common.AttackPath + `.*\?`}, response.Text)[0]
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

func (u UsecaseTravel) PerformTravelAndAttackLoop(token string, apiToken string, hash string, specialAttack bool, loop int) (types.TravelAttackLoop, error) {
	if loop > 50 {
		loop = 50
	}
	var result types.TravelAttackLoop
	var summary types.TravelAttackLoopSummary
	var allResponse []interface{}
	var timeDelay int
	for i := 0; i < loop; i++ {
		itemRaw := ""
		time.Sleep(time.Duration(timeDelay) * time.Second)
		responses, err := u.PerformTravelAndAttack(token, apiToken, hash, specialAttack)
		if err != nil {
			log.Println(err.Error() + logTagUsecaseTravelAttackLoop)
			result.Journey = allResponse
			result.Summary = summary
			return result, err
		}
		allResponse = append(allResponse, responses...)
		if len(responses) == 0 {
			timeDelay = DefaultTimeTravel
			continue
		}
		travelResp, ok := responses[0].(types.TravelPostResponse)
		if !ok {
			timeDelay = DefaultTimeTravel
			continue
		}

		// Update Summary
		summary.TotalEvent += 1
		if travelResp.RewardType == "exp" {
			summary.TotalEXP += travelResp.RewardAmount
			summary.TotalEventEXP += 1
		}
		if travelResp.RewardType == "gold" {
			summary.TotalGold += travelResp.RewardAmount
			summary.TotalEventGold += 1
		}
		if travelResp.RewardType == "item" {
			summary.TotalEventItem += 1
			itemRaw = travelResp.ResultText
		}
		// Update Summary Monster Battle(s)
		if len(responses) > 1 {
			attackResp, ok := responses[len(responses)-1].(types.AttackPostResponse)
			if ok {
				summary.TotalEventBattle += len(responses) - 1
				summary.TotalEventMonster += 1
				if attackResp.EnemyDeath {
					//Get EXP and Gold from monster battle
					resultRegex := utils.RegexFindMultipleString([]string{`\d+\sexperience`, `\d+\sgold coins`}, attackResp.EnemyDeathText)
					if len(resultRegex) == 2 && resultRegex[0] != "" && resultRegex[1] != "" {
						exp, err := strconv.ParseFloat(strings.Replace(resultRegex[0], " experience", "", -1), 64)
						if err != nil {
							log.Println(err.Error() + logTagUsecaseTravelAttackLoop)
						} else {
							summary.TotalEXP += exp
						}
						gold, err := strconv.ParseFloat(strings.Replace(resultRegex[1], " gold coins", "", -1), 64)
						if err != nil {
							log.Println(err.Error() + logTagUsecaseTravelAttackLoop)
						} else {
							summary.TotalGold += gold
						}
					}
					//Get Item Drop from monster battle
					itemRaw = attackResp.ItemDrop
				}
			}
		}
		// Parse item raw message
		if itemRaw != "" {
			itemRegex := utils.RegexFindMultipleString([]string{`[a-zA-Z\s0-9]*</span>`, `img src=\".+?\"`}, itemRaw)
			if len(itemRegex) == 2 && itemRegex[0] != "" && itemRegex[1] != "" {
				item := types.AttackItemDrop{
					ItemURL:  common.OriginSMMO + strings.Replace(itemRegex[1][:len(itemRegex[1])-1], `img src="`, ``, -1),
					ItemName: strings.Replace(itemRegex[0], "</span>", "", -1),
				}
				summary.ItemList = append(summary.ItemList, item)
			}
			summary.TotalItem += 1
		}

		// Define time delay
		if travelResp.NextWait == 0 {
			timeDelay = DefaultTimeTravel
			continue
		}
		timeDelay = int(travelResp.NextWait) + 1
	}
	result.Journey = allResponse
	result.Summary = summary
	return result, nil
}
