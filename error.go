package slingr

import (
	"fmt"
	"net/http"
	"strings"
)

var apiError APIError

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Errors  []struct {
		Field          string      `json:"field"`
		FieldLabel     string      `json:"fieldLabel"`
		Code           string      `json:"code"`
		Message        string      `json:"message"`
		AdditionalInfo interface{} `json:"additionalInfo"`
	} `json:"errors"`
	URL        string
	HttpMethod string
	HttpStatus int
}

func (e APIError) Error() string {
	httpErrorMsg := fmt.Sprintf("%d %s", e.HttpStatus, http.StatusText(e.HttpStatus))
	var str strings.Builder
	if len(e.Errors) > 0 {
		for _, v := range e.Errors {
			str.WriteString(fmt.Sprintf("\n  - %s %s", v.Field, v.Message))
		}
	}
	return fmt.Sprintf("%s %s\n%s\n%s%s", e.HttpMethod, e.URL, httpErrorMsg, e.Message, str.String())
}
