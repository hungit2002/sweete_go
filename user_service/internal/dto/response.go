package dto

import "net/http"

type Meta struct {
	Code    int    `json:"code" yaml:"code"`
	Message string `json:"message" yaml:"message"`
	Status  string `json:"status" yaml:"status"`
}

type HttpSuccessResponse struct {
	Meta   Meta        `json:"meta" yaml:"meta"`
	Result interface{} `json:"result" yaml:"result"`
}

type HttpError struct {
	Meta   Meta        `json:"meta" yaml:"meta"`
	Result interface{} `json:"result" yaml:"result"`
}

func ErrorResponse(code int, message interface{}, errors error) *HttpError {
	//logger.GetInstance().Error(errors)

	messageNew := http.StatusText(code)
	if messageString, ok := message.(string); ok {
		messageNew = messageString
	}
	res := HttpError{
		Meta: Meta{
			Code:    code,
			Message: messageNew,
		},
		Result: errors,
	}
	return &res
}

func NoContentHttpResponse(message interface{}, err error) *HttpError {
	return ErrorResponse(http.StatusNoContent, message, err)
}

func BadRequestErrorResponse(message interface{}, err error) *HttpError {
	return ErrorResponse(http.StatusBadRequest, message, err)
}

func InternalErrorResponse(message interface{}, err error) *HttpError {
	return ErrorResponse(http.StatusInternalServerError, message, err)
}

func SuccessResponse(message string, data interface{}) *HttpSuccessResponse {
	res := HttpSuccessResponse{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
			Status:  http.StatusText(http.StatusOK),
		},
		Result: data,
	}
	return &res
}
