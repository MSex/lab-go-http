package postusers

import (
	"net/http"
	"net/url"
	"testing"

	loggermocks "github.com/MSex/lab-go-http/app/logger/mocks"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)


func TestHandler(t *testing.T) {
	suite.Run(t, new(handleTests))
}

type handleTests struct {
	suite.Suite
}

func (suite *handleTests) SetupTest() {
}

func (suite *handleTests) TestParseRequestFail() {
	request := &http.Request{
		URL: &url.URL{},
	}
	err := errors.Errorf("Mocked Error")

	logger := new(loggermocks.Logger)
	logger.On("Warn", mock.Anything)

	requestHelper := new(RequestHelperMock)
	requestHelper.On("ParseRequest", httprouter.Params(nil), "", nil).Return(nil, err)

	responseHelper := new(ResponseHelperMock)
	responseHelper.On("WriteError", 400, mock.Anything, mock.Anything, mock.Anything)

	handler := &Handler{
		logger: logger,
		requestHelper: requestHelper,
		responseHelper: responseHelper,
	}


	handler.Handle(nil, request, nil)

	logger.AssertExpectations(suite.T())
	requestHelper.AssertExpectations(suite.T())
	responseHelper.AssertExpectations(suite.T())
}

func (suite *handleTests) TestValidateRequestFail() {
	request := &http.Request{
		URL: &url.URL{},
	}
	err := errors.Errorf("Mocked Error")
	
	logger := new(loggermocks.Logger)
	logger.On("Warn", mock.Anything)

	requestHelper := new(RequestHelperMock)
	requestHelper.On("ParseRequest", httprouter.Params(nil), "", nil).Return(nil, nil)
	requestHelper.On("ValidateRequest", mock.Anything).Return(validation.Errors{"Any": err})

	responseHelper := new(ResponseHelperMock)
	responseHelper.On("WriteError", 400, mock.Anything, mock.Anything, mock.Anything)

	handler := &Handler{
		logger: logger,
		requestHelper: requestHelper,
		responseHelper: responseHelper,
	}


	handler.Handle(nil, request, nil)

	logger.AssertExpectations(suite.T())
	requestHelper.AssertExpectations(suite.T())
	responseHelper.AssertExpectations(suite.T())
}


// func (suite *handleTests) TestParseRequestFail2() {
// 	logger := new(loggermocks.Logger)
// 	logger.On("Warn", mock.Anything)
// 	logger.On("Error", mock.Anything)

// 	writer := new(utilmocks.ResponseWriter)
// 	writer.On("WriteHeader", 400)
// 	writer.On("Write", mock.Anything).Return(1, errors.Errorf("Error"))
// 	// writer.On("WriteHeader", 500)


// 	handler := &Handler{
// 		logger: logger,
// 		parseRequest: func(pathParams httprouter.Params, query string, body io.ReadCloser) (*RequestData, error) {
// 			return nil, errors.Errorf("Error")
// 		},
// 	}
// 	request := &http.Request{
// 		URL: &url.URL{},
// 	}

// 	handler.Handle(writer, request, nil)
// 	 logger.AssertExpectations(suite.T())
// 	 writer.AssertExpectations(suite.T())


// }
	// got, err := tt.target.Handle(tt.args.ctx, tt.args.request)
	// 		errCode := status.Code(err)
	// 		if tt.wantErr != (err != nil) {
	// 			t.Errorf("handler.Handle() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if tt.wantErr && errCode != tt.wantErrCode {
	// 			t.Errorf("handler.Handle() err = %v errorCode = %v, wantErrCode %v", err, errCode, tt.wantErrCode)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("handler.Handle() = %v, want %v", got, tt.want)
	// 		}
		
		
	// 	}


// {
// 			name: "Auth fails",
// 			target: &Handler{
// 				auth: &mock.Auth{
// 					OnAuthenticateAndAuthorize: func(ctx context.Context, request interface{}) (info data.UserInfo, e error) {
// 						return nil, status.Error(codes.Unauthenticated, "Auth Error")
// 					},
// 				},
// 				authorizeOwner: func(handler *Handler, userId string, request *hoth_operacoes.CreateOperacaoRequest) error {
// 					return nil
// 				},
// 				validateRequestSyntax: func(req *hoth_operacoes.CreateOperacaoRequest) validation.Errors { return validation.Errors{} },
// 				validateContext:       func(ctx context.Context) error { return nil },
// 				validatePrecondition:  func(*Handler, *data.Operacao) error { return nil },
// 				ativos:                &mock.Ativos{},
// 				boletas:               &mock.Boletas{},
// 				carteiras:             &mock.CarteiraViews{},
// 				operacoes:             &mock.Operacoes{},
// 				logger:                &mock.Logger{},
// 			},
// 			args: args{
// 				ctx:     &mock.Context{},
// 				request: &hoth_operacoes.CreateOperacaoRequest{},
// 			},
// 			wantErr:     true,
// 			wantErrCode: codes.Unauthenticated,
// 		},