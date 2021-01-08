package postusers

import (
	"encoding/json"
	"io"

	"github.com/MSex/lab-go-http/app/server/util"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

func parseValidatePath(p httprouter.Params, req *ParsedRequest) validation.Errors {
	err := validation.Errors{
		"Path Parameters": validation.Validate(p, validation.Length(0, 0)),
	}

	return err
}

func parseValidateQuery(rawQuery string, req *ParsedRequest) validation.Errors {
	err := validation.Errors{
		"Query String": validation.Validate(rawQuery, util.NotAllowed),
	}

	return err
}

func parseValidateBody(body io.ReadCloser, parsed *ParsedRequest) validation.Errors {
	parsedBody := RequestBody{}

	if err := json.NewDecoder(body).Decode(&parsedBody); err != nil {
		return validation.Errors{
			"Body": errors.Wrap(err, "Malformed"),
		}
	}

	parsed.User = parsedBody.User

	err := validation.Errors{
		"User": validation.Validate(parsedBody.User, validation.Required),
	}
	if err["User"] != nil {
		return err
	}

	err = validation.Errors{
		"User.Name":  validation.Validate(parsedBody.User.Name, validation.Required, validation.Length(5, 20)),
		"User.Login": validation.Validate(parsedBody.User.Login, validation.Required, validation.Length(1, 5), is.Email),
		"User.Birth": validation.Validate(parsedBody.User.Birth, validation.Required, validation.Date("2006-01-02")),
	}

	return validation.Errors(err)
}
