package zmodels

import (
	"encoding/json"
	"fmt"
)

// SuccessResponse is the SendHTMLEmail() response object
type SuccessResponse struct {
	Data      []SuccessData `json:"data"`
	Message   string        `json:"message"`
	RequestId string        `json:"request_id"`
	Object    string        `json:"object"`
}

// SuccessData is the Data object for the SuccessResponse object
type SuccessData struct {
	Code           string        `json:"code"`
	AdditionalInfo []interface{} `json:"additional_info"`
	Message        string        `json:"message"`
}

// Failure mapping

type FailureResponse struct {
	ErrorData Error `json:"error"`
}
type Error struct {
	Code      string         `json:"code,omitempty"`
	Details   []ErrorDetails `json:"details,omitempty"`
	Message   string         `json:"message,omitempty"`
	RequestID string         `json:"request_id,omitempty"`
}
type ErrorDetails struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Target  string `json:"target,omitempty"`
}

func (f *FailureResponse) Error() string {
	detail, _ := json.Marshal(f.ErrorData.Details)
	return fmt.Sprint("Error Code:-" + f.ErrorData.Code + " Error Details:-" + (string(detail)) + " Error Msg:- " + f.ErrorData.Message + " Error request id:- " + f.ErrorData.RequestID)
}
