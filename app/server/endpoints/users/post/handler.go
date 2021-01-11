package postusers

import (
	"net/http"

	"github.com/MSex/lab-go-http/app/data"
	"github.com/MSex/lab-go-http/app/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Handler struct {
	logger logger.Logger

	requestHelper RequestHelper
	responseHelper ResponseHelper

	validatePrecondition func(handler *Handler, user *data.User) (error, int)
	users                data.Users
}

func ProvideHandler(
	users data.Users,
	logger *zap.Logger,
) (*Handler, error) {
	handler := &Handler{
		logger: logger,

		requestHelper: &RequestHelperImpl{},
		responseHelper: &ResponseHelperImpl{},

		users:                users,
		validatePrecondition: validatePrecondition,
	}

	return handler, nil
}

func (handler *Handler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//TODO auth

	requestData, err := handler.requestHelper.ParseRequest(p, r.URL.RawQuery, r.Body)
	if err != nil {
		err := errors.Wrap(err, "Error parsing request")
		handler.logger.Warn(err.Error())
		handler.responseHelper.WriteError(http.StatusBadRequest, err, w, handler.logger)
		return
	}

	if err := handler.requestHelper.ValidateRequest(requestData).Filter(); err != nil {
		err := errors.Wrap(err, "Invalid request")
		handler.logger.Warn(err.Error())
		handler.responseHelper.WriteError(http.StatusBadRequest, err, w, handler.logger)
		return
	}

	user := &data.User{
		Name:  requestData.Body.User.Name, 
		Login: requestData.Body.User.Login,
		Birth: requestData.Body.User.Birth,
	}

	err, code := handler.validatePrecondition(handler, user)
	if err != nil {
		handler.logger.Error(err.Error())
		handler.responseHelper.WriteError(code, err, w, handler.logger)
		return
	}

	id, err := handler.users.Create(user)
	if err != nil {
		err = errors.Wrap(err, "Error presisting data")
		handler.logger.Error(err.Error())
		handler.responseHelper.WriteError(http.StatusInternalServerError, err, w, handler.logger)
		return
	}

	msg := ResponseBody{
		Id: id.String(),
	}

	handler.responseHelper.WriteBody(msg,  w, handler.logger)



}

