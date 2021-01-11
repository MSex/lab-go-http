package postusers

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/MSex/lab-go-http/app/server/util"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type RequestData struct {
	Path  *httprouter.Params
	Query *url.Values
	Body  *RequestBody
}

type RequestUser struct {
	Name  string
	Login string
	Birth string
}

type RequestBody struct {
	User *RequestUser
}

type RequestHelper interface {
	ParseRequest(pathParams httprouter.Params, queryString string, body io.ReadCloser) (*RequestData, error)
	ValidateRequest(request *RequestData) validation.Errors
}

type RequestHelperImpl struct {
}

func (helper *RequestHelperImpl) ParseRequest(pathParams httprouter.Params, queryString string, body io.ReadCloser) (*RequestData, error) {
	request := &RequestData{
		Path: &pathParams,
		Body: &RequestBody{},
	}

	qValues, err := url.ParseQuery(queryString)
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing query string")
	}
	request.Query = &qValues

	if body == nil {
		return nil, errors.Errorf("Error request body can't be nil")
	}
	err = json.NewDecoder(body).Decode(request.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing request body")
	}

	return request, nil
}

func (helper *RequestHelperImpl) ValidateRequest(request *RequestData) validation.Errors {
	errors := validation.Errors{}

	errors["Path Params"] = validation.Validate(request.Path, util.NotAllowed)
	errors["Query Params"] = validation.Validate(request.Query, util.NotAllowed)

	errors["Body"] = validation.Validate(request.Body, validation.Required)
	if errors["Body"] != nil {
		return errors
	}

	errors["User"] = validation.Validate(request.Body.User, validation.Required)
	if errors["User"] != nil {
		return errors
	}

	errors["User.Name"] = validation.Validate(request.Body.User.Name, validation.Required, validation.Length(5, 20))
	errors["User.Login"] = validation.Validate(request.Body.User.Login, validation.Required, validation.Length(7, 20), is.EmailFormat)
	errors["User.Birth"] = validation.Validate(request.Body.User.Birth, validation.Required, validation.Date("2006-01-02"))

	return errors
}
