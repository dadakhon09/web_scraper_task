package models

type Empty struct{}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseOK struct {
	Message interface{}
}

type Obj struct {
	Link  string `json:"link"`
	Title string `json:"title"`
}
type Response struct {
	NumOfSuccessCalls int32 `json:"num_of_success_calls"`
	NumOfFailedCalls  int32 `json:"num_of_failed_calls"`
	Titles            []Obj `json:"titles"`
}
