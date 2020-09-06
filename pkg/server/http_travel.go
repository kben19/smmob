package server

import (
	"net/http"
	"strconv"

	"github.com/kben19/smmob/pkg/common"
	"github.com/kben19/smmob/pkg/usecase"
)

const (
	failedToPerformTravel           = "failed to perform travel"
	failedToPerformTravelAttack     = "failed to perform travel attack"
	failedToPerformTravelAttackLoop = "failed to perform travel attack loop"
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
	if err == common.ErrInvalidPayload {
		HTTPError(w, err, failedToPerformTravelAttack, http.StatusBadRequest, token, apiToken, hash)
		return
	}
	if err != nil {
		HTTPError(w, err, failedToPerformTravel, http.StatusInternalServerError, token, apiToken, hash)
		return
	}
	WriteHTTPResponse(w, res)
	return
}

func (h HandlerTravel) HandleTravelAttack(w http.ResponseWriter, r *http.Request) {
	token, apiToken, hash := r.FormValue(common.ParamToken), r.FormValue(common.ParamAPIToken), r.FormValue(common.ParamHash)
	specialAtkStr := r.FormValue(common.ParamSpecialAttack)
	specialAtk, err := strconv.ParseBool(specialAtkStr)
	if err != nil {
		HTTPError(w, err, failedToPerformTravelAttack, http.StatusBadRequest, specialAtkStr)
		return
	}
	res, err := h.usecaseTravel.PerformTravelAndAttack(token, apiToken, hash, specialAtk)
	if err == common.ErrInvalidPayload {
		HTTPError(w, err, failedToPerformTravelAttack, http.StatusBadRequest, token, apiToken, hash)
		return
	}
	if err != nil {
		HTTPError(w, err, failedToPerformTravelAttack, http.StatusInternalServerError, token, apiToken, hash)
		return
	}
	WriteHTTPResponse(w, res)
	return
}

func (h HandlerTravel) HandleTravelAttackLoop(w http.ResponseWriter, r *http.Request) {
	token, apiToken, hash := r.FormValue(common.ParamToken), r.FormValue(common.ParamAPIToken), r.FormValue(common.ParamHash)
	specialAtkStr := r.FormValue(common.ParamSpecialAttack)
	specialAtk, err := strconv.ParseBool(specialAtkStr)
	if err != nil {
		HTTPError(w, err, failedToPerformTravelAttack, http.StatusBadRequest, specialAtkStr)
		return
	}
	loopStr := r.FormValue(common.ParamLoop)
	loop, err := strconv.Atoi(loopStr)
	if err != nil {
		HTTPError(w, err, failedToPerformTravelAttackLoop, http.StatusBadRequest, loopStr)
		return
	}

	res, err := h.usecaseTravel.PerformTravelAndAttackLoop(token, apiToken, hash, specialAtk, loop)
	if err == common.ErrInvalidPayload {
		HTTPError(w, err, failedToPerformTravelAttack, http.StatusBadRequest, token, apiToken, hash)
		return
	}
	if err != nil {
		HTTPError(w, err, failedToPerformTravelAttack, http.StatusInternalServerError, token, apiToken, hash)
		return
	}
	WriteHTTPResponse(w, res)
	return
}
