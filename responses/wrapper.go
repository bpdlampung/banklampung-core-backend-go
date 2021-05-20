package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"net/http"
	"time"
)

type modelSuccess struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type modelError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code,omitempty"`
	Success   bool   `json:"success"`
	Message   *struct {
		Indonesian *string `json:"indonesian"`
		English    *string `json:"english"`
	} `json:"message,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

func Success(c *gin.Context, data interface{}) {
	logs.GetAppLogger().InfoInterface(data)

	c.JSON(http.StatusOK, modelSuccess{
		Data:      data,
		Success:   true,
		Code:      http.StatusOK,
		Timestamp: time.Now().Unix(),
	})

	c.Abort()

	return
}

func Error(c *gin.Context, error error) {
	statusCode := getErrorStatusCode(error)
	errorMsg := error.Error()

	logs.GetAppLogger().Error(error.Error())

	c.JSON(statusCode, modelError{
		Success: false,
		Code:    statusCode,
		Message: &struct {
			Indonesian *string `json:"indonesian"`
			English    *string `json:"english"`
		}{Indonesian: &errorMsg, English: &errorMsg},
		Timestamp: time.Now().Unix(),
	})

	c.Abort()

	return
}

func ErrorWithErrorCode(c *gin.Context, error error, errorCode string) {
	statusCode := getErrorStatusCode(error)
	errorMsg := error.Error()

	logs.GetAppLogger().Error(error.Error())

	c.JSON(statusCode, modelError{
		ErrorCode: errorCode,
		Success:   false,
		Code:      statusCode,
		Message: &struct {
			Indonesian *string `json:"indonesian"`
			English    *string `json:"english"`
		}{Indonesian: &errorMsg, English: &errorMsg},
		Timestamp: time.Now().Unix(),
	})

	c.Abort()
	return
}

func ErrorWithMessage(c *gin.Context, statusCode int, messageId, messageEn, errorCode string) {
	logs.GetAppLogger().Error(messageId)

	c.JSON(statusCode, modelError{
		ErrorCode: errorCode,
		Success:   false,
		Code:      statusCode,
		Message: &struct {
			Indonesian *string `json:"indonesian"`
			English    *string `json:"english"`
		}{Indonesian: &messageId, English: &messageEn},
		Timestamp: time.Now().Unix(),
	})

	c.Abort()

	return
}

func ErrorWithIdMessage(c *gin.Context, statusCode int, messageId, errorCode string) {
	logs.GetAppLogger().Error(messageId)

	c.JSON(statusCode, modelError{
		ErrorCode: errorCode,
		Success:   false,
		Code:      statusCode,
		Message: &struct {
			Indonesian *string `json:"indonesian"`
			English    *string `json:"english"`
		}{Indonesian: &messageId, English: nil},
		Timestamp: time.Now().Unix(),
	})

	c.Abort()

	return
}

func ErrorWithEnMessage(c *gin.Context, statusCode int, messageEn, errorCode string) {
	logs.GetAppLogger().Error(messageEn)

	c.JSON(statusCode, modelError{
		ErrorCode: errorCode,
		Success:   false,
		Code:      statusCode,
		Message: &struct {
			Indonesian *string `json:"indonesian"`
			English    *string `json:"english"`
		}{Indonesian: nil, English: &messageEn},
		Timestamp: time.Now().Unix(),
	})

	c.Abort()

	return
}

func getErrorStatusCode(err error) int {
	errString, ok := err.(*errors.ErrorString)
	if ok {
		return errString.Code()
	}

	// default http status code
	return http.StatusInternalServerError
}
