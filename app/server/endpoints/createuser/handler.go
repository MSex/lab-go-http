package createuser

import (
	"io"
	"net/http"

	"encoding/json"

	"github.com/MSex/lab-go-http/app/data"
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
	users  data.Users
	logger *zap.Logger
}

func ProvideHandler(
	users data.Users,
	logger *zap.Logger,
) (*Handler, error) {
	handler := &Handler{
		users:  users,
		logger: logger,
	}

	return handler, nil
}

func (handler *Handler) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO ctx := r.Context()

	//TODO auth

	//TODO authorize owner

	//TODO validate context?

	//TODO validate syntax

	cursor, err := handler.users.LoadCursor()
	if err != nil {
		err := errors.Wrap(err, "Error getting loader")
		handler.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cursor.Close()

	for {
		user, err := cursor.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			err := errors.Wrap(err, "Error getting current cursor value")
			handler.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//TODO msg := toMessage(user)

		asJSON, err := json.Marshal(user)
		if err != nil {
			err := errors.Wrap(err, "Error getting current cursor value")
			handler.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(asJSON)       //TODO error check
		w.Write([]byte("\n")) //TODO error check
		// if err := stream.Send(msg); err != nil {
		// 	if status.Code(err) == codes.Canceled {
		// 		return status.Error(codes.Canceled, err.Error())
		// 	}
		// 	err := errors.Wrap(err, "Error sending message")
		// 	handler.logger.Error(err)
		// 	return status.Error(codes.Internal, err.Error())
		// }
		// w.Flush()
	}
}
