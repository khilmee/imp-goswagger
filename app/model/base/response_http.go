package base

import (
	"context"
	"encoding/json"
	"imp-goswagger/helper/message"
	"net/http"
	"reflect"
)

//swagger:response
type ResponseHTTP struct {
	//in: body
	Response responseHttp `json:"response"`
}

// swagger:model SuccessResponse
type responseHttp struct {
	// Meta is the API response information
	// in: struct{}
	Meta metaResponse `json:"meta"`
	// Data is our data
	// in: struct{}
	Data data `json:"data"`
	// Errors is the response message
	//in: string
	Errors interface{} `json:"errors,omitempty"`
}

type metaResponse struct {
	// Code is the response code
	// example: 1000
	Code int `json:"code"`
	// Message is the response message
	// example: Success
	Message string `json:"message"`
}

type emptyStruct struct{}

type data struct {
	Records interface{} `json:"records,omitempty"`
	Record  interface{} `json:"record,omitempty"`
}

func SetHttpResponse(code int, message string, result interface{}) interface{} {
	dt := data{}
	isSlice := reflect.ValueOf(result).Kind() == reflect.Slice
	isNil := IsNil(result)
	if isSlice {
		if isNil {
			dt.Records = []emptyStruct{}
		} else {
			dt.Records = result
		}
	} else {
		if isNil {
			dt.Record = emptyStruct{}
		} else {
			dt.Record = result
		}
	}

	return responseHttp{
		Meta: metaResponse{
			Code:    code,
			Message: message,
		},
		Data: dt,
	}
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

type errorer interface {
	error() error
}

type errorResponse struct {
	// Meta is the API response information
	// in: struct{}
	Meta struct {
		// Code is the response code
		//in: int
		Code int `json:"code"`
		// Message is the response message
		//in: string
		Message string `json:"message"`
	} `json:"meta"`
	// Data is our data
	// in: struct{}
	Data interface{} `json:"data"`
	// Errors is the response message
	//in: string
	Errors interface{} `json:"errors,omitempty"`
}

func GetHttpResponse(resp interface{}) *responseHttp {
	result, ok := resp.(responseHttp)

	if ok {
		return &result
	}
	return nil
}

func EncodeResponseHTTP(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	if err, ok := resp.(errorer); ok && err.error() != nil {
		EncodeError(ctx, err.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	result := GetHttpResponse(resp)
	code := result.Meta.Code
	switch code {
	case message.ErrNotFound.Code:
		w.WriteHeader(http.StatusBadRequest)
	case message.SuccessMsg.Code:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	return json.NewEncoder(w).Encode(resp)
}

// Encode error, for HTTP
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result := &errorResponse{}
	result.Meta.Code = message.ErrServerTimeout.Code
	result.Meta.Message = message.ErrServerTimeout.Message

	if err != nil {
		result.Meta.Code = message.ErrValidation.Code
		result.Meta.Message = message.ErrValidation.Message
		result.Errors = err.Error()
	}

	_ = json.NewEncoder(w).Encode(result)
}
