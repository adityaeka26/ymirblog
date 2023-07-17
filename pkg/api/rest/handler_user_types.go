// Package rest is port handler.
package rest

import "github.com/abialemuel/ymirblog/pkg/entity"

// GetUserRequest Get a User request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type GetUserRequest struct {
	
}

type GetUserRequestID struct {
	ID    int    `json:"id"`
}
// GetUserResponse Get a User response.  /** PLEASE EDIT THIS EXAMPLE, return handler response */.
type GetUserResponse struct {
	Message string        `json:"message"`
	User    *entity.User  `json:"user,omitempty"`
	Users   []entity.User `json:"users,omitempty"`
}

type UserResponse struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type GetAllUserRespone struct{
	Message string        `json:"message"`
	Items []UserResponse `json:"items"`
}



// GetUserCSVRequest Get a User request.  /** PLEASE EDIT THIS EXAMPLE, request handler */.
type GetUserCSVRequest struct {
}

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}
