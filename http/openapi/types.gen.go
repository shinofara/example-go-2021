// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package oapi

// Signup defines model for signup.
type Signup struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Todo defines model for todo.
type Todo struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// PostSignupJSONBody defines parameters for PostSignup.
type PostSignupJSONBody Signup

// TodoJSONBody defines parameters for Todo.
type TodoJSONBody Todo

// PostSignupJSONRequestBody defines body for PostSignup for application/json ContentType.
type PostSignupJSONRequestBody PostSignupJSONBody

// TodoJSONRequestBody defines body for Todo for application/json ContentType.
type TodoJSONRequestBody TodoJSONBody

