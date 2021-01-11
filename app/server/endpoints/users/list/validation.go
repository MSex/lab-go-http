package listusers

import (
	"io"
	"net/url"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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
	queryValues, perr := url.ParseQuery(rawQuery)
	if perr != nil {
		return validation.Errors{
			"Query String": errors.Wrap(perr, "Malformed"),
		}
	}

	err := validation.Errors{
		//TODO validate not extra parms
		"Query String": validation.Validate(queryValues, validation.Length(0,1)),
		"StartDate": validation.Validate(
			queryValues["startdate"],
			validation.Length(1, 1),
			validation.Each(validation.Date("2006-01-02")),
		),
	}

	req.Date = queryValues.Get("startdate")
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
