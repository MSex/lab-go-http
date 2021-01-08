package getusers

import (
	"io"
	"net/http"

	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/MSex/lab-go-http/app/data"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger

	parseValidatePath  func(pathParams httprouter.Params, parsed *ParsedRequest) validation.Errors
	parseValidateQuery func(rawQuery string, parsed *ParsedRequest) validation.Errors
	parseValidateBody  func(body io.ReadCloser, parse *ParsedRequest) validation.Errors

	users            data.Users
	buildResposeUser func(*data.User) (*User, error)
}

func ProvideHandler(
	users data.Users,
	logger *zap.Logger,
) (*Handler, error) {
	handler := &Handler{
		logger:             logger,
		parseValidatePath:  parseValidatePath,
		parseValidateQuery: parseValidateQuery,
		parseValidateBody:  parseValidateBody,
		users:              users,
		buildResposeUser:   buildResposeUser,
	}

	return handler, nil
}

func (handler *Handler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO ctx := r.Context()

	//TODO auth

	//TODO authorize owner

	//TODO validate context?

	parsed := ParsedRequest{}

	if err := handler.parseValidatePath(p, &parsed).Filter(); err != nil {
		err := errors.Wrap(err, "Invalid request path")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			err := errors.Wrap(err, "Error writing response")
			handler.logger.Error(err.Error())
		}
		return
	}

	if err := handler.parseValidateQuery(r.URL.RawQuery, &parsed).Filter(); err != nil {
		err := errors.Wrap(err, "Invalid request query")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			err := errors.Wrap(err, "Error writing response")
			handler.logger.Error(err.Error())
		}
		return
	}

	if err := handler.parseValidateBody(r.Body, &parsed).Filter(); err != nil {
		err := errors.Wrap(err, "Invalid request body")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			err := errors.Wrap(err, "Error writing response")
			handler.logger.Error(err.Error())
		}
		return
	}

	userId, err := data.UserIdFromString(parsed.UserId)
	if err != nil {
		err := errors.Wrap(err, "Error parsing id")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			err := errors.Wrap(err, "Error writing response")
			handler.logger.Error(err.Error())
		}
		return
	}

	user, err := handler.users.Get(userId)
	if err == data.NotFoundError {
		err := errors.Wrap(err, "User not found")
		handler.logger.Info(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		err := errors.Wrap(err, "Error getting user")
		handler.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msg, err := handler.buildResposeUser(user)
	if err != nil {
		err := errors.Wrap(err, "Error building response")
		handler.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	asJSON, err := json.Marshal(msg)
	if err != nil {
		err := errors.Wrap(err, "Error marshaling response")
		handler.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(asJSON)
	if err != nil {
		err := errors.Wrap(err, "Error writing response")
		handler.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		err := errors.Wrap(err, "Error writing response")
		handler.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
