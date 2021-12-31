package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/pulkit-tanwar/omh-users-management/lib/constant"
	"github.com/pulkit-tanwar/omh-users-management/lib/database"
	"github.com/pulkit-tanwar/omh-users-management/lib/model"
	log "github.com/sirupsen/logrus"
)

// Ping - This function will ping the echo server
func (s *Server) Ping(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"Ping": "OK"})
}

// CreateUser - This function will create a new user
func (s *Server) CreateUser(c echo.Context) error {
	contentType := c.Request().Header.Get("Content-Type")
	if contentType != "application/json" {
		log.WithFields(log.Fields{
			"ErrorDescription": "Content-Type Not Supported ",
		}).Error("Content-Type not supported for RegisterUser Request")
		return c.JSON(http.StatusBadRequest, model.NewErrorStructure(constant.ContentTypeNotSupported, constant.ContentTypeNotSupportedMessage))
	}

	userItem, err := validateRequestPayload(c)
	if err != nil {
		return err
	}

	timeNow := time.Now().Format(time.RFC3339)
	userItem.DateCreated = timeNow
	userItem.DateModified = timeNow

	err = database.DB.CreateUser(userItem)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":            err,
			"ErrorCode":        constant.FailedToCreateUser,
			"ErrorDescription": constant.FailedToCreateUserMessage,
		}).Error("Error while creating new user")

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return c.JSON(http.StatusConflict, model.NewErrorStructure(constant.UserNameAlreadyExists, constant.UserNameAlreadyExistsMessage))
		}

		return c.JSON(http.StatusInternalServerError, model.NewErrorStructure(constant.FailedToCreateUser, constant.FailedToCreateUserMessage))
	}
	return c.JSON(http.StatusCreated, userItem)
}

func validateRequestPayload(c echo.Context) (model.User, error) {

	user := model.User{}
	defer c.Request().Body.Close()

	receivedJSON, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":            err,
			"ErrorCode":        constant.FailedToReadRequestBody,
			"ErrorDescription": constant.FailedToReadRequestBodyMessage,
		}).Error("Error while reading payload body")

		c.JSON(http.StatusInternalServerError, model.NewErrorStructure(constant.FailedToReadRequestBody, constant.FailedToReadRequestBodyMessage))
		return user, errors.New(constant.FailedToReadRequestBodyMessage)
	}

	err = json.Unmarshal([]byte(receivedJSON), &user)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":            err,
			"ErrorCode":        constant.FailedToUnmarshalRequestBody,
			"ErrorDescription": constant.FailedToUnmarshalRequestBodyMessage,
		}).Error("Error while unmarshalling payload body")

		c.JSON(http.StatusInternalServerError, model.NewErrorStructure(constant.FailedToUnmarshalRequestBody, constant.FailedToUnmarshalRequestBodyMessage))
		return user, errors.New("Error while unmarshalling payload body")
	}

	if user.User_Name == "" {
		log.WithFields(log.Fields{
			"ErrorCode":        constant.UserNameMissing,
			"ErrorDescription": constant.UserNameMissingMessage,
		}).Error("user name missing from payload body")
		c.JSON(http.StatusBadRequest, model.NewErrorStructure(constant.UserNameMissing, constant.UserNameMissingMessage))
		return user, errors.New("user name missing")
	}

	if user.First_Name == "" {
		log.WithFields(log.Fields{
			"ErrorCode":        constant.FirstNameMissing,
			"ErrorDescription": constant.FirstNameMissingMessage,
		}).Error("first name missing from payload body")
		c.JSON(http.StatusBadRequest, model.NewErrorStructure(constant.FirstNameMissing, constant.FirstNameMissingMessage))
		return user, errors.New("first name missing")
	}

	if user.Last_Name == "" {
		log.WithFields(log.Fields{
			"ErrorCode":        constant.LastNameMissing,
			"ErrorDescription": constant.LastNameMissingMessage,
		}).Error("last name missing from payload body")
		c.JSON(http.StatusBadRequest, model.NewErrorStructure(constant.LastNameMissing, constant.LastNameMissingMessage))
		return user, errors.New("last name missing")
	}

	return user, nil
}
