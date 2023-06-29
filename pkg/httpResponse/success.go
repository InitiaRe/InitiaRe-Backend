package httpResponse

import "fmt"

// Rest response interface
type RestRes interface {
	Status() int
	Message() string
	Res() interface{}
}

// Rest response struct
type RestResponse struct {
	ResStatus  int         `json:"status,omitempty"`
	ResMessage string      `json:"message,omitempty"`
	Result     interface{} `json:"res,omitempty"`
}

// Response Message() interface method
func (r RestResponse) Message() string {
	return fmt.Sprintf("status: %d - message: %s - response: %v", r.ResStatus, r.ResMessage, r.Result)
}

// Response status
func (r RestResponse) Status() int {
	return r.ResStatus
}

// Rest get result
func (r RestResponse) Res() interface{} {
	return r.Result
}

// New Rest Response
func NewRestResponse(status int, message string, response interface{}) RestRes {
	return RestResponse{
		ResStatus:  status,
		ResMessage: message,
		Result:     response,
	}
}
