package api

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Response struct {
	Message *string     `json:"message,omitempty"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type Context struct {
	DB     *gorm.DB
	Body   map[string]interface{}
	Params map[string]string
}

func SuccessAPIResponse(message *string, data *interface{}) *Response {
	return &Response{
		Data:    data,
		Message: message,
		Status:  "success",
	}
}

func ErrorAPIResponse(message string) *Response {
	return &Response{
		Data:    nil,
		Message: &message,
		Status:  "error",
	}
}

type RequestHandlerType func(*Context) (interface{}, error)

func RequestHandler(handler RequestHandlerType) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context().Value("apiContext").(*Context)
		res, err := handler(c)

		encoder := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			encoder.Encode(ErrorAPIResponse(err.Error()))
		} else {
			encoder.Encode(SuccessAPIResponse(nil, &res))
		}
	}
}
