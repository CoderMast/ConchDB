package result

import (
	"encoding/json"
	"fmt"
)

// Result 返回值对象
type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功的返回值封装
func Success(message string, data interface{}) string {
	result := new(Result)
	result.Code = successCode
	result.Message = message
	result.Data = data

	resultJson, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf("%s", err)
	}

	return fmt.Sprintf("%s", resultJson)
}

// Error 失败的返回值封装
func Error(message string) string {
	result := new(Result)
	result.Code = errorCode
	result.Message = message

	resultJson, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf("%s", err)
	}
	return fmt.Sprintf("%s", resultJson)
}
