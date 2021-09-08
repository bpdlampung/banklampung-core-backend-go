package errors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	FileNameErrorCodeCollection = "./configs/resources/error-codes.json"
	DefaultErrorCode            = "CORE-999"
)

var errorCodes = make(map[string]ErrorCode)

type ErrorCode struct {
	ErrorCode  string `json:"error-code"`
	StatusCode int    `json:"status-code"`
	MsgEn      string `json:"msg-en"`
	MsgId      string `json:"msg-id"`
}

func init() {
	var errCodes []ErrorCode

	plan, _ := ioutil.ReadFile(FileNameErrorCodeCollection)

	err := json.Unmarshal(plan, &errCodes)

	if err != nil {
		panic(err)
	}

	for _, errCode := range errCodes {
		errorCodes[errCode.ErrorCode] = errCode
	}
}

func GetErrorInfo(errorCode string) ErrorCode {
	val, ok := errorCodes[errorCode]

	if ok {
		return val
	}

	return ErrorCode{
		ErrorCode:  DefaultErrorCode,
		StatusCode: http.StatusInternalServerError,
		MsgEn:      fmt.Sprintf("Error Code %s is not defined", errorCode),
		MsgId:      fmt.Sprintf("Error Code %s tidak terdefinisi", errorCode),
	}
}
