package server

import (
	"net/http"

	"github.com/kben19/smmob/pkg/common"
	"github.com/kben19/smmob/pkg/usecase"
)

const (
	failedToPerformTravel = "failed to perform travel"
)

type HandlerTravel struct {
	usecaseTravel usecase.UsecaseTravelItf
}

func InitHandlerTravel(usecaseTravel usecase.UsecaseTravelItf) HandlerTravel {
	return HandlerTravel{usecaseTravel: usecaseTravel}
}

func (h HandlerTravel) HandlePerformTravel(w http.ResponseWriter, r *http.Request) {
	token, apiToken, hash := r.FormValue(common.ParamToken), r.FormValue(common.ParamAPIToken), r.FormValue(common.ParamHash)
	res, err := h.usecaseTravel.PerformTravel(token, apiToken, hash)
	if err != nil {
		HTTPError(w, err, failedToPerformTravel, http.StatusInternalServerError, token, apiToken, hash)
		return
	}
	WriteHTTPResponse(w, res)
	return
}
