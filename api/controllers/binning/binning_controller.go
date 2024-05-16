package binning

import "net/http"

type BinningController interface {
	FindAll(writer http.ResponseWriter, request *http.Request)
	GetAll(writer http.ResponseWriter, request *http.Request)
}
