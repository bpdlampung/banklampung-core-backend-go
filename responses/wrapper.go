package responses

import (
	"github.com/bpdlampung/banklampung-core-backend-go/entities"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type modelSuccess struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type modelPagingSuccess struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Paging    paging      `json:"paging"`
	Timestamp int64       `json:"timestamp"`
}

type paging struct {
	Page        int64 `json:"page"`
	TotalPage   int64 `json:"total_page"`
	ItemPerPage int64 `json:"item_per_page"`
	TotalItem   int64 `json:"total_item"`
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

func PagingSuccess(c *gin.Context, data interface{}, total int64, pagingFilter entities.Paging) {
	logs.GetAppLogger().InfoInterface(data)

	c.JSON(http.StatusOK, modelPagingSuccess{
		Data: data,
		Paging: paging{
			Page:        pagingFilter.Page,
			TotalPage:   total / pagingFilter.ItemPerPage,
			ItemPerPage: pagingFilter.ItemPerPage,
			TotalItem:   total,
		},
		Success:   true,
		Code:      http.StatusOK,
		Timestamp: time.Now().Unix(),
	})

	c.Abort()

	return
}

func Error(c *gin.Context, error error) {
	logs.GetAppLogger().Error(error.Error())

	statusCode := getErrorStatusCode(error)
	errorCode := errors.DefaultErrorCode
	errorMsgEn := error.Error()
	errorMsgId := error.Error()

	if getErrorResponseCode(error) != "-" {
		errorInfo := errors.GetErrorInfo(getErrorResponseCode(error))

		statusCode = errorInfo.StatusCode
		errorCode = errorInfo.ErrorCode
		if errorInfo.ErrorCode == errors.DefaultErrorCode || getErrorMessage(error) == "-" {
			errorMsgEn = errorInfo.MsgEn
			errorMsgId = errorInfo.MsgId
		}
	}

	c.JSON(statusCode, modelError{
		ErrorCode: errorCode,
		Success:   false,
		Code:      statusCode,
		Message: &struct {
			Indonesian *string `json:"indonesian"`
			English    *string `json:"english"`
		}{Indonesian: &errorMsgId, English: &errorMsgEn},
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

func getErrorResponseCode(err error) string {
	errString, ok := err.(*errors.ErrorString)

	if ok && errString.ResponseCode() != "" {
		return errString.ResponseCode()
	}

	return "-"
}

func getErrorStatusCode(err error) int {
	errString, ok := err.(*errors.ErrorString)

	if ok && errString.StatusCode() != 0 {
		return errString.StatusCode()
	}

	// default http status code
	return http.StatusInternalServerError
}

func getErrorMessage(err error) string {
	errString, ok := err.(*errors.ErrorString)

	if ok && errString.Message() != "" {
		return errString.Message()
	}

	return "-"
}
