package attack

import (
	"github.com/kben19/smmob/pkg/common/types"
	"github.com/kben19/smmob/pkg/resource/api"
)

type DomainAttack struct {
	rscAPI api.ResourceAPIItf
}

type DomainAttackItf interface {
	DoAttack(payload types.AttackPostPayload, header types.AttackPostHeader, path string) (types.AttackPostResponse, error)
}

func InitDomainAttack(rscAPI api.ResourceAPIItf) DomainAttackItf {
	return &DomainAttack{rscAPI: rscAPI}
}
