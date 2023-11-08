package returnvalue

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Responses interface {
	ResponseOK(code int) error
	ResponseErr(code int, description string) error
	ResponseOKWithValue(int, interface{}) error
}

type returnResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type returnValue struct {
	Val interface{}
}

type Response struct {
	w http.ResponseWriter
}

var (
	MarshalErr = errors.New("JSON marshalling error!")
)

func NewResponse(w http.ResponseWriter) Responses {
	return &Response{
		w: w,
	}
}

func (ret *Response) ResponseOK(code int) error {
	return ret.makeResponse(code, "")
}

func (ret *Response) ResponseOKWithValue(code int, value interface{}) error {
	jsonResp, err := json.Marshal(returnValue{
		Val: value,
	})
	if err != nil {
		return fmt.Errorf("%w", MarshalErr)
	}

	ret.w.WriteHeader(code)
	ret.w.Header().Set("Content-Type", "application/json")
	ret.w.Write(jsonResp)

	return nil
}

func (ret *Response) ResponseErr(code int, description string) error {
	return ret.makeResponse(code, description)
}

func (ret *Response) makeResponse(code int, description string) error {
	jsonResp, err := json.Marshal(returnResponse{
		Code:        code,
		Description: description,
	})
	if err != nil {
		return fmt.Errorf("%w", MarshalErr)
	}

	ret.w.WriteHeader(code)
	ret.w.Header().Set("Content-Type", "application/json")
	ret.w.Write(jsonResp)

	return nil
}
