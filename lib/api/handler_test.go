package api_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pulkit-tanwar/omh-users-management/lib/config"
	"github.com/pulkit-tanwar/omh-users-management/lib/constant"
	"github.com/pulkit-tanwar/omh-users-management/lib/database"
	"github.com/pulkit-tanwar/omh-users-management/lib/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var createEndpoint = "/api/v1/users"
var getUserEndpoint = "/api/v1/users/pulkit123"

type MockDatabaseClient struct {
	mock.Mock
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

//DBConnect - Mocked method
func (m *MockDatabaseClient) DBConnect() error {
	args := m.Mock.Called()
	return args.Error(0)
}

func (m *MockDatabaseClient) CreateUser(user model.User) error {
	resp := m.Mock.Called(nil)
	return resp.Error(0)
}

func (m *MockDatabaseClient) RetrieveUser(userName string) (model.User, error) {
	resp := m.Mock.Called(userName)
	user := resp.Get(0).(model.User)
	return user, resp.Error(1)
}

func (m *MockDatabaseClient) ModifyUserDetails(user model.User) (model.User, error) {
	resp := m.Mock.Called(nil)
	userDetails := resp.Get(0).(model.User)
	return userDetails, resp.Error(1)
}

func (m *MockDatabaseClient) DeleteUser(userName string) error {
	resp := m.Mock.Called(userName)
	return resp.Error(1)
}

func (m *MockDatabaseClient) GetAllUsers() ([]model.User, error) {
	resp := m.Mock.Called(nil)
	userDetails := resp.Get(0).([]model.User)
	return userDetails, resp.Error(1)
}

func TestPing(t *testing.T) {
	rr := serve(t, get("/api/v1/ping"), config.DefaultConfig())
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", strings.ToLower(rr.HeaderMap["Content-Type"][0]))

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))

	assert.Equal(t, map[string]interface{}{"Ping": "OK"}, response)
}

func TestCreateUserSuccessfull(t *testing.T) {
	data := map[string]interface{}{"username": "22112", "firstname": "pulkit", "lastname": "tanwar"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		assert.Fail(t, "json.Marshal Failed.", err)
	}
	mockDb := &MockDatabaseClient{}
	mockDb.On("CreateUser", nil).Return(nil)
	database.DB = mockDb
	req := post(createEndpoint, string(jsonData))
	req.Header.Add("Content-Type", "application/json")
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusCreated)
	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))

	userName := response["userName"].(string)
	assert.Equal(t, userName, "22112")

	firstName := response["firstName"].(string)
	assert.Equal(t, firstName, "pulkit")

	lastName := response["lastName"].(string)
	assert.Equal(t, lastName, "tanwar")
}

func TestCreateUserDBFailed(t *testing.T) {
	data := map[string]interface{}{"username": "22112", "firstname": "pulkit", "lastname": "tanwar"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		assert.Fail(t, "json.Marshal Failed.", err)
	}
	mockDb := &MockDatabaseClient{}
	mockDb.On("CreateUser", nil).Return(errors.New("error"))
	database.DB = mockDb
	req := post(createEndpoint, string(jsonData))
	req.Header.Add("Content-Type", "application/json")
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusInternalServerError)

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	errorCode := response["errorCode"].(float64)
	assert.Equal(t, errorCode, float64(constant.FailedToCreateUser))

	errorDescription := response["errorDescription"].(string)
	assert.Equal(t, errorDescription, constant.FailedToCreateUserMessage)
}

func TestCreateUserInvalidBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, createEndpoint, errReader(0))
	req.Header.Add("Content-Type", "application/json")
	rr := serve(t, req, config.DefaultConfig())

	assert.Equal(t, rr.Code, http.StatusInternalServerError)
	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	errorCode := response["errorCode"].(float64)
	assert.Equal(t, errorCode, float64(constant.FailedToReadRequestBody))

	errorDescription := response["errorDescription"].(string)
	assert.Equal(t, errorDescription, constant.FailedToReadRequestBodyMessage)
}

