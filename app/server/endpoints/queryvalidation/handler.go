package createuser

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Handler struct {
	// TODO auth                  data.Auth
	// TODO authorizeOwner        func(handler *Handler, userId string, request *hoth_operacoes.CreateOperacaoRequest) error
	// TODO validateRequestSyntax func(req *hoth_operacoes.ListMoedaViewsRequest) validation.Errors
	// TODO validateContext       func(ctx context.Context) error
	// TODO validatePrecondition  func(handler *Handler, operacao *data.Operacao) error
	validateRouteParams func(ps httprouter.Params) validation.Errors
	validateQueryParams func(url.Values) validation.Errors
	logger              *zap.Logger
}

func ProvideHandler(
	logger *zap.Logger,
) (*Handler, error) {
	handler := &Handler{
		logger:              logger,
		validateRouteParams: validateRouteParams,
		validateQueryParams: validateQueryParams,
	}

	return handler, nil
}

type User struct {
	Name string
	Age  int
}

type RequestBody struct {
	User User
}

func (handler *Handler) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO ctx := r.Context()

	//TODO auth

	//TODO authorize owner

	//TODO validate context?

	if err := handler.validateRouteParams(ps).Filter(); err != nil {
		err := errors.Wrap(err, "Invalid parameters")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) //TODO error check
		return
	}

	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		err := errors.Wrap(err, "Malformed URI query")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) //TODO error check
		return
	}

	if err := handler.validateQueryParams(queryValues).Filter(); err != nil {
		err := errors.Wrap(err, "Invalid parameters")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) //TODO error check
		return
	}

	var body RequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		err := errors.Wrap(err, "Invalid body")
		handler.logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) //TODO error check
		return
	}

	// if err := handler.validateRequestSyntax(request).Filter(); err != nil {
	// 	handler.logger.Warn(err)
	// 	return nil, status.Error(codes.InvalidArgument, err.Error())
	// }

	//TODO validate syntax

	// cursor, err := handler.users.LoadCursor()
	// if err != nil {
	// 	err := errors.Wrap(err, "Error getting loader")
	// 	handler.logger.Error(err.Error())
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// defer cursor.Close()

	// for {
	// 	user, err := cursor.Next()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		err := errors.Wrap(err, "Error getting current cursor value")
	// 		handler.logger.Error(err.Error())
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}

	// 	//TODO msg := toMessage(user)

	// 	asJSON, err := json.Marshal(user)
	// 	if err != nil {
	// 		err := errors.Wrap(err, "Error getting current cursor value")
	// 		handler.logger.Error(err.Error())
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}

	// 	w.Write(asJSON)       //TODO error check
	// 	w.Write([]byte("\n")) //TODO error check
	// if err := stream.Send(msg); err != nil {
	// 	if status.Code(err) == codes.Canceled {
	// 		return status.Error(codes.Canceled, err.Error())
	// 	}
	// 	err := errors.Wrap(err, "Error sending message")
	// 	handler.logger.Error(err)
	// 	return status.Error(codes.Internal, err.Error())
	// }
	// w.Flush()
	// }
}

func validateRouteParams(ps httprouter.Params) validation.Errors {
	return validation.Errors{
		"UserId": validation.Validate(ps.ByName("id"), validation.Required, is.Digit, validation.Length(5, 20)),
	}
}

func validateQueryParams(q url.Values) validation.Errors {
	return validation.Errors{
		"StartDate": validation.Validate(
			q.Get("startDate"),
			validation.Required,
			validation.Match(regexp.MustCompile("^/d{4}-/d{2}-/d{2}$")),
		),
	}
}

func validateRequestBody(b RequestBody) validation.Errors {
	err := map[string]error{}
	err["User"] = validation.Validate(b.User, validation.Required)

	if err["User"] != nil {
		return err
	}

	err["User.Name"] = validation.Validate(b.User.Name, validation.Required)
	err["User.Age"] = validation.Validate(b.User.Age, validation.Required)

	return validation.Errors(err)
}
