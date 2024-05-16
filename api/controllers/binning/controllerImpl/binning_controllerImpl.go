package controllerImpl

import (
	"encoding/json"
	"net/http"
	binning2 "user-services/api/controllers/binning"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	"user-services/api/payloads"
	request2 "user-services/api/payloads/request"
	errorresponses "user-services/api/repositories/error"
	"user-services/api/services/binning"
)

type BinningControllerImpl struct {
	BinningService binning.BinningService
}

func NewBinningControllerImpl(services binning.BinningService) binning2.BinningController {
	return &BinningControllerImpl{BinningService: services}
}
func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}
}
func MakeErrorResponse(errMsg string, errorResponses *errorresponses.ErrorRespones) {
	errorResponses.Success = false
	errorResponses.LogSysNo = 0
	errorResponses.Message = errMsg
}

func (b *BinningControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	writer.Header().Add("Content-Type", "application/json")

	//headerString, err2 := json.Marshal(request.Header)
	var requestbody []request2.BinningHeaderRequest
	encoder := json.NewEncoder(writer)
	ReadFromRequestBody(request, &requestbody)
	binningresponses, err, errmsg := b.BinningService.FindAll(request.Context(), requestbody)
	if err != nil {
		var errr errorresponses.ErrorRespones
		MakeErrorResponse(errmsg, &errr)
	}
	err = encoder.Encode(binningresponses)
	if err != nil {
		panic(err)
	}
	panic("implement me")
}

// GetAll Get All Bining List Via Header
//
// @Security BearerAuth
//
//	@Summary		Show An Binning List
//	@Description	Get Binning List By Header
//	@Tags			Binning
//	@Accept			json
//	@Produce		json
//	@Param			request	body		[]request2.BinningHeaderRequest	true	"Insert Header Request"
//	@Success		200		{object}	[]response.BinningHeaderResponses
//	@Router			/api/binning/getAll [post]
func (b *BinningControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request) {
	//query := request.URL.Query()
	//limit, _ := strconv.Atoi(query.Get("limit"))
	//page, _ := strconv.Atoi(query.Get("page"))
	//tokenn := request.Header.Get("Authorization")
	//panic(tokenn)
	var ParameterRequest []request2.BinningHeaderRequest
	err := jsonchecker.ReadFromRequestBody(request, &ParameterRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}
	get, errFetch := b.BinningService.GetAll(request.Context(), ParameterRequest)

	if errFetch != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, get, "Get Approval Success", http.StatusOK)
}
