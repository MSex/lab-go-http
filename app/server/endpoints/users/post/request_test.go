package postusers

import (
	"io"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/suite"
)

func TestRequest(t *testing.T) {
	suite.Run(t, new(parseRequestTests))
			suite.Run(t, new(validateRequestTests))
}

type parseRequestTests struct {
	suite.Suite

	target RequestHelper
	validBody io.ReadCloser
}

func (suite *parseRequestTests) SetupTest() {
	suite.validBody = ioutil.NopCloser(strings.NewReader("{}"))
	suite.target = &RequestHelperImpl{}
}

func (suite *parseRequestTests) TestPathParams() {
	params := httprouter.Params{}

	request, err := suite.target.ParseRequest(params, "", suite.validBody)

	suite.NoError(err)
	suite.Equal(&params, request.Path)
}

func (suite *parseRequestTests) TestQueryString() {
	query := "a=b"

	request, err := suite.target.ParseRequest(nil, query, suite.validBody)

	suite.NoError(err)
	suite.Equal(request.Query.Get("a"), "b")
}

func (suite *parseRequestTests) TestErrorParsingQueryString() {
	query := "a=b%TT"

	_, err := suite.target.ParseRequest(nil, query, nil)

	suite.Error(err)
}

func (suite *parseRequestTests) TestNilBody() {
	_, err := suite.target.ParseRequest(nil, "", nil)

	suite.Error(err)
}

func (suite *parseRequestTests) TestEmptyBody() {
	body := ioutil.NopCloser(strings.NewReader(""))

	_, err := suite.target.ParseRequest(nil, "", body)

	suite.Error(err)
}

func (suite *parseRequestTests) TestNonJsonBody() {
	body := ioutil.NopCloser(strings.NewReader("bla"))

	_, err := suite.target.ParseRequest(nil, "", body)

	suite.Error(err)
}

func (suite *parseRequestTests) TestMalformedBody() {
	body := ioutil.NopCloser(strings.NewReader(`{"user": "not a user"}`))

	_, err := suite.target.ParseRequest(nil, "", body)

	suite.Error(err)
}


type validateRequestTests struct {
	suite.Suite

	target RequestHelper
	validRequest *RequestData
}

func (suite *validateRequestTests) SetupTest() {
	suite.target = &RequestHelperImpl{}
	suite.validRequest = &RequestData{
		Body: &RequestBody{
			User: &RequestUser{
				Name:  "Mauricio",
				Login: "mauricio@exemple.com",
				Birth: "1977-05-24",
			},
		},
	}
}

func (suite *validateRequestTests) TestAcceptNilPathParams() {
	request := suite.validRequest
	request.Path = nil

	err := suite.target.ValidateRequest(request)

	suite.NoError(err.Filter())
}

func (suite *validateRequestTests) TestAcceptEmptyPathParams() {
	request := suite.validRequest
	request.Path = &httprouter.Params{}

	err := suite.target.ValidateRequest(request)

	suite.NoError(err.Filter())
}

func (suite *validateRequestTests) TestRejectAnyPathParameter() {
	request := suite.validRequest
	request.Path = &httprouter.Params{
		{
			Key:   "any",
			Value: "any",
		},
	}

	err := suite.target.ValidateRequest(request)

	suite.Error(err["Path Params"])
}

func (suite *validateRequestTests) TestAcceptNilQuery() {
	request := suite.validRequest
	request.Query = nil

	err := suite.target.ValidateRequest(request)

	suite.NoError(err.Filter())
}

func (suite *validateRequestTests) TestAcceptEmptyQuery() {
	request := suite.validRequest
	request.Query = &url.Values{}

	err := suite.target.ValidateRequest(request)

	suite.NoError(err.Filter())
}

func (suite *validateRequestTests) TestRejectNonEmptyQuery() {
	request := suite.validRequest
	request.Query = &url.Values{
		"a": nil,
	}

	err := suite.target.ValidateRequest(request)

	suite.Error(err["Query Params"])
}

func (suite *validateRequestTests) TestRejectNilBody() {
	request := suite.validRequest
	request.Body = nil

	err := suite.target.ValidateRequest(request)

	suite.Error(err["Body"])
}

func (suite *validateRequestTests) TestRejectNilUser() {
	request := suite.validRequest
	request.Body.User = nil

	err := suite.target.ValidateRequest(request)

	suite.Error(err["User"])
}

func (suite *validateRequestTests) TestValidateName() {
	request := suite.validRequest
	err := suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Name = ""
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Name"])

	request.Body.User.Name = "12345"
	err = suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Name = "1234"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Name"])

	request.Body.User.Name = "12345678901234567890"
	err = suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Name = "123456789012345678901"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Name"])
}

func (suite *validateRequestTests) TestValidateLogin() {
	request := suite.validRequest
	err := suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Login = ""
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Login"])

	request.Body.User.Login = "a@b.com"
	err = suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Login = "a@b.co"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Login"])

	request.Body.User.Login = "a234567890123@bb.com"
	err = suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Login = "a2345678901234@bb.com"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Login"])

	request.Body.User.Login = "abcdefg"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Login"])
}

func (suite *validateRequestTests) TestValidateBirth() {
	request := suite.validRequest
	err := suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Birth = ""
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Birth"])

	request.Body.User.Birth = "blabla"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Birth"])

	request.Body.User.Birth = "2020-30-30"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Birth"])

	request.Body.User.Birth = "2021-02-29"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Birth"])

	request.Body.User.Birth = "2020-02-29"
	err = suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Birth = "2100-02-29"
	err = suite.target.ValidateRequest(request)
	suite.Error(err["User.Birth"])

	request.Body.User.Birth = "2000-02-29"
	err = suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())

	request.Body.User.Birth = "1977-05-24"
	err = suite.target.ValidateRequest(request)
	suite.NoError(err.Filter())
}
