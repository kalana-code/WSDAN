package model

import "net/http"

//InnerResponse used for exchange info between utils and service file
type InnerResponse struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

//Response is used for arrange response in standed way
type Response struct {
	Program string
	Version string
	Status  string
	Code    int
	Message string
	Data    map[string]interface{}
}

// Default set default value to the response
func (obj *Response) Default() {
	obj.Program = "Beq"
	obj.Version = "0.01"
}

// BadRequest set as Bad Request
func (obj *Response) BadRequest() {
	obj.Code = http.StatusBadRequest
	obj.Status = "Failed"
	obj.Message = "Bad Request"
}

// InternalServer set as Bad Request
func (obj *Response) InternalServerError() {
	obj.Code = http.StatusInternalServerError
	obj.Status = "Failed"
	obj.Message = "Internal Server Error"
}

// UserLogin is model is used for login process
type UserLogin struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// UserInfo is used for fetch user data from database
type UserInfo struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Role      string `json:"Role"`
	Gender    string `json:"Gender"`
	BirthDay  string `json:"BirthDay"`
}