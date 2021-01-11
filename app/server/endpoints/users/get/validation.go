package getusers

import (
	"io"

	"github.com/MSex/lab-go-http/app/server/util"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

func parseValidatePath(p httprouter.Params, req *ParsedRequest) validation.Errors {
	err := validation.Errors{
		"Path Parameters": validation.Validate(p, validation.Length(1, 1)),
		"UserId": validation.Validate(p.ByName("userId"), validation.Required, is.Digit, validation.Length(1, 20)),
	}

	req.UserId = p.ByName("userId")
	return err
}

func parseValidateQuery(rawQuery string, req *ParsedRequest) validation.Errors {
	err := validation.Errors{
		"Query String": validation.Validate(rawQuery, util.NotAllowed),
	}

	return err
}

func parseValidateBody(body io.ReadCloser, _ *ParsedRequest) validation.Errors {
	buf := make([]byte, 256)
	n, err := body.Read(buf)
	if n > 0 || err != io.EOF {
		return validation.Errors{
			"Body": errors.Errorf("should be empty"),
		}
	}

	return validation.Errors{}
}
