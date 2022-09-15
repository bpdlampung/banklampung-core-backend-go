package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/entities"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
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

type ModelError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code,omitempty"`
	Success   bool   `json:"success"`
	Message   *struct {
		Indonesian *string `json:"indonesian"`
		English    *string `json:"english"`
	} `json:"message,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

func prettyPrint(b []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")

	if err != nil {
		return b
	}

	return out.Bytes()
}

func Success(c *gin.Context, data interface{}) {
	marshaled, _ := json.Marshal(data)

	prettyData := prettyPrint(marshaled)

	logs.GetAppLogger().InfoInterface(fmt.Sprintf("[RESPONSE] path : %s | response : %s", c.FullPath(), string(prettyData)))

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
	marshaled, _ := json.Marshal(data)

	prettyData := prettyPrint(marshaled)

	logs.GetAppLogger().InfoInterface(fmt.Sprintf("[RESPONSE] path : %s | response : %s", c.FullPath(), string(prettyData)))

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

	stacktrace := fmt.Sprintf(" Message: %s", error.Error())

	for i := 1; i <= 10; i++ {
		pc, file, line, _ := runtime.Caller(i)
		f := runtime.FuncForPC(pc)
		if f == nil || line == 0 {
			break
		}

		stacktrace += fmt.Sprintf("\n --- at %s:%d ---", file, line)
	}

	logs.GetAppLogger().Error(stacktrace)
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

	c.JSON(statusCode, ModelError{
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

	c.JSON(statusCode, ModelError{
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

	c.JSON(statusCode, ModelError{
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

	c.JSON(statusCode, ModelError{
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
