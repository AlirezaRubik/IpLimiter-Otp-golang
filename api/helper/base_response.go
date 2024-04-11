package helper

import "information/api/validation"

type BaseResponse struct {
	Result        interface{}                               `json:"result"`
	Success       bool                                      `json:"success"`
	ClientCode    int                                       `json:"client_code"`
	Err           any                                     `json:"error"`
	ErrorProvider *[]validation.BaseResponseErrorValidation `json:"error_provider"`
}
func GenerateBaseReseponse(result any, success bool, resultClientCode int) *BaseResponse {
	return &BaseResponse{Result: result, Success: success, ClientCode: resultClientCode}

}
func GenerateBaseReseponseWithError(result any, success bool, resultClientCode int, err error) *BaseResponse {
	return &BaseResponse{Result: result, Success: success, ClientCode: resultClientCode, Err: err.Error()}

}
func GenerateBaseReseponseWithValidationError(result any, success bool, resultClientCode int, err error) *BaseResponse {
	return &BaseResponse{Result: result, Success: success, ClientCode: resultClientCode, ErrorProvider: validation.ErrorProvider(err)}

}