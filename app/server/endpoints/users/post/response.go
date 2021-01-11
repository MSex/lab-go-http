package postusers

import (
	"encoding/json"
	"net/http"

	"github.com/MSex/lab-go-http/app/logger"
	"github.com/pkg/errors"
)

type ResponseBody struct {
	Id string
}

type ResponseHelper interface {
	WriteError(code int, err error, httpWriter http.ResponseWriter, logger logger.Logger)
	WriteBody(body ResponseBody, httpWriter http.ResponseWriter, logger logger.Logger)
}


type ResponseHelperImpl struct {
}

func (helper *ResponseHelperImpl) WriteError(code int, err error, httpWriter http.ResponseWriter, logger logger.Logger) {
	httpWriter.WriteHeader(code)
	helper.writeData([]byte(err.Error()), httpWriter, logger)
}

func (helper *ResponseHelperImpl) WriteBody(body ResponseBody, httpWriter http.ResponseWriter, logger logger.Logger) {
	asJSON, err := json.Marshal(body)
	if err != nil {
		err := errors.Wrap(err, "Error marshaling response")
		logger.Error(err.Error())
		helper.WriteError(http.StatusInternalServerError, err, httpWriter, logger)
		return
	}

	httpWriter.WriteHeader(http.StatusOK)
	helper.writeData(asJSON, httpWriter, logger)
	helper.writeData([]byte("\n"), httpWriter, logger)
}

func (h *ResponseHelperImpl) writeData(data []byte , httpWriter http.ResponseWriter, logger logger.Logger) {
	_, err := httpWriter.Write(data)
	if err != nil {
		err := errors.Wrap(err, "Error writing response")
		logger.Error(err.Error())
	}
}
