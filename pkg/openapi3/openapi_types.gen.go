// Package openapi3 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package openapi3

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

// CreateUserResponse defines model for CreateUserResponse.
type CreateUserResponse struct {
	Id       *string `json:"id,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// LoginUserRequest defines model for LoginUserRequest.
type LoginUserRequest struct {
	// The password for login in clear text
	Password string `json:"password"`

	// The user name for login
	UserName string `json:"userName"`
}

// LoginUserResonse defines model for LoginUserResonse.
type LoginUserResonse struct {
	// A url for websoket API with a one-time token for starting chat
	Url string `json:"url"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody CreateUserRequest

// LoginUserJSONBody defines parameters for LoginUser.
type LoginUserJSONBody LoginUserRequest

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// LoginUserJSONRequestBody defines body for LoginUser for application/json ContentType.
type LoginUserJSONRequestBody LoginUserJSONBody