func TestCreateUserInvalidContentType(t *testing.T) {
	data := map[string]interface{}{"username": "22112", "firstname": "pulkit", "lastname": "tanwar"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		assert.Fail(t, "json.Marshal Failed.", err)
	}
	mockDb := &MockDatabaseClient{}
	mockDb.On("CreateUser", nil).Return(errors.New("error"))
	database.DB = mockDb
	req := post(createEndpoint, string(jsonData))
	req.Header.Add("Content-Type", "text/csv")
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusBadRequest)

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	errorCode := response["errorCode"].(float64)
	assert.Equal(t, errorCode, float64(constant.ContentTypeNotSupported))

	errorDescription := response["errorDescription"].(string)
	assert.Equal(t, errorDescription, constant.ContentTypeNotSupportedMessage)
}

func TestCreateUserMissingUserName(t *testing.T) {
	data := map[string]interface{}{"firstname": "pulkit", "lastname": "tanwar"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		assert.Fail(t, "json.Marshal Failed.", err)
	}
	mockDb := &MockDatabaseClient{}
	mockDb.On("CreateUser", nil).Return(errors.New("error"))
	database.DB = mockDb
	req := post(createEndpoint, string(jsonData))
	req.Header.Add("Content-Type", "application/json")
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusBadRequest)

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	errorCode := response["errorCode"].(float64)
	assert.Equal(t, errorCode, float64(constant.UserNameMissing))

	errorDescription := response["errorDescription"].(string)
	assert.Equal(t, errorDescription, constant.UserNameMissingMessage)
}

func TestCreateUserMissingFirstName(t *testing.T) {
	data := map[string]interface{}{"username": "22112", "lastname": "tanwar"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		assert.Fail(t, "json.Marshal Failed.", err)
	}
	mockDb := &MockDatabaseClient{}
	mockDb.On("CreateUser", nil).Return(errors.New("error"))
	database.DB = mockDb
	req := post(createEndpoint, string(jsonData))
	req.Header.Add("Content-Type", "application/json")
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusBadRequest)

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	errorCode := response["errorCode"].(float64)
	assert.Equal(t, errorCode, float64(constant.FirstNameMissing))

	errorDescription := response["errorDescription"].(string)
	assert.Equal(t, errorDescription, constant.FirstNameMissingMessage)
}

func TestCreateUserMissingLastName(t *testing.T) {
	data := map[string]interface{}{"username": "22112", "firstname": "pulkit"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		assert.Fail(t, "json.Marshal Failed.", err)
	}
	mockDb := &MockDatabaseClient{}
	mockDb.On("CreateUser", nil).Return(errors.New("error"))
	database.DB = mockDb
	req := post(createEndpoint, string(jsonData))
	req.Header.Add("Content-Type", "application/json")
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusBadRequest)

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	errorCode := response["errorCode"].(float64)
	assert.Equal(t, errorCode, float64(constant.LastNameMissing))

	errorDescription := response["errorDescription"].(string)
	assert.Equal(t, errorDescription, constant.LastNameMissingMessage)
}

func TestCreateUserErrorWhileUnmarshling(t *testing.T) {
	data := "{"
	jsonData, err := json.Marshal(data)
	if err != nil {
		assert.Fail(t, "json.Marshal Failed.", err)
	}
	mockDb := &MockDatabaseClient{}
	mockDb.On("CreateUser", nil).Return(errors.New("error"))
	database.DB = mockDb
	req := post(createEndpoint, string(jsonData))
	req.Header.Add("Content-Type", "application/json")
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusInternalServerError)

	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))
	errorCode := response["errorCode"].(float64)
	assert.Equal(t, errorCode, float64(constant.FailedToUnmarshalRequestBody))

	errorDescription := response["errorDescription"].(string)
	assert.Equal(t, errorDescription, constant.FailedToUnmarshalRequestBodyMessage)
}

func TestGetUserByUserNameSuccssfull(t *testing.T) {
	userr := model.User{User_Name: "pulkit123", First_Name: "pulkit", Last_Name: "tanwar"}

	mockDb := &MockDatabaseClient{}
	mockDb.On("RetrieveUser", "pulkit123").Return(userr, nil)
	database.DB = mockDb
	req := get(getUserEndpoint)
	rr := serve(t, req, config.DefaultConfig())
	assert.Equal(t, rr.Code, http.StatusOK)
	var response map[string]interface{}
	require.Nil(t, json.NewDecoder(rr.Body).Decode(&response))

	userName := response["userName"].(string)
	assert.Equal(t, userName, "pulkit123")

	firstName := response["firstName"].(string)
	assert.Equal(t, firstName, "pulkit")

	lastName := response["lastName"].(string)
	assert.Equal(t, lastName, "tanwar")
}
