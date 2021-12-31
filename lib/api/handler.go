package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pulkit-tanwar/omh-users-management/lib/constant"
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
		return c.JSON(http.StatusBadRequest, model.NewErrorStructure(100, "Content-Type Not Supported"))
	}

	reqPayload, err := validateRequestPayload(c)
	if err != nil {
		return err
	}

	// TODO : Add Db Client and persist data

	return c.JSON(http.StatusOK, reqPayload)
}

func validateRequestPayload(c echo.Context) (map[string]interface{}, error) {

	defer c.Request().Body.Close()
	var unmarshalledJSON map[string]interface{}

	receivedJSON, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":            err,
			"ErrorCode":        constant.FailedToReadRequestBody,
			"ErrorDescription": constant.FailedToReadRequestBodyMessage,
		}).Error("Error while reading payload body")

		c.JSON(http.StatusInternalServerError, model.NewErrorStructure(constant.FailedToReadRequestBody, constant.FailedToReadRequestBodyMessage))
		return nil, errors.New(constant.FailedToReadRequestBodyMessage)
	}

	err = json.Unmarshal([]byte(receivedJSON), &unmarshalledJSON)
	if err != nil {

		log.WithFields(log.Fields{
			"Error":            err,
			"ErrorCode":        constant.FailedToReadRequestBody,
			"ErrorDescription": constant.FailedToReadRequestBodyMessage,
		}).Error("Error while unmarshalling payload body")

		c.JSON(http.StatusInternalServerError, model.NewErrorStructure(constant.FailedToReadRequestBody, constant.FailedToReadRequestBodyMessage))
		return nil, errors.New(constant.FailedToReadRequestBodyMessage)
	}

	return unmarshalledJSON, nil
}
