package postusers

import (
	"net/http"

	"github.com/MSex/lab-go-http/app/data"
	"github.com/pkg/errors"
)



func validatePrecondition(handler *Handler, user *data.User) ( error, int) {
	exists, err := handler.users.ExistsLogin(user.Login)
	if err != nil {
		err := errors.Wrap(err, "Error checking login existence")
		return err, http.StatusInternalServerError
	}
	if exists {
		return errors.Errorf("Login already in use"), http.StatusConflict
	}

	return nil, 0
}